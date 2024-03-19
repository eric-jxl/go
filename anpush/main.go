package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var S, T *string

func init() {
	T = flag.String("t", "通知", "输入通知标题")
	S = flag.String("w", "", "输入通知内容(支持Markdown语法)")
	flag.Parse()
}
func main() {
	url := fmt.Sprintf("https://api.anpush.com/push/%s", os.Getenv("ANPUSH_TOKEN"))
	payload := strings.NewReader(fmt.Sprintf("title=%s&content=%s&channel=79692", *T, *S))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
