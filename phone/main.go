package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func requests(phonenum string) (string,error) {
	url := fmt.Sprintf("https://www.baifubao.com/callback?cmd=1059&callback=phone&phone=%s", phonenum)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	bd, _ := ioutil.ReadAll(resp.Body)
	return string(bd),nil
}
func main() {
	phone := flag.String("p","","Usage phone number")
	flag.Parse()
	res,_ := requests(*phone)
	fmt.Println(res)
}
