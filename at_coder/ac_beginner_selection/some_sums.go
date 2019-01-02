package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d\n", &n, &a, &b)

	var sum int
	for i := 1; i <= n; i++ {
		total := score(i)
		if total >= a && total <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}

func score(num int) int {
	var total int
	for num > 0 {
		total += num % 10
		num /= 10
	}
	return total
}
