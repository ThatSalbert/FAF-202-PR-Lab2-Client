package item

import (
	"math/rand"
)

type Cl_order_post struct {
	Client_Id int   `json:client_id`
	Orders    Order `json:orders`
}

func GenClorderpost() *Cl_order_post {
	var new_order = new(Cl_order_post)
	var order = Genorder()
	new_order.Client_Id = rand.Intn(9999-1000) + 1000
	new_order.Orders = *order
	return new_order
}
