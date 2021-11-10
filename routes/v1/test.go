/*
 * @Author: your name
 * @Date: 2020-09-17 18:18:29
 * @LastEditTime: 2021-11-10 12:32:02
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /tinycms/Users/wang/go/templates/apiservice/routes/v1/test.go
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tinybear1976/tinycms/apis"
	mw "github.com/tinybear1976/tinycms/middleware"
)

func initTestRoutes(r *gin.Engine) {
	var url string = "/api/v1/test"
	v1 := r.Group(url)
	v1.Use(mw.CrossDomain())
	{
		v1.POST("/", apis.TestApi)
	}
}
