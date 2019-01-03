package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n int
	)
	fmt.Scan(&n)
	a := make(number, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Sort(a)

	var score int

	for i, n := range a {
		if i%2 == 0 {
			score += n
		} else {
			score -= n
		}
	}

	fmt.Println(score)
}

type number []int

func (n number) Len() int {
	return len(n)
}

func (n number) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n number) Less(i, j int) bool {
	return n[i] > n[j]
}
