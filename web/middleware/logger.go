package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		c.Set("example", "12345")
		// 请求前

		c.Next()
		// 请求后
		latency := time.Since(t)

		// 获取发送的 status
		status := c.Writer.Status()
		reqBody, _ := c.GetRawData()
		log.Printf("[INFO] %s %s %s %s %s %d\n", gin.Mode(), c.Request.Method, c.Request.RequestURI, reqBody, latency, status)
	}
}
