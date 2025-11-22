package main

import "fmt"

func main() {
	cdInputs := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}

	var cdResult []string
	for _, i := range cdInputs {
		if containsDuplicate(i) {
			cdResult = append(cdResult, fmt.Sprintf("%v => contains duplicates.", i))
		} else {
			cdResult = append(cdResult, fmt.Sprintf("%v => is not contain duplicates.", i))
		}
	}

	fmt.Println("===========")
	for _, j := range cdResult {
		fmt.Println(j)
	}
}
