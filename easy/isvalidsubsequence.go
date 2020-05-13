package main

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	pointer := 0
	for _, i := range(array) {
		if i == sequence[pointer] {
			pointer += 1
			if pointer == len(sequence) {
				return true
			}
		}
		
	}
	return false
}