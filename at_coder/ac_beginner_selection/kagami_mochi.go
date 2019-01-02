package main

import (
	"fmt"
)

func main() {
	var (
		n int
	)
	fmt.Scan(&n)
	d := make(map[int]struct{})
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		d[tmp] = struct{}{}
	}

	fmt.Println(len(d))
}
