package item

import (
	"math/rand"
	"time"
)

type Order struct {
	Restaurant_Id int     `json:"id"`
	Items         []int   `json:"items"`
	Priority      int     `json:"priority"`
	Max_wait      float32 `json:"max_wait"`
	Created_Time  int64   `json:"created_time"`
}

func Genorder() *Order {
	var client_order = new(Order)
	rand.Seed(time.Now().UnixNano())
	var nitems int = rand.Intn(6-1) + 1
	client_order.Restaurant_Id = 1
	client_order.Items = Genlist(nitems)
	client_order.Priority = rand.Intn(5-1) + 1
	client_order.Max_wait = float32(Get_max_wait(client_order.Items)) * 1.3
	client_order.Created_Time = time.Now().Unix()
	return client_order
}
