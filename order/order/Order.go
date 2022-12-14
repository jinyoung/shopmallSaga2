package order

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
)
type Order struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	ProductId int `json:"productId" type:"int"`
	Qty int `json:"qty" type:"int"`
	ProductName string `json:"productName" type:"string"`
}

func (self *Order) AfterCreate(tx *gorm.DB) (err error){
	orderPlaced := NewOrderPlaced()
	model.Copy(orderPlaced, self)

	Publish(orderPlaced)
	
	
	return nil
}
func (self *Order) BeforeCreate(tx *gorm.DB) (err error){
	return nil
}
func (self *Order) BeforeDelete(tx *gorm.DB) (err error){
	orderCancelled := NewOrderCancelled()
	model.Copy(orderCancelled, self)

	Publish(orderCancelled)
	
	
	return nil
}
