package main

func findMax(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	}
	return c
}
