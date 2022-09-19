package item

import (
	"math/rand"
	"time"
)

type Order struct {
	Id       int     `json:"id"`
	Items    []int   `json:"items"`
	Priority int     `json:"priority"`
	Max_wait float32 `json:"max_wait"`
}

func Genorder(id int) Order {
	rand.Seed(time.Now().UnixNano())
	var nitems int = rand.Intn(6-1) + 1
	var table_order Order
	table_order.Id = id
	table_order.Items = Genlist(nitems)
	table_order.Priority = rand.Intn(5-1) + 1
	table_order.Max_wait = float32(Get_max_wait(table_order.Items)) * 1.3
	return table_order
}
