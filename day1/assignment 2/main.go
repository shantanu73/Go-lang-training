package main

import "fmt"

func main()  {

	starcount := 5

	for i := 1;  i<=starcount; i++ {
		for j := 1;  j<=i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}