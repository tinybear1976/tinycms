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
	"github.com/spf13/viper"
	mariadb "github.com/tinybear1976/database-mariadb"
	"github.com/tinybear1976/localsystem/logger"
	"github.com/tinybear1976/tinycms/defines"
)

func InitMariadb() {
	initTinyCMS()
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
