package main

import (
	"Workspace/shortUrl/controller"
	"flag"
	"fmt"
)

func main() {
	str := flag.String("u","https://www.icourse163.org","请输入要缩短的长链接")
	flag.Parse()
	resp ,err := controller.GenShortUrl(*str)
	if err != nil {
		return
	}
	fmt.Printf("short_url: %s\n",resp.UrlShort)
}
