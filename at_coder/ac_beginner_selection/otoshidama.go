package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		x, y int
	)
	fmt.Scanf("%d %d\n", &x, &y)
	for i := 0; i <= x; i++ {
		for j := 0; j <= x-i; j++ {
			if i*10000+j*5000+(x-i-j)*1000 == y {
				fmt.Printf("%d %d %d\n", i, j, x-i-j)
				os.Exit(0)
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
