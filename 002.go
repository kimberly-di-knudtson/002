package main

import (
	"fmt"
)

var input = []int{5, 2, 1, 3}

func main() {
	fmt.Print("Input Array: ")
	fmt.Println(input)
	output := withDivision()
	fmt.Print("Output Array, divide out: ")
	fmt.Println(output)

	output = withoutDivision()
	fmt.Print("Output Array, products only: ")
	fmt.Println(output)
}

func withDivision() []int {
	prod := 1
	count_zeros := 0 //handle 0s - if > 1 zero values, all zero output
	zero_index := -1 //if 1 zero, capture its index only that is non-zero in output

	//get total product of all elements
	for i, v := range input {
		if v == 0 {
			count_zeros++
			zero_index = i
		} else {
			prod = prod * v
		}
		if count_zeros > 1 { //no need to continue, output is all zero
			break
		}
	}

	output := make([]int, len(input)) //initialize output - all zeros
	if count_zeros == 0 {             //no zeros in input - populate output values
		for i, v := range input {
			output[i] = prod / v
		}
	} else if count_zeros == 1 { //one zero in input - its is the only index with a value
		output[zero_index] = prod
	}
	return output
}

func withoutDivision() []int {
	after := make([]int, len(input))
	prod := 1
	for i := len(input) - 1; i >= 0; i-- { //start at last index
		after[i] = prod
		prod = prod * input[i]
	}
	fmt.Print("After array: ")
	fmt.Println(after)

	//before := make([]int, len(input)) -- save memory, don't really need a before array
	output := make([]int, len(input))
	prod = 1
	for i := 0; i < len(input); i++ {
		//before[i] = prod
		//output[i] = before[i] * after[i]
		output[i] = prod * after[i]
		prod = prod * input[i]
	}
	//fmt.Print("Before array: ")
	//fmt.Println(before)

	return output
}
