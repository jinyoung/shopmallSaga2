package showview

import (
	"time"
)

type DeliveryStarted struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"`
	OrderId int `json:"orderId" type:"int"`
	ProductId int `json:"productId" type:"int"`
	ProductName string `json:"productName" type:"string"`
}

func NewDeliveryStarted() *DeliveryStarted{
	event := &DeliveryStarted{EventType:"DeliveryStarted", TimeStamp:time.Now().String()}
	return event
}
