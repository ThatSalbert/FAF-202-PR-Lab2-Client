package item

type Waiter struct {
	Id           int    `json:"id"`
	State        int8   `json:"state"`
	Picked_order *Order `json:"picked_order"`
}

var Waiters []Waiter

func Genwaiters(n int) {
	for i := 0; i < n; i++ {
		var newwaiter Waiter
		newwaiter.Id = i + 1
		newwaiter.State = 0
		newwaiter.Picked_order = nil
		Waiters = append(Waiters, newwaiter)
	}
}
