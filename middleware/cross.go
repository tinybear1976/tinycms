/*
 * @Author: your name
 * @Date: 2021-11-10 12:29:18
 * @LastEditTime: 2021-11-10 12:31:55
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /tinycms/middleware/cross.go
 */
package middleware

import "github.com/gin-gonic/gin"

func CrossDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		// origin := c.Request.Header.Get("origin")
		// if len(origin) == 0 {
		// 	origin = c.Request.Header.Get("Origin")
		// }
		// c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// func corsWavFile() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// origin := c.Request.Header.Get("origin")
// 		// if len(origin) == 0 {
// 		// 	origin = c.Request.Header.Get("Origin")
// 		// }
// 		// c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Content-Type", "audio/x-wav")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}
// 		c.Next()
// 	}
// }
