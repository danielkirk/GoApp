package main

import "fmt"

const (
	active   = true
	inactive = false
)

type Security struct {
	items map[string]bool
}

func activate(security *Security, key string) {
	security.items[key] = active
}

func deactivate(security *Security, key string) {
	security.items[key] = inactive
}

func checkout(security *Security) {
	for key := range security.items {
		deactivate(security, key)
	}
	fmt.Println(security.items)
}

func main() {
	security := Security{make(map[string]bool)}
	security.items["Web"] = active
	security.items["Desktop"] = active
	security.items["Tabket"] = active
	security.items["Laptop"] = active

	fmt.Println(security.items)

	secPtr := &security
	fmt.Printf("security address: %p, secPtr address: %p\n", &security, secPtr)

	deactivate(secPtr, "Laptop")
	fmt.Println(security.items)
	activate(secPtr, "Laptop")
	fmt.Println(security.items)
	checkout(secPtr)
}
