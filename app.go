package main

import "fmt"

type Coordinate struct {
	X, Y int
}

func shiftBy(x, y int, coord *Coordinate) {
	coord.X += x
	coord.Y += y
}

func (coord *Coordinate) shiftBy(x, y int) {
	coord.X += x
	coord.Y += y
}

func (c Coordinate) shiftDist(other Coordinate) Coordinate {
	return Coordinate{other.X - c.X, other.Y - c.Y}
}

func main() {
	coord := Coordinate{5, 5}
	fmt.Println(coord.X, coord.Y)
	shiftBy(2, 3, &coord)
	fmt.Println(coord.X, coord.Y)
	coord.shiftBy(1, 2)
	fmt.Println(coord.X, coord.Y)

	first := Coordinate{12, 23}
	second := Coordinate{1, 24}
	distance := first.shiftDist(second)
	fmt.Println(distance)
}
