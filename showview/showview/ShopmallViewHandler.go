package showview

import (
	"github.com/mitchellh/mapstructure"
	"log"
)

func whenOrderPlaced_then_CREATE_1 (inputEvent map[string]interface{}) {

	orderPlaced := NewOrderPlaced()
	mapstructure.Decode(inputEvent, &orderPlaced)

	shopmall := &Shopmall{}
	shopmall.OrderId = orderPlaced.Id
	shopmall.ProductId = orderPlaced.ProductId
	shopmall.Qty = orderPlaced.Qty

	// view 레파지 토리에 save
	repository := ShopmallRepository()
	err := repository.save(shopmall)
	if err != nil {
		// TODO error control
		log.Printf("Create error: %v \n", err)
	}
}

func whenDeliveryStarted_then_UPDATE_1 (inputEvent map[string]interface{}) {

	deliveryStarted := NewDeliveryStarted()
	mapstructure.Decode(inputEvent,&deliveryStarted)

	var shopmalls []Shopmall
	repository := ShopmallRepository()
	// FIXME geom lib define snake_case as column name (eg: user_id), so if your query column is 'userId' then change 'user_id'
	err := repository.db.Where("orderId = ?", deliveryStarted.OrderId).Find(&shopmalls).Error
	if err != nil {
		// TODO error control
		log.Printf("Select error: %v \n", err)
	}
	for _, viewEntity := range shopmalls {
		viewEntity.DeliveryId = deliveryStarted.Id
		err1 := repository.db.Updates(viewEntity).Error
		if err1 != nil {
			// TODO error control
			log.Printf("Update error: %v \n", err1)
		}
	}
}

func whenOrderCancelled_then_DELETE_1 (inputEvent map[string]interface{}) {
	orderCancelled := NewOrderCancelled()
	mapstructure.Decode(inputEvent,&orderCancelled)

	var shopmalls []Shopmall
	repository := ShopmallRepository()
	// FIXME geom lib define snake_case as column name (eg: user_id), so if your query column is 'userId' then change 'user_id'
	err := repository.db.Where("orderId = ?", orderCancelled.Id).Find(&shopmalls).Error
	if err != nil {
		// TODO error control
		log.Printf("Select error: %v \n", err)
	}
	for _, viewEntity := range shopmalls {
		err1 := repository.db.Delete(viewEntity).Error
		if err1 != nil {
			// TODO error control
			log.Printf("Delete error: %v \n", err1)
		}
	}
}
