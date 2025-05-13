package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["this thang"] = 1
	myMap["that thang"] = 2
	myMap["them thangs"] = 3

	price, found := myMap["this thang"]
	if !found {
		fmt.Println("price not found")
	} else {
		fmt.Println(price)
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	delete(myMap, "this thang")

	newPrice, newFound := myMap["this thang"]
	if !newFound {
		fmt.Println("price not found")
	} else {
		fmt.Println(newPrice)
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
