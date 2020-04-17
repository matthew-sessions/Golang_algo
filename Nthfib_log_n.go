package main

func GetNthFib(n int) int {
	// Write your code here.
	if n == 1 {
		return 0
	} else if n == 0 {
		return 0
	} else {
		f1 := 0
		f2 := 1
		counter := 2
		for counter < n {
			newval := f1 + f2
			f1 = f2
			f2 = newval
			counter++
		}
		return f2
	}
}
