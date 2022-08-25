package external

type Order struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	ProductId int `json:"productId" type:"int"`
	Qty int `json:"qty" type:"int"`
	ProductName string `json:"productName" type:"string"`
}

func (self *Order) getPrimaryKey() int {
	// FIXME if PrimaryKey is multi value, than change this method
	return self.Id
}
