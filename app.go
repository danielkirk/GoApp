package main

import "fmt"

type Product struct {
	price int
	name  string
}

func sum(products []Product) int {
	total := 0
	for i := range products {
		total += products[i].price
	}
	return total
}

func printStats(products []Product) {
	var cost, totalItems int

	for i := range products {
		cost += products[i].price
		if products[i].name != "" {
			totalItems++
		}
	}

	println(cost)
	println(totalItems)
	println(products[totalItems-1].name)
	println(len(products), cap(products))

}
func main() {

	shoppingList := []Product{{12, "bacon"}, {30, "eggs"}, {4, "cheese"}, {5, "meal"}}
	newList := []Product{}

	//omits the first index item and returns the rest
	printStats(shoppingList[1:])

	newList = append(newList, shoppingList...)

	//omits the first two indexed items and returns the rest
	printStats(newList)
	newList = append(newList, shoppingList...)
	//returns the entire array
	printStats(newList)
	newList = append(newList, shoppingList...)
	//returns the entire array without explicitly stating it
	printStats(newList)

	//since I know the array will be 4 I can preallocate the slice.
	preAllocatedSlice := make([]Product, 0, len(shoppingList)*4)
	preAllocatedSlice = append(preAllocatedSlice, shoppingList...)
	fmt.Printf("d: %p, len: %d, cap: %d\n", preAllocatedSlice, len(preAllocatedSlice), cap(preAllocatedSlice))

	preAllocatedSlice = append(preAllocatedSlice, shoppingList...)
	fmt.Printf("d: %p, len: %d, cap: %d\n", preAllocatedSlice, len(preAllocatedSlice), cap(preAllocatedSlice))
	preAllocatedSlice = append(preAllocatedSlice, shoppingList...)
	fmt.Printf("d: %p, len: %d, cap: %d\n", preAllocatedSlice, len(preAllocatedSlice), cap(preAllocatedSlice))
	preAllocatedSlice = append(preAllocatedSlice, shoppingList...)
	fmt.Printf("d: %p, len: %d, cap: %d\n", preAllocatedSlice, len(preAllocatedSlice), cap(preAllocatedSlice))

}
