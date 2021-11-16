package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tinybear1976/tinycms/apis"
	mw "github.com/tinybear1976/tinycms/middleware"
)

func initRbacRoutes(r *gin.Engine) {
	var url string = "/api/v1/rbac"
	v1 := r.Group(url)
	v1.Use(mw.CrossDomain())
	//v1.Use(mw.Authenticate())
	{
		v1.POST("/", apis.Rbac_GetUiPerimissionApi)
	}
}