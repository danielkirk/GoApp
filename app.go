package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {
	serverList := []string{"darkstar", "aiur", "omicron", "w359", "base"}

	serverStatus := make(map[string]int)

	for _, server := range serverList {
		serverStatus[server] = Online
	}

	fmt.Println(serverStatus)
}
