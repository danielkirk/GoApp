package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	const (
		diceSides = 12
		diceRolls = 2
		dice      = 4
	)

	for r := 1; r <= diceRolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(diceSides)
			sum += rolled
			fmt.Println("Roll #", r, ", die #", d, ":", rolled)
		}

		fmt.Println("Total Rolled", sum)
		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("Snake Eyes")
		case sum == 7:
			fmt.Println("Lucky Seven!")
		}
	}
}
