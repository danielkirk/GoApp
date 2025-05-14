package main

import "fmt"

func increment(x *int) {
	*x += 1
	fmt.Println(*x)
}

func main() {
	value := 10
	valuePtr := &value
	increment(valuePtr)
	increment(valuePtr)
	fmt.Println(value)

	i := 1
	increment(&i)
}
