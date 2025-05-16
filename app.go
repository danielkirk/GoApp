package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	currentLot := ParkingLot{make([]Space, 10)}
	fmt.Println(currentLot)
	currentLot.occupySpace(1)
	fmt.Println(currentLot)

	occupySpace(&currentLot, 4)
	fmt.Println(currentLot)
	currentLot.vacateSpace(1)
	fmt.Println(currentLot)
}
