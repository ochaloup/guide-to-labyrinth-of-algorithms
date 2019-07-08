package main

import (
	"fmt"

	"github.com/ochaloup/guide-to-labyrinth-of-algorithms/chapter1/sections"
)

var sequence1 = []int{1, -2, 4, 5, -1, -5, 2, 7}
var sequence2 = []int{1, -2, 4, 5, -1, -5, 2, 7, -1, 111}

func main() {
	var result int
	var left, right int

	// run task 0
	// sequence1 = append(sequence1, 1)
	// result = sections.MaxSum3(sequence1)

	// run task 1
	result, left, right = sections.MaxSum3WithIndex(sequence2)

	fmt.Printf("Result is: %v, index left: %v, index right: %v\n", result, left, right)
}
