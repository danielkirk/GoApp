package main

type Product struct {
	price int
	name  string
}

func sum(products [4]Product) int {
	total := 0
	for i := range products {
		total += products[i].price
	}
	return total
}

func printStats(products [4]Product) {
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

}
func main() {

	shoppingList := [4]Product{{12, "bacon"}, {30, "eggs"}, {4, "cheese"}}

	printStats(shoppingList)

	shoppingList[3] = Product{5, "meal"}

	printStats(shoppingList)

}
