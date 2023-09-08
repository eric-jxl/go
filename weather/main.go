package main

import (
	"flag"
	"fmt"

	"github.com/eric-jxl/go/weather/app"
)

var S *string

func init() {
	S = flag.String("u", "苏州", "输入要查询天气的城市!")
	flag.Parse()
}
func main() {
	local, _ := app.New()
	resp, err := local.GetCityWeather(*S)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", resp)
}
