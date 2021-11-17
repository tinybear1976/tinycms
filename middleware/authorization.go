package middleware

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	crsa "github.com/tinybear1976/cross_authenticate"
	mariadb "github.com/tinybear1976/database-mariadb"
	"github.com/tinybear1976/localsystem/logger"
	"github.com/tinybear1976/redisdb2"
	"github.com/tinybear1976/tinycms/debugging"
	"github.com/tinybear1976/tinycms/defines"
)

const (
	AUTH_RED  = "Red"
	AUTH_BLUE = "Blue"
)

type Author struct {
	Url      string
	ClientIP string
	Mode     string
	User     string
	E        string
}

func (a *Author) authorization(c *gin.Context) (ok bool, err error) {
	ok = false
	err = nil
	if a == nil {
		return
	}
	if a.Mode == AUTH_RED {
		c.Set(defines.KEYUserInfo, a.User)
		ts := get_TS(a.User)
		ok, err = crsa.Authenticate(ts, a.E)
	} else {
		c.Set(defines.KEYUserInfo, "")
		t := get_T()
		ok, err = crsa.Authenticate(t, a.E)
	}
	return
}

func get_TS(userid string) string {
	//todo
	return ""
}

func get_T() string {
	//todo
	return ""
}

func getAuthor(auth_str, ip, url_path string) (*Author, error) {
	a := &Author{
		Url:      url_path,
		ClientIP: ip,
		Mode:     "",
		User:     "",
		E:        "",
	}
	if len(auth_str) <= 0 {
		go rec_clientip_user(a)
		go audit_auth_token_direct(ip, "")
		return nil, errors.New("authorization info is empty")
	}
	spl := strings.Split(auth_str, " ")
	if len(spl) != 2 {
		go rec_clientip_user(a)
		go audit_auth_token_direct(ip, "")
		return nil, errors.New("authorization info split must be to 2 parts")
	}
	if spl[0] != AUTH_RED && spl[0] != AUTH_BLUE {
		go rec_clientip_user(a)
		go audit_auth_token_direct(ip, "")
		return nil, errors.New("unrecognized authentication mode: " + spl[0])
	}
	bytes, err := base64.StdEncoding.DecodeString(spl[1])
	if err != nil {
		go rec_clientip_user(a)
		go audit_auth_token_direct(ip, "")
		return nil, err
	}
	data := strings.Split(string(bytes), ":")
	if len(data) != 2 {
		go rec_clientip_user(a)
		go audit_auth_token_direct(ip, "")
		return nil, errors.New("auth_data split must be to 2 parts")
	}
	// data[0] -> userid
	// data[1] -> E(len=36)
	a = &Author{
		Url:      url_path,
		ClientIP: ip,
		Mode:     string(spl[0]),
		User:     string(data[0]),
		E:        string(data[1]),
	}
	go rec_clientip_user(a)
	return a, nil
}

func authenticateByRequest(c *gin.Context, sAuth string) (ok bool, err error) {
	// request.head["authorization"]
	//sAuth = "Red dXNlcjE6cGFzc3dvcmQ="
	a, err := getAuthor(sAuth, c.ClientIP(), c.Request.URL.Path)
	if err != nil {
		audit_sleep()
		return false, err
	}
	if !audit_auth_token(a) {
		audit_sleep()
	}
	ok, err = a.authorization(c)
	return
}

func audit_sleep() {
	if !viper.GetBool(defines.YML_ADT_AUTHTOKEN_ALLOW) {
		return
	}
	//fmt.Println("audit false!!!!!!!")
	jam := time.Duration(viper.GetInt64(defines.YML_ADT_AUTHTOKEN_JAM))
	//fmt.Println(jam)
	//fmt.Println("audit sleep start at ", time.Now().Format(defines.FORMATDATETIME))
	time.Sleep(jam * time.Millisecond)
	//fmt.Println("audit sleep stop at ", time.Now().Format(defines.FORMATDATETIME))
}

func audit_auth_token_direct(ip, E string) {
	if !viper.GetBool(defines.YML_ADT_AUTHTOKEN_ALLOW) {
		return
	}
	conn, err := redisdb2.Connect(defines.REDIS_AUTH_TOKEN)
	if err != nil {
		logger.LogContainer[defines.LOG_APP].Error(err.Error())
		return
	}
	redisdb2.SET(conn, "BL::"+ip, E+";"+time.Now().Format(defines.FORMATDATETIME14))
	redisdb2.Diconnect(conn)
}

