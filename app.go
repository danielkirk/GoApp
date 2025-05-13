package main

import "fmt"

type Part string

func printParts(parts []Part) {
	var partsList string = "these parts"
	for i := 0; i < len(parts); i++ {
		partsList += ` ` + string(parts[i])
	}

	fmt.Println(partsList)
}

func main() {
	assemblyLine := []Part{"wrench", "screwdriver", "hammer"}

	printParts(assemblyLine)

	assemblyLine = append(assemblyLine, "nailgun", "pipe")

	printParts(assemblyLine)

	assemblyLine = assemblyLine[3:]

	printParts(assemblyLine)
}
