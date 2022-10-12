package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"client/item"

	"github.com/gorilla/mux"
)

func startSim() {
	GenRandomOrder()
}

func GenRandomOrder() {
	go RandomOrder()
}

func RandomOrder() {
	time.Sleep(time.Second * 10)
	for {
		n := rand.Intn(10)
		if n <= 4 {
			ordertobesent, err := json.Marshal(item.GenClorderpost())
			fmt.Println("Client Generated Order: " + string(ordertobesent))
			response, err := http.Post("http://food-order:7000/order", "application/json", bytes.NewBuffer(ordertobesent))
			fmt.Println("Order sent to Food Ordering Service.")
			if err != nil {
				fmt.Print("Could not make POST request to the Food Ordering Service.")
			}
			defer response.Body.Close()
			time.Sleep(time.Second * 4)
		} else {
			fmt.Println("Client didn't generate any order.")
			time.Sleep(time.Second * 4)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := mux.NewRouter()

	//router.HandleFunc("/order", Orders).Methods("GET")
	//router.HandleFunc("/order", Posts).Methods("POST")

	startSim()

	http.ListenAndServe(":6000", router)
}

// func Posts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var order item.K_order_post
// 	_ = json.NewDecoder(r.Body).Decode(&order)
// 	item.ReceivedOrder = append(item.ReceivedOrder, order)
// 	json.NewEncoder(w).Encode(&order)

// 	ordertobesent, err := json.Marshal(order)
// 	if err != nil {
// 		fmt.Print("No order.")
// 	}
// 	fmt.Println("Food Ordering Service got the order: " + string(ordertobesent))
// }

// func Orders(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	json.NewEncoder(w).Encode(item.ReceivedOrder)
// }
