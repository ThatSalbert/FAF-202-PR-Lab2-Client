package main

import (
	"dining_hall/item"
)

var ntables int = 5
var nwaiters int = 3

func startSim() {
	item.Gentables(ntables)
	item.Genwaiter(nwaiters)
}

func main() {
	startSim()
}
