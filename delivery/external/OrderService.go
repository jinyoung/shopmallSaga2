package external 

import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

var client = resty.New()

func CancelOrder(entity *Order) *resty.Response{
	target := "https://order:8080/orders/"+strconv.Itoa(entity.getPrimaryKey())

	resp, _ := client.R().
		Delete(target)

	return resp
}

