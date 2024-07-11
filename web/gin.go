package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/eric-jxl/go/web/response"
	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return fmt.Sprintf("%d/%02d/%02d %d:%d:%02d", year, month, day, hour, min, sec)
}

//go:embed sub.go
var content string

func main() {
	//r := gin.New()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(response.Exception())
	//_ = r.SetTrustedProxies([]string{"192.168.1.15:7890"})
	r.GET("/ping", func(c *gin.Context) {
		data := "Ping"
		response.Success(c, data)
	})

	r.POST("/create", func(c *gin.Context) {
		response.Success(c, "创建成功")
	})

	r.GET("/error", func(c *gin.Context) {
		response.Fail(c, http.StatusBadRequest, "参数错误")
	})

	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	//r.Delims("{[{", "}]}") //自定义界定符
	r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLFiles("templates/index.tmpl") //等价于
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Eric",
			"Unix":  time.Now(),
		})
	})

	r.GET("/img", func(c *gin.Context) {
		response.Success(c, content)
	})
	_ = r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
