package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		3
		100000
		30000
		20000
	*/
	var n, p int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p)
		if p%2 != 0 {
			fmt.Println("first")
			os.Exit(0)
		}
	}
	fmt.Println("second")
}
