package item

import "time"

type Dh_order_post struct {
	Order_id     int     `json:"order_id"`
	Table_id     int     `json:"table_id"`
	Waiter_id    int     `json:"waiter_id"`
	Items        []int   `json:"items"`
	Priority     int     `json:"priority"`
	Max_wait     float32 `json:"max_wait"`
	Pick_up_time int64   `json:"pick_up_time"`
}

func Genorderpost(order Order, waiter Waiter, table Table) *Dh_order_post {
	var order_post = new(Dh_order_post)
	order_post.Order_id = order.Id
	order_post.Table_id = table.Id
	order_post.Order_id = waiter.Id
	order_post.Items = order.Items
	order_post.Priority = order.Priority
	order_post.Max_wait = float32(order.Max_wait)
	order_post.Pick_up_time = time.Now().Unix()

	return order_post
}
