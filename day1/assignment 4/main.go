package main

import "fmt"

func main()  {

	n := 5

	fact := factorial(n)

	fmt.Print(fact)
}

func factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n*factorial(n-1)
}