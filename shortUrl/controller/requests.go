package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eric-jxl/go/shortUrl/models"
)

func GenShortUrl(url string) (*models.NResp, error) {
	//resp, err := http.Get("https://api.sumt.cn/api/short.php?dwz=suoim&url=" + url) ⚠️废弃
	resp, err := http.Get("https://api.uomg.com/api/long2dwz?dwzapi=urlcn&url=" + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rd, _ := ioutil.ReadAll(resp.Body)
	resu := models.NResp{}
	_ = json.Unmarshal(rd, &resu)
	return &resu, nil
}
