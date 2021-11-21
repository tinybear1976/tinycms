package apis

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tinybear1976/largefileupload"
	"github.com/tinybear1976/tinycms/defines"
)

func TestApi(c *gin.Context) {
	//c.ClientIP()
	t1 := setToken("U001")
	fmt.Println("t1", t1)
	t2 := setToken2(t1, "d502")
	fmt.Println("t2", t2)

	err, ok := checkToken(t1, t2)
	fmt.Println(err, ok)
	fmt.Println(c.Request.Header.Get("abc"))
	fmt.Println(c.Request.Header)
	// c64 := c.Request.Header.Get("authorization")
	// infos := strings.Split(c64, " ")
	// sdec, _ := base64.StdEncoding.DecodeString(infos[1])
	// fmt.Println(string(sdec))

	// filename := "d:/test/1/2/3/4/main.go"
	// s := localsystem.FileNameOnly(filename)
	// println(s)
	// s = localsystem.FileNameWithoutPath(filename)
	// //~~~~~~~~~~~~~~ db demo start ~~~~~~~~~~~~~~~~~~~~
	// sql := fmt.Sprintf("select userid,username from users order by userid")
	// if viper.GetBool("logs.gz_optcount") {
	// 	logger.Log.Debug("话务员列表  |  " + sql)
	// }
	// ret := make([]models.TestModels, 0)
	// db, err := mariadb.Connect(defines.MAINDB)
	// if err != nil {
	// 	//return ret, err
	// }
	// err = db.Select(&ret, sql)
	// if err != nil {
	// 	//return ret, err
	// }
	// //~~~~~~~~~~~~~~ db demo end ~~~~~~~~~~~~~~~~~~~~~~

	// sendOKJson(c, "")
}

func setToken(userid string) string {
	niao := time.Now().UnixNano()
	seed := userid + strconv.FormatInt(niao, 16)
	ciphertext := md5.Sum([]byte(seed))
	t := hex.EncodeToString(ciphertext[:])
	return t
}

func setToken2(t1, key string) string {
	//niao := time.Now().UnixNano()
	//seed := userid + strconv.FormatInt(niao, 16)
	seed := key + t1
	m := md5.Sum([]byte(seed))
	token := hex.EncodeToString(m[:])
	var sb strings.Builder
	sb.WriteString(token[0:8])
	sb.WriteString(string(key[0]))
	sb.WriteString(token[8:16])
	sb.WriteString(string(key[1]))
	sb.WriteString(token[16:24])
	sb.WriteString(string(key[2]))
	sb.WriteString(token[24:32])
	sb.WriteString(string(key[3]))
	return sb.String()
}

func checkToken(t1, t2 string) (error, bool) {
	if len(t2) != 36 {
		return errors.New("t2 len must eq 36"), false
	}
	pos := 8
	step := 9
	var keybuild strings.Builder
	for i := 0; i < 4; i++ {
		keybuild.WriteString(string(t2[pos]))
		println(string(t2[pos]))
		pos += step
	}
	key := keybuild.String()
	if t2 != setToken2(t1, key) {
		return nil, false
	}
	return nil, true
}

func T_slice(c *gin.Context) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	new_ints := copy_slice(ints)
	fmt.Println("old:", ints)
	fmt.Println("new:", new_ints)
}

func copy_slice(ints []int) []int {
	my := make([]int, 3)
	copy(my, ints)
	return my
}

func T_LargefileUploadApi(c *gin.Context) {
	max_size := viper.GetInt64(defines.YML_UPLOAD_MAX_SIZE)
	fmt.Println("max_size", max_size)
	upload_path := viper.GetString(defines.YML_UPLOAD_SAVE_PATH)
	fmt.Println("save_path", upload_path)
	files, err := largefileupload.LargeFileUpload(c.Request, upload_path, max_size)
	fmt.Println("err", err)
	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}
