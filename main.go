package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"dining_hall/item"

	"github.com/gorilla/mux"
)

var ntables int = 5
var nwaiters int = 3

const simtime = 350

func startSim() {
	//item.Gentables(ntables)
	//item.Genwaiters(nwaiters)
	GenRandomOrder()
}

func GenRandomOrder() {
	go RandomOrder()
}

func RandomOrder() {
	for {
		n := rand.Intn(10)
		if n <= 4 {
			ordertobesent, err := json.Marshal(item.Genorder())
			fmt.Println("Order Generated: " + string(ordertobesent))
			response, err := http.Post("http://kitchen:8000/order", "application/json", bytes.NewBuffer(ordertobesent))
			fmt.Println("Order sent to kitchen.")
			if err != nil {
				fmt.Print("Could not make POST request to the kitchen.")
			}
			defer response.Body.Close()
			time.Sleep(time.Second * 2)
		} else {
			fmt.Println("Didn't generate any order.")
			time.Sleep(time.Second * 2)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := mux.NewRouter()

	router.HandleFunc("/distribution", Orders).Methods("GET")
	router.HandleFunc("/distribution", Posts).Methods("POST")

	startSim()

	http.ListenAndServe(":8080", router)
}

func Posts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order item.K_order_post
	_ = json.NewDecoder(r.Body).Decode(&order)
	item.ReceivedOrder = append(item.ReceivedOrder, order)
	json.NewEncoder(w).Encode(&order)

	ordertobesent, err := json.Marshal(order)
	if err != nil {
		fmt.Print("No order.")
	}
	fmt.Println("Dining Hall got the order: " + string(ordertobesent))
}

func Orders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(item.ReceivedOrder)
}
