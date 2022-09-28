package main


//@GetOrders godoc
//@summary Get details all orders
//@Tags Orders
//@Accept json
//@Produce json
//@Success 200 {array} Order
//@Router /orders [get]
import (
	"encoding/json"
	
)

func getOrders(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content.Type","application/json")
    json.NewDecoder(w),Encode(orders)
}	