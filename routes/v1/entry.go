package v1

import "github.com/gin-gonic/gin"

func InitAllRoutes() *gin.Engine {
	router := gin.Default()
	initTestRoutes(router)
	return router
}
