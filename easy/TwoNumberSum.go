package main

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	data := map[int]bool{}
	for _, i := range(array) {
		needed := target - i
		if data[needed] {
			return []int{needed, i}
		} else {
			data[i] = true
		}
	}
	return []int{}
}
