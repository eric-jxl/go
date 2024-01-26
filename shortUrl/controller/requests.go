package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eric-jxl/go/shortUrl/models"
)

func GenShortUrl(url string) (*models.Resp, error) {
	resp, err := http.Get("https://api.uomg.com/api/long2dwz?dwzapi=urlcn&url=" + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rd, _ := io.ReadAll(resp.Body)
	resu := models.Resp{}
	_ = json.Unmarshal(rd, &resu)
	return &resu, nil
}
