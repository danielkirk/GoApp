package main

import "fmt"

// This is a simple Go program that prints "Hello World" to the console.
func main() {
	var favColor = `Green`

	fmt.Println(favColor)

	birthYear, age := 1996, 29

	fmt.Println(birthYear, age)
	var (
		firstInitial = `d`
		lastInitial  = `k`
	)

	fmt.Println(firstInitial, lastInitial)

	var ageDays int

	ageDays = multiplyAge(365, age)

	fmt.Println(ageDays)

	a, b, c := multiReturn()

	fmt.Println(a, b, c)
}

func multiplyAge(firstValue, secondValue int) int {
	return firstValue * secondValue
}

func multiReturn() (int, int, int) {
	return 1, 2, 3
}
