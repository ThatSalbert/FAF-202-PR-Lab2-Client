package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"dining_hall/item"

	"github.com/gorilla/mux"
)

var ntables int = 5
var nwaiters int = 3

const simtime = 350

func startSim() {
	item.Gentables(ntables)
	item.Genwaiters(nwaiters)
	item.WaitersStartWork()
	item.TablesStartOrder()
}

// func GenRandomOrder() {
// 	for {
// 		time.Sleep(4 * time.Second)
// 		go RandomOrder()
// 	}
// }

// func RandomOrder() {
// 	n := rand.Intn(10)
// 	if n <= 4 {
// 		fmt.Println("Order Generated.")
// 		item.Order_list = append(item.Order_list, item.Genorder(rand.Intn(9999-1000)+1000))
// 	} else {
// 		fmt.Println("Didn't generate any order.")
// 	}
// }

var Received_order []item.Order

func main() {
	rand.Seed(time.Now().UnixNano())
	startSim()

	router := mux.NewRouter()

	router.HandleFunc("/distribution", Orders).Methods("GET")
	//PostOrder()

	http.ListenAndServe(":8080", router)
}

func Orders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(item.Order_list)
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
