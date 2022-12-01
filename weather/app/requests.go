package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eric-jxl/go/weather/models"
)

type ShortDao struct{}

func New() (*ShortDao, error) {
	return &ShortDao{}, nil
}

// GetCityWeather 获取目标未来一周的城市天气
func (dao *ShortDao) GetCityWeather(location string) (*models.Response, error) {
	var urls = fmt.Sprintf("https://api.66mz8.com/api/weather.php?location=%s",location)
	resp, err := http.Get(urls)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	var response models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
