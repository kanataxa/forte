package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var cnt int
	for {
		for idx, num := range a {
			if num%2 != 0 {
				fmt.Println(cnt)
				os.Exit(0)
			} else {
				a[idx] = num / 2
			}
		}
		cnt++
	}
}
