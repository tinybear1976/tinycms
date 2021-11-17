package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinybear1976/tinycms/rbac"
)

func Front_GetUiPermissionApi(c *gin.Context) {
	jq := getRequestJsonQ(c)
	find_rst := jq.Find("role_id")
	if find_rst == nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": "missing key 'role_id'",
		})
		return
	}
	role_id, ok := find_rst.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": "role_id must be a string",
		})
		return
	}
	j, err := rbac.GetUiPermission_front(role_id)
	fmt.Println(j)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error1": err.Error(),
		})
		return
	}
	r, err := stringToMap(j)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error2": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}
