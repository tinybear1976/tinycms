package apis

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq"
	"github.com/tinybear1976/tinycms/defines"
)

func currentUser(c *gin.Context) string {
	return c.GetString(defines.KEYUserInfo)
}

func getRequestJsonQ(c *gin.Context) *gojsonq.JSONQ {
	raw, _ := c.GetRawData()
	sjsonbody := string(raw)
	jq := gojsonq.New().FromString(sjsonbody)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(raw))
	return jq
}

func getRequestJsonString(c *gin.Context) string {
	raw, _ := c.GetRawData()
	sjsonbody := string(raw)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(raw))
	return sjsonbody
}

// func stringToJson(sjson string) (*interface{}, error) {
// 	var r interface{}
// 	err := json.Unmarshal([]byte(sjson), &r)
// 	return &r, err
// }

func stringToMap(sjson string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	bytes := []byte(sjson)
	err := json.Unmarshal(bytes, &dat)
	return dat, err
}
