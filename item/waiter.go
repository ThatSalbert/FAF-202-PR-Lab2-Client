package item

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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

func WaitersStartWork() {
	for _, waiter := range Waiters {
		go waiter.ProcessOrder()
	}
}

func (waiter *Waiter) ProcessOrder() {
	for {
		select {
		case recievedOrder := <-OrderChannel:
			defer close(OrderChannel)
			ordertobesent, err := json.Marshal(recievedOrder)
			fmt.Println("Waiter " + strconv.Itoa(waiter.Id) + " recieved order " + string(ordertobesent))
			response, err := http.Post("http://kitchen:8000/order", "application/json", bytes.NewBuffer(ordertobesent))
			if err != nil {
				fmt.Print("Could not make POST request to the kitchen.")
			}
			defer response.Body.Close()
		}
	}
}
