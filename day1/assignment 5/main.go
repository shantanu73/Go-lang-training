package main

import "fmt"

type numlist [] int
func main()  {
	n := 10

	prime := primeList(n)
	//prime := numlist{}
	//prime = append(prime, 2)
	fmt.Println("First",n,"prime numbers are :-")
	fmt.Print(prime)
}

func primeList(n int) numlist {
	primeList := numlist{}
	stopCond := 0
	for i := 2; stopCond == 0; i++{
		if isPrime(i) {
			primeList = append(primeList, i)
			n--
		}
		if n == 0 {
			stopCond = 1
		}
	}
	return primeList
}
