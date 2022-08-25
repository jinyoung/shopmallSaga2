package main

import(
	"log"
	"delivery/delivery"
	"delivery/config"
)

func main() {
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	log.Printf("option : %v", options)
	delivery.DeliveryDBInit()
	go delivery.InitProducer(config.GetMode())
	// view 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go delivery.InitConsumer(config.GetMode())
	e := delivery.RouteInit()

	e.Logger.Fatal(e.Start(":8082"))
}
