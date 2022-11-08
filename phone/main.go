package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

const appKey = "e6a0d40abd02adbeb1105c4d757a6284"

type PhoneStruct struct {
	Code   string `json:"code"`
	Charge bool   `json:"charge"`
	Msg    string `json:"msg"`
	Result struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
		Result struct {
			Shouji   string `json:"shouji"`
			Province string `json:"province"`
			City     string `json:"city"`
			Company  string `json:"company"`
			Cardtype string `json:"cardtype"`
			Areacode string `json:"areacode"`
		} `json:"result"`
	} `json:"result"`
	RequestId string `json:"requestId"`
}

func requests(phonenum string) (string, error) {
	url := fmt.Sprintf("https://www.baifubao.com/callback?cmd=1059&callback=phone&phone=%s", phonenum)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bd, _ := ioutil.ReadAll(resp.Body)
	return string(bd), nil
}

func request_jdApi(phone string) error {
	url := fmt.Sprintf("https://way.jd.com/jisuapi/query4?shouji=%s&appkey=%s", phone, appKey)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bd, _ := ioutil.ReadAll(resp.Body)
	var phones PhoneStruct
	_ = json.Unmarshal(bd, &phones)
	var out bytes.Buffer
	ps, _ := json.Marshal(phones.Result)
	json.Indent(&out, ps, " ", "\t")
	fmt.Printf("手机归属地：%v\n", out.String())
	return nil
}
func main() {
	phone := flag.String("p", "", "Usage phone number")
	flag.Parse()
	request_jdApi(*phone)
}
