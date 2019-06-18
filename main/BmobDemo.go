package main

import (
	"github.com/bmob/bmob-go-sdk"
	"log"
)

var (
	appConfig = bmob.RestConfig{"12299351c40d78301ebf709efe6d926e", "7bd844783eec87058e645d27597afc1f"}
)

type TestData struct {
	Score string
	//data  DataType
}

type MyRes struct {
	bmob.RestResponse
	bmob.ImageResponse
}


type TestDataRes struct {
	TestData
	MyRes
}

func main() {

	a := bmob.RestResponse{}
	log.Println(a)
	log.Println("****************************************")
	var respDst = TestDataRes{}

	header, err :=bmob.DoRestReq(appConfig,
		bmob.RestRequest{
			bmob.BaseReq{
				"GET",
				bmob.ApiRestURL("MainBannerBean") + "/",
				""},
			"application/json",
			nil},
		&respDst)

	if err == nil {
		log.Println(header)
		log.Println(respDst)
	} else {
		log.Panic(err)
	}

	log.Println("****************************************")

}

