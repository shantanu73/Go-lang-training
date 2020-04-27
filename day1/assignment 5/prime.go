package main

func isPrime(n int) bool {
	val := true
	for i := 2; i <= sqrtInt(n); i++ {
		if n%i == 0 {
			val = false
		}
	}
	return val
}
