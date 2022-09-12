package item

type table struct {
	id          int
	state       int8
	table_order *order
}

var tables []table

func Gentables(n int) {
	for i := 0; i < n; i++ {
		var newtable table
		newtable.id = i + 1
		newtable.state = 0
		newtable.table_order = nil
		tables = append(tables, newtable)
	}
}
