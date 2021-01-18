package main

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func main() {

	client, err := sdk.NewClientWithAccessKey("cn-shanghai", "LTAI4G6M6CAHVYMgUCP5ZGiH", "NpTjtcqGop59TMb406QXytmLiX2B9D") // ak,sk信息
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Product = "FaceAPI"
	request.Domain = "face.cn-shanghai.aliyuncs.com"
	request.Version = "2018-12-03"
	request.Scheme = "https"
	request.ApiName = "DetectFace"
	request.QueryParams["ImageUrl"] = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1559655604341&di=3d6995f6dee1c4795d1827e754a00452&imgtype=0&src=http%3A%2F%2Fimg0.ph.126.net%2F90u9atgu46nnziAm1NMAGw%3D%3D%2F6631853916514183512.jpg" // Specify the requested regionId, if not specified, use the client regionId, then default regionId
	request.AcceptFormat = "json"
	request.TransToAcsRequest()
	client.ProcessCommonRequest(request)

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf(response.GetHttpContentString())
}
