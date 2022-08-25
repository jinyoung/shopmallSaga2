package delivery
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type DeliveryDB struct{
	db *gorm.DB
}

var deliveryrepository *DeliveryDB

func DeliveryDBInit() {
	var err error
	deliveryrepository = &DeliveryDB{}
	deliveryrepository.db, err = gorm.Open(sqlite.Open("Delivery_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	deliveryrepository.db.AutoMigrate(&Delivery{})

}

func DeliveryRepository() *DeliveryDB {
	return deliveryrepository
}

func (self *DeliveryDB)save(entity interface{}) error {
	
	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *DeliveryDB)GetList() []Delivery{
	
	entities := []Delivery{}
	self.db.Find(&entities)

	return entities
}

func (self *DeliveryDB)GetID(id int) *Delivery{
	entity := &Delivery{}
	self.db.Where("id = ?", id).First(entity)

	return entity
}

func (self *DeliveryDB) Delete(entity *Delivery) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *DeliveryDB) Update(id int, params map[string]string) error{
	entity := &Delivery{}
	err1 := self.db.Where("id = ?", id).First(entity).Error
	if err1 != nil {
		return err1
	}else {
		update := &Delivery{}
		ObjectMapping(update, params)

		err2 := self.db.Model(&entity).Updates(update).Error
		return err2
	}

}