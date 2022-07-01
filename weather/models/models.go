package models


type RespData struct {
	Days string `json:"days"`
	Week string `json:"week"`
	Temperature string `json:"temperature"`
	Weather string `json:"weather"`
	WeatherIcon string `json:"weather_icon"`
	Wind string `json:"wind"`
}


type Response struct {
	Code int
	Citynm string
	Cityno string
	Data []RespData
}