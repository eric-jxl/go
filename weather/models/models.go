package models

type RespData struct {
	Days        string `json:"days"`
	Week        string `json:"week"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
	Wind        string `json:"wind"`
	WeatherIcon string `json:"weather_icon"`
}

type Response struct {
	Code   int
	Citynm string
	Cityno string
	Data   []RespData
}

type AmapWeather struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Lives    []struct {
		Province         string `json:"province"`
		City             string `json:"city"`
		Adcode           string `json:"adcode"`
		Weather          string `json:"weather"`
		Temperature      string `json:"temperature"`
		Winddirection    string `json:"winddirection"`
		Windpower        string `json:"windpower"`
		Humidity         string `json:"humidity"`
		Reporttime       string `json:"reporttime"`
		TemperatureFloat string `json:"temperature_float"`
		HumidityFloat    string `json:"humidity_float"`
	} `json:"lives"`
}
