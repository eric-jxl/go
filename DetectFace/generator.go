package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk"

func main() {

	client, err := sdk.NewClientWithAccessKey("REGION_ID", "ACCESS_KEY_ID", "ACCESS_KEY_SECRET")
	if err != nil {
		// Handle exceptions
		panic(err)
	}
}
