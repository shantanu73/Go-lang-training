package main

import "fmt"

func main()  {

	starcount := 5
	val := 1

	for i := 1;  i<=starcount; i++ {
		for j := 1;  j<=i; j++ {
			fmt.Print(val, " ")
			val = val + 1
		}
		fmt.Println()
	}
}