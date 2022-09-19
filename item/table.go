package item

type Table struct {
	Id          int    `json:"id"`
	State       int8   `json:"state"`
	Table_order *Order `json:"table_order"`
}

var tables []Table

func Gentables(n int) {
	for i := 0; i < n; i++ {
		var newtable Table
		newtable.Id = i + 1
		newtable.State = 0
		newtable.Table_order = nil
		tables = append(tables, newtable)
	}
}
