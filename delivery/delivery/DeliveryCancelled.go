package delivery

import (
	"time"
)

type DeliveryCancelled struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	OrderId int `json:"orderId" type:"int"` 
	
}

func NewDeliveryCancelled() *DeliveryCancelled{
	event := &DeliveryCancelled{EventType:"DeliveryCancelled", TimeStamp:time.Now().String()}

	return event
}
