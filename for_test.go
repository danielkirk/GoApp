package main

import "testing"

var myArray = [...]int{1, 2, 3}

func BenchmarkIterator(b *testing.B) {
	sum := 0
	for b.Loop() {
		for i := range myArray {
			sum += myArray[i]
		}
	}
}

func BenchmarkRange(b *testing.B) {
	sum := 0
	for b.Loop() {
		for i := 0; i < len(myArray); i++ {
			sum += myArray[i]
		}
	}
}

// func forTest() {

// 	BenchmarkIterator()
// 	BenchmarkRange()

// }
