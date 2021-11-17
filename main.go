/*
 * @Author: wang
 * @Date: 2021-11-10 12:05:18
 * @LastEditTime: 2021-11-10 13:52:59
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /tinycms/main.go
 */
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/tinybear1976/localsystem"
	"github.com/tinybear1976/tinycms/defines"
	routing "github.com/tinybear1976/tinycms/routes/v1"

	"github.com/spf13/viper"
	"github.com/tinybear1976/localsystem/logger"
	"github.com/tinybear1976/tinycms/config"
)

var (
	arg_show_version bool
)

func main() {
	flag.BoolVar(&arg_show_version, "version", false, "usage version show version")
	flag.Parse()
	if arg_show_version {
		fmt.Printf("TinyCMS Api Service current version: %s\n", defines.VERSION)
		return
	}
	p, err := localsystem.CurrentDirectory()
	if err != nil {
		p = "."
	}
	config.InitLoggers(p)
	config.InitSpecificConfig("cfg_tinycms", "yaml", p)
	config.InitMariadb()
	config.InitRedis()
	if !viper.GetBool(defines.YML_DEBUG_GIN) {
		gin.SetMode(gin.ReleaseMode)
	}
	r := routing.InitAllRoutes()
	ip := viper.GetString("publishing.server")
	s_info := fmt.Sprintf("TinyCMS Api Service [ver:%s] Running at %s  (start time: %s)", defines.VERSION, ip, time.Now().Format(defines.FORMATDATETIME))
	logger.LogContainer[defines.LOG_APP].Info(s_info)
	fmt.Print(defines.GetLogo(ip))
	if viper.GetBool(defines.YML_DEBUG_PPROF) {
		pprof.Register(r)
	}
	r.Run(ip)
}
