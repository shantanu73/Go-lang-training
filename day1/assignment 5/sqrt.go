package main

func sqrtInt(n int) int {
	var sqrt int
	for i:=1; i*i <= n; i++ {
		sqrt = i
	}
	return sqrt
}