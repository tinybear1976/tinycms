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
	//v1.Use(mw.Authenticate())
	{
		v1.POST("/", apis.TestApi)
		v1.POST("/copy_slice", apis.T_slice)
		v1.POST("/luf", apis.T_LargefileUploadApi)
	}
	r.StaticFile("/upload.html", "./html/upload.html")
}
