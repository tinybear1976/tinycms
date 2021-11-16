/*
 * @Author: your name
 * @Date: 2021-11-10 12:17:55
 * @LastEditTime: 2021-11-10 13:50:48
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /tinycms/config/mariadb.go
 */
package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
	mariadb "github.com/tinybear1976/database-mariadb"
	"github.com/tinybear1976/localsystem/logger"
	"github.com/tinybear1976/tinycms/debugging"
	"github.com/tinybear1976/tinycms/defines"
)

func InitMariadb() {
	initTinyCMS()
	initLogClientIp()
}

func initTinyCMS() {
	err := mariadb.New(
		defines.DB_MAIN,
		viper.GetString("mariadb.tcms.server"),
		viper.GetString("mariadb.tcms.port"),
		viper.GetString("mariadb.tcms.uid"),
		viper.GetString("mariadb.tcms.pwd"),
		viper.GetString("mariadb.tcms.db"))
	if err != nil {
		logger.Log.Panic("init db for tinycms faile." + err.Error())
		panic("init db for tinycms faile " + err.Error())
	}
}
func initLogClientIp() {
	err := mariadb.New(
		defines.DB_LOG_CLIENTIP,
		viper.GetString("mariadb.logclientip.server"),
		viper.GetString("mariadb.logclientip.port"),
		viper.GetString("mariadb.logclientip.uid"),
		viper.GetString("mariadb.logclientip.pwd"),
		viper.GetString("mariadb.logclientip.db"))
	if err != nil {
		logger.Log.Panic("init db for logclientip faile." + err.Error())
		panic("init db for logclientip faile " + err.Error())
	}
	once_check_log_clientip_table_exists()
}

func once_check_log_clientip_table_exists() {
	filename := "./scripts/logclientip.sql"

	f, err := os.Open(filename)
	if err != nil {
		logger.Log.Panic("missing script file: " + filename)
		panic("missing script file: " + filename)
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Log.Panic("read script file error: " + filename)
		panic("read script file error: " + filename)
	}
	temp := string(bytes)
	sql := strings.Replace(temp, defines.SCRP_TAG_LOG_CLIENTIP_TABLENAME, viper.GetString("mariadb.logclientip.table"), -1)
	conn, err := mariadb.Connect(defines.DB_LOG_CLIENTIP)
	if err != nil {
		logger.Log.Panic("exec script file connection error: " + filename)
		panic("exec script file connection error: " + filename)
	}
	debugging.Debug_ShowSql("log_clientip_create_table", sql)
	_, err = conn.Exec(sql)
	if err != nil {
		logger.Log.Panic("exec script file error: " + filename)
		panic("exec script file error: " + filename)
	}
}
