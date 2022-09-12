package item

import (
	"math/rand"
	"time"
)

type order struct {
	id       int     `json:"id"`
	items    []int   `json:"items"`
	priority int     `json:"priority"`
	max_wait float32 `json:"max_wait"`
}

func Genorder(id int) order {
	rand.Seed(time.Now().UnixNano())
	var nitems int = rand.Intn(6-1) + 1
	var table_order order
	table_order.id = id
	table_order.items = Genlist(nitems)
	table_order.priority = rand.Intn(5-1) + 1
	table_order.max_wait = float32(Get_max_wait(table_order.items)) * 1.3
	return table_order
}