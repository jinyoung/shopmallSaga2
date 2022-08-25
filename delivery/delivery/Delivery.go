package delivery

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"

	"delivery/external"
)
type Delivery struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	OrderId int `json:"orderId" type:"int"`
	ProductId int `json:"productId" type:"int"`
	ProductName string `json:"productName" type:"string"`
}

func (self *Delivery) AfterCreate(tx *gorm.DB) (err error){
	deliveryStarted := NewDeliveryStarted()
	model.Copy(deliveryStarted, self)

	Publish(deliveryStarted)
	
	
	return nil
}
func (self *Delivery) BeforeDelete(tx *gorm.DB) (err error){
	deliveryCancelled := NewDeliveryCancelled()
	model.Copy(deliveryCancelled, self)

	Publish(deliveryCancelled)
	
	
	order := &external.Order{}
	resp := external.CancelOrder(order)
	fmt.Println(resp)
	return nil
}
