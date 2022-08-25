package main

import(
	"log"
	"showview/showview"
	"showview/config"
)

func main() {
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	log.Printf("option : %v", options)
	showview.ShopmallDBInit()
	go showview.InitProducer(config.GetMode())
	// policy 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go showview.InitConsumer(config.GetMode())
	e := showview.RouteInit()

	e.Logger.Fatal(e.Start(":8083"))
}
