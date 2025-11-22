// 217. Contains Duplicate https://neetcode.io/problems/duplicate-integer?list=neetcode150
package main

func containsDuplicate(input []int) bool {
	checked := make(map[int]bool, len(input))

	for _, i := range input {
		if checked[i] {
			return true
		}
		checked[i] = true
	}

	return false
}
