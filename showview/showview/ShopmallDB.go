package showview
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type ShopmallDB struct{
	db *gorm.DB
}

var shopmallrepository *ShopmallDB

func ShopmallDBInit() {
	var err error
	shopmallrepository = &ShopmallDB{}
	shopmallrepository.db, err = gorm.Open(sqlite.Open("Shopmall_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	shopmallrepository.db.AutoMigrate(&Shopmall{})

}

func ShopmallRepository() *ShopmallDB {
	return shopmallrepository
}

func (self *ShopmallDB)save(entity *Shopmall) error {
	tx := self.db.Create(entity)
	if tx.Error != nil {
		tx1 := self.db.Save(entity)
		if tx1.Error != nil {
			log.Print(tx1.Error)
			return tx1.Error
		}
	}
	return nil
}

func (self *ShopmallDB)GetList() []Shopmall{
	
	entities := []Shopmall{}
	self.db.Find(&entities)

	return entities
}

func (self *ShopmallDB)GetID(id int) *Shopmall{
	entity := &Shopmall{}
	self.db.Where("id = ?", id).First(entity)

	return entity
}

func (self *ShopmallDB) Delete(entity *Shopmall) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *ShopmallDB) Update(id int, params map[string]string) error{
	entity := &Shopmall{}
	err1 := self.db.Where("id = ?", id).First(entity).Error
	if err1 != nil {
		return err1
	}else {
		update := &Shopmall{}
		ObjectMapping(update, params)

		err2 := self.db.Model(&entity).Updates(update).Error
		return err2
	}

}