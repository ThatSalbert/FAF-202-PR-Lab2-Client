package item

type waiter struct {
	id           int
	state        int8
	picked_order *order
}

var waiters []waiter

func Genwaiter(n int) {
	for i := 0; i < n; i++ {
		var newwaiter waiter
		newwaiter.id = i + 1
		newwaiter.state = 0
		newwaiter.picked_order = nil
		waiters = append(waiters, newwaiter)
	}
}
