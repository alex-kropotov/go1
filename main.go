package main

import "fmt"

func main() {
	var input = []int{1, 8, 5, 3, 6, 12}
	sortInsertion(input)
	fmt.Println(input)
}

func sortInsertion(input []int) {
	var curVal int
	for pos, val := range input[1:] {
		curVal = val
		pos++
		for (pos > 0) && (input[pos-1] > curVal) {
			input[pos] = input[pos-1]
			pos--
		}
		input[pos] = curVal
	}
}