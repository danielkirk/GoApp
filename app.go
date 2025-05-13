package main

import "fmt"

func main() {
	slice := []string{"hey", "how's", "it", "going"}

	for i, element := range slice {
		fmt.Println(i, element)
		for _, ch := range element {
			fmt.Printf("      %q\n", ch)
		}
	}
}
