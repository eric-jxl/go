package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eric-jxl/go/weather/models"
)

type ShortDao struct{}

func New() (*ShortDao, error) {
	return &ShortDao{}, nil
}

// GetCityWeather 获取目标城市天气
// city 城市编码
func (dao *ShortDao) GetCityWeather(location string) (string, error) {
	var key = "xxxxxx" //高德地图key密钥
	var urls = fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&output=json&city=%s", key, location)
	resp, err := http.Get(urls)
	if err != nil {
		return "nil", err
	}

	// another Way To Write
	//req, err := http.NewRequest("GET", urls, nil)
	//req.Header.Set("Content-Type", "application/json")
	//resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var response models.AmapWeather
	err = json.Unmarshal(body, &response)
	res, _ := json.Marshal(response.Lives)
	var out bytes.Buffer
	err = json.Indent(&out, res, " ", "\t")
	if err != nil {
		return "nil", err
	}
	return out.String(), nil
}
