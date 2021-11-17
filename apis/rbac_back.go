package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinybear1976/tinycms/rbac"
)

func Rbac_GetUiPermissionApi(c *gin.Context) {
	jq := getRequestJson(c)
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
	j, err := rbac.GetUiPermission_back(role_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := stringToMap(j)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}

func Rbac_SaveUiPermissionApi(c *gin.Context) {
	j := getRequestJsonString(c)

	err := rbac.SaveUiPermission_back(j)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "save success",
	})
}
