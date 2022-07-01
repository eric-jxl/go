package models

type Response struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	UrlLong  string `json:"url_long"`
	UrlShort string `json:"url_short"`
}
