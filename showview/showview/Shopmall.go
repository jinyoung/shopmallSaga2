package showview

type Shopmall struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	OrderId int `json:"orderId" type:"int"`
	ProductId int `json:"productId" type:"int"`
	Qty int `json:"qty" type:"int"`
	DeliveryId int `json:"deliveryId" type:"int"`
}