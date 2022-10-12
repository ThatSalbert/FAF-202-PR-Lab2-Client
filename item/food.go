package item

import (
	"math/rand"
)

type Food struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Preparation_time  int    `json:"preparation_time"`
	Complexity        int    `json:"complexity"`
	Cooking_apparatus string `json:"cooking_apparatus"`
}

var pizza = Food{Id: 1, Name: "Pizza", Preparation_time: 20, Complexity: 2, Cooking_apparatus: "oven"}
var salad = Food{Id: 2, Name: "Salad", Preparation_time: 10, Complexity: 1, Cooking_apparatus: ""}
var zeama = Food{Id: 3, Name: "Zeama", Preparation_time: 7, Complexity: 1, Cooking_apparatus: "stove"}
var scallop = Food{Id: 4, Name: "Scallop Sashimi with Meyer Lemon Confit", Preparation_time: 32, Complexity: 3, Cooking_apparatus: ""}
var island_duck = Food{Id: 5, Name: "Island Duck with Mulberry Mustard", Preparation_time: 35, Complexity: 3, Cooking_apparatus: "oven"}
var waffles = Food{Id: 6, Name: "Waffles", Preparation_time: 10, Complexity: 1, Cooking_apparatus: "stove"}
var aubergine = Food{Id: 7, Name: "Aubergine", Preparation_time: 20, Complexity: 2, Cooking_apparatus: "oven"}
var lasagna = Food{Id: 8, Name: "Lasagna", Preparation_time: 30, Complexity: 2, Cooking_apparatus: "oven"}
var burger = Food{Id: 9, Name: "Burger", Preparation_time: 15, Complexity: 1, Cooking_apparatus: "stove"}
var gyros = Food{Id: 10, Name: "Gyros", Preparation_time: 15, Complexity: 1, Cooking_apparatus: ""}
var kebab = Food{Id: 11, Name: "Kebab", Preparation_time: 15, Complexity: 1, Cooking_apparatus: ""}
var unagi = Food{Id: 12, Name: "Unagi Maki", Preparation_time: 20, Complexity: 2, Cooking_apparatus: ""}
var tobacco = Food{Id: 13, Name: "Tobacco Chicken", Preparation_time: 30, Complexity: 2, Cooking_apparatus: "oven"}

var menu = []Food{pizza, salad, zeama, scallop, island_duck, waffles, aubergine, lasagna, burger, gyros, kebab, unagi, tobacco}

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
		if menu[itemarray[i]-1].Preparation_time > max_wait {
			max_wait = menu[itemarray[i]-1].Preparation_time
		}
	}
	return max_wait
}
