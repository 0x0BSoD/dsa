// 217. Contains Duplicate https://neetcode.io/problems/duplicate-integer?list=neetcode150
package main

import "fmt"

func containsDuplicate(input []int) bool {
	n := len(input)
	sortedInput := make(map[int]int, n)
	cycles := 0

	// Insert Sorting
	for i := range n {
		j := i - 1
		key := input[i]
		fmt.Printf(" i: %v | j: %v | key: %v\n", i, j, key)

		for j > 0 && input[j] > key {
			fmt.Printf("\tinput[j]:%v bigger than key:%v\n", input[j], key)
			fmt.Printf("\tinserting %v to result[%v]\n", input[j], j+1)
			sortedInput[j+1] = input[j]
			j = j - 1
			cycles += 1
		}
		cycles += 1
		fmt.Printf("inserting %v to result[%v]\n", key, j+1)
		sortedInput[j+1] = key
	}

	fmt.Printf("Loops: 2, Cycles: %v\n", cycles)

	// Check duplicates
	j := 0
	for j < len(sortedInput) {
		first := sortedInput[j]
		second := sortedInput[j+1]

		fmt.Printf("First: %v, Second: %v\n", first, second)

		if first == second {
			return true
		}

		j += 1
	}

	return false
}
