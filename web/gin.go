package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	hour,min,sec := t.Clock()
	return fmt.Sprintf("%d/%02d/%02d %d:%d:%02d", year, month, day,hour,min,sec)
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.15:7890"})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


	r.POST("/create",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"OK",
		})
	})

	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	//r.Delims("{[{", "}]}") //自定义界定符
	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLFiles("templates/index.tmpl") //等价于
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl", gin.H{
			"title":"Eric",
			"Unix":time.Now(),
		})
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}