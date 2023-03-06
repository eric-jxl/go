package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eric-jxl/go/shortUrl/models"
)

func GenShortUrl(url string) (*models.Response, error) {
	resp, err := http.Get("https://api.sumt.cn/api/short.php?dwz=suoim&url=" + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rd, _ := ioutil.ReadAll(resp.Body)
	resu := models.Response{}
	_ = json.Unmarshal(rd, &resu)
	return &resu, nil
}
