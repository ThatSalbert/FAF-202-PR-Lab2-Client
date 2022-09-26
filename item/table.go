package item

import (
	"fmt"
	"strconv"
	"time"
)

type Table struct {
	Id          int    `json:"id"`
	State       int8   `json:"state"`
	Table_order *Order `json:"table_order"`
}

var OrderChannel = make(chan Order)

var Tables []Table

func Gentables(n int) {
	for i := 0; i < n; i++ {
		var newtable Table
		newtable.Id = i + 1
		newtable.State = 0
		newtable.Table_order = nil
		Tables = append(Tables, newtable)
	}
}

func TablesStartOrder() {
	for _, table := range Tables {
		go table.HandleOrder()
	}
}

func (table *Table) HandleOrder() {
	for {
		if table.State == 0 {
			table.State = 1
			table.Table_order = Genorder()
			fmt.Println("Order generated for table " + strconv.Itoa(table.Id))
			time.Sleep(2 * time.Second)
			OrderChannel <- *table.Table_order
		}
	}
}
