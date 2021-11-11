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

	routing "github.com/tinybear1976/tinycms/routes/v1"

	"path"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tinybear1976/localsystem"
	"github.com/tinybear1976/localsystem/logger"
	"github.com/tinybear1976/tinycms/config"
)

var (
	arg_show_version bool
)

func main() {
	flag.BoolVar(&arg_show_version, "version", false, "usage version show version")
	flag.Parse()
	version := "1.0.0"
	if arg_show_version {
		fmt.Printf("TinyCMS Api Service current version: %s\n", version)
		return
	}
	gin.SetMode(gin.ReleaseMode)
	p, err := localsystem.CurrentDirectory()
	if err != nil {
		logger.Log.Warn("failed to get the current service running path")
		p = "."
	}

	logger.Log = logger.InitLogger(path.Join(p, "tinycms.log"), "")
	if logger.Log == nil {
		panic("log creation failed, program operation aborted")
	}
	config.InitSpecificConfig("cfg_tinycms", "yaml", p)
	config.InitMariadb()
	r := routing.InitAllRoutes()
	ip := viper.GetString("publishing.server")
	s_info := fmt.Sprintf("TinyCMS Api Service [ver:%s] Running at %s", version, ip)
	logger.Log.Info(s_info)
	fmt.Printf("%s\n", s_info)
	r.Run(ip)
}