// 默认返回true，表示审计通过
func audit_auth_token(a *Author) (ret_ok bool) {
	ret_ok = true
	if !viper.GetBool(defines.YML_ADT_AUTHTOKEN_ALLOW) {
		return
	}
	conn, err := redisdb2.Connect(defines.REDIS_AUTH_TOKEN)
	if err != nil {
		logger.LogContainer[defines.LOG_APP].Error(err.Error())
		return
	}
	// 先检查审计出来的黑名单
	b_key := "BL::" + a.ClientIP
	ok, err := redisdb2.EXISTS(conn, b_key)
	if err == nil && ok {
		redisdb2.Diconnect(conn)
		return false
	}
	key := "token::" + a.E
	ok, err = redisdb2.EXISTS(conn, key)
	if err == nil {
		expire := uint(viper.GetInt64(defines.YML_ADT_AUTHTOKEN_EXPIRE))
		repeat := viper.GetInt(defines.YML_ADT_AUTHTOKEN_REPEAT)
		if ok {
			redisdb2.EXEC(conn, "PEXPIRE", key, expire)
			redisdb2.EXEC(conn, "INCR", key)
		} else {
			redisdb2.SETPX(conn, key, "0", expire)
		}
		s, err1 := redisdb2.GET(conn, key)
		v, err2 := strconv.Atoi(s)
		if err1 == nil && err2 == nil {
			if v >= repeat {
				// audit
				ret_ok = false
				redisdb2.SET(conn, "BL::"+a.ClientIP, a.E+";"+time.Now().Format(defines.FORMATDATETIME14))
			}
		}
	} else {
		logger.LogContainer[defines.LOG_APP].Error(err.Error())
	}
	redisdb2.Diconnect(conn)
	return
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if viper.GetBool(defines.YML_ADT_AUTHTOKEN_ALLOW) {
			auth := c.Request.Header.Get("authorization")
			if len(auth) == 0 {
				// Unauthorized
				// c.AbortWithStatus(401)
				go rec_clientip(c)
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"error": "Unauthorized",
				})
				return
			}

			ok, err := authenticateByRequest(c, auth)
			if err != nil {
				go rec_clientip(c)
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			if !ok {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"error": "Unauthorized",
				})
				return
			}
		} else {
			go rec_clientip(c)
		}
		c.Next()
	}
}

func rec_clientip(c *gin.Context) {
	if !viper.GetBool(defines.YML_ADT_LOG_CLIENTIP_ALLOW) {
		return
	}
	//fmt.Println("request url: ", c.Request.URL.Path)
	//fmt.Println("from ip: ", c.ClientIP())
	sql := "INSERT INTO " + viper.GetString("mariadb.logclientip.table") + " (recdate, clientip, urlpath, userid) VALUES (now(), '" + c.ClientIP() + "', '" + c.Request.URL.Path + "', '');"
	debugging.Debug_ShowSql("rec web client ip(no user)", sql)
	exec_record_clientip(sql)
}

func rec_clientip_user(a *Author) {
	if !viper.GetBool(defines.YML_ADT_LOG_CLIENTIP_ALLOW) {
		return
	}
	//fmt.Println("request url: ", a.Url)
	//fmt.Println("from ip: ", a.ClientIP)
	sql := "INSERT INTO " + viper.GetString("mariadb.logclientip.table") + " (recdate, clientip, urlpath, userid) VALUES (now(), '" + a.ClientIP + "', '" + a.Url + "', '" + a.User + "');"
	debugging.Debug_ShowSql("rec web client ip(user)", sql)
	exec_record_clientip(sql)
}

func exec_record_clientip(sql string) {
	conn, err := mariadb.Connect(defines.DB_LOG_CLIENTIP)
	if err != nil {
		logger.LogContainer[defines.LOG_APP].Error("rec client ip connection error: " + err.Error())
		return
	}
	_, err = conn.Exec(sql)
	if err != nil {
		logger.LogContainer[defines.LOG_APP].Error("rec client ip insert error: " + err.Error())
	}
}
