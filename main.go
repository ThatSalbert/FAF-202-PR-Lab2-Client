package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"dining_hall/item"

	"github.com/gorilla/mux"
)

var ntables int = 5
var nwaiters int = 3

// func startSim() {
// 	item.Gentables(ntables)
// 	item.Genwaiter(nwaiters)
// }

var order_list []item.Order
var received_order []item.Order

func main() {

	var n = 5

	for i := 1; i <= n; i++ {
		order_list = append(order_list, item.Genorder(i))
	}

	item.Genwaiter(3)

	fmt.Println(order_list)
	fmt.Println(item.Waiters)

	//	startSim()
	router := mux.NewRouter()

	router.HandleFunc("/distribution", Orders).Methods("GET")
	PostOrder()

	log.Println("Listening...")
	http.ListenAndServe(":8080", router)
}

func Orders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(order_list)
}

func PostOrder() {
	const url = "http://kitchen:8000/order"

	requestBody := strings.NewReader(`
		{
			"order_id": 6,
			"table_id": 3,
			"waiter_id": 3,
			"items": [3, 4, 5],
			"priority": 3,
			"max_wait": 35,
			"pick_up_time": 4829
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
