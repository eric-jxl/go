package models

type Response struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	UrlLong  string `json:"url_long"`
	UrlShort string `json:"url_short"`
}

type Resp struct {
	Code       int    `json:"code"`   //返回的状态码
	ShortenUrl string `json:"ae_Url"` //返回缩短后的短网址
	Msg        string `json:"msg"`    //返回错误提示信息！
}
