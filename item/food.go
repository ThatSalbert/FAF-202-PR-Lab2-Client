package item

import (
	"math/rand"
)

type food struct {
	id                int    `json:"id"`
	name              string `json:"name"`
	preparation_time  int    `json:"preparation-time"`
	complexity        int    `json:"complexity"`
	cooking_apparatus string `json:"cooking-aparatus"`
}

var pizza = food{id: 1, name: "pizza", preparation_time: 20, complexity: 2, cooking_apparatus: "oven"}
var salad = food{id: 2, name: "salad", preparation_time: 10, complexity: 1, cooking_apparatus: ""}
var zeama = food{id: 3, name: "zeama", preparation_time: 7, complexity: 1, cooking_apparatus: "stove"}
var scallop = food{id: 4, name: "Scallop Sashimi with Meyer Lemon Confit", preparation_time: 32, complexity: 3, cooking_apparatus: ""}
var island_duck = food{id: 5, name: "Island Duck with Mulberry Mustard", preparation_time: 35, complexity: 3, cooking_apparatus: "oven"}
var waffles = food{id: 6, name: "Waffles", preparation_time: 10, complexity: 1, cooking_apparatus: "stove"}
var aubergine = food{id: 7, name: "Aubergine", preparation_time: 20, complexity: 2, cooking_apparatus: "oven"}
var lasagna = food{id: 8, name: "Lasagna", preparation_time: 30, complexity: 2, cooking_apparatus: "oven"}
var burger = food{id: 9, name: "Burger", preparation_time: 15, complexity: 1, cooking_apparatus: "stove"}
var gyros = food{id: 10, name: "Gyros", preparation_time: 15, complexity: 1, cooking_apparatus: ""}
var kebab = food{id: 11, name: "Kebab", preparation_time: 15, complexity: 1, cooking_apparatus: ""}
var unagi = food{id: 12, name: "Unagi Maki", preparation_time: 20, complexity: 2, cooking_apparatus: ""}
var tobacco = food{id: 13, name: "Tobacco Chicken", preparation_time: 30, complexity: 2, cooking_apparatus: "oven"}

var menu = []food{pizza, salad, zeama, scallop, island_duck, waffles, aubergine, lasagna, burger, gyros, kebab, unagi, tobacco}

func Genlist(nitems int) []int {
	var itemarray []int = make([]int, nitems)
	for i := 0; i < nitems; i++ {
		itemarray[i] = rand.Intn(13-1) + 1
	}
	return itemarray
}

func Get_max_wait(itemarray []int) int {
	var max_wait int = -1000
	for i := 0; i < len(itemarray); i++ {
		if menu[itemarray[i]-1].preparation_time > max_wait {
			max_wait = menu[itemarray[i]-1].preparation_time
		}
	}
	return max_wait
}
