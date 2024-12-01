package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/eric-jxl/go/web/middleware"
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
		response.Fatal(c, "参数错误")
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
	r.Use(middleware.Logger())
	api := r.Group("/api")
	{
		api.GET("/v1", func(c *gin.Context) {
			response.Success(c, "v1")
		})
		api.GET("/img", func(c *gin.Context) {
			response.Success(c, content)
		})
		api.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
		api.GET("/cookie", func(c *gin.Context) {
			cookie, err := c.Cookie("gin_cookie")
			if err != nil {
				cookie = "NotSet"
				c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
			}

			fmt.Printf("Cookie value: %s \n", cookie)
		})

		r.MaxMultipartMemory = 8 << 20 // 8 MiB
		r.POST("/upload", func(c *gin.Context) {
			// 单文件
			file, _ := c.FormFile("file")
			log.Println(file.Filename)
			dst := "./" + file.Filename
			// 上传文件至指定的完整文件路径
			c.SaveUploadedFile(file, dst)
			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!\n", file.Filename))
		})
	}

	_ = r.Run(":8999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
