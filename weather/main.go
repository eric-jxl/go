package main

import (
	"Workspace/weather/app"
	"flag"
	"fmt"
)

func main() {
	s := flag.String("u","上海","输入要查询天气的城市!")
	flag.Parse()
	local, _ := app.New()
	resp, err := local.GetCityWeather(*s)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _,v := range resp.Data{
		fmt.Println(v)
	}
}
