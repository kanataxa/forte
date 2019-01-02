package main

import (
"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	t, x, y := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d\n", &t[i+1], &x[i+1], &y[i+1])
	}
	ok := true
	for i := 0; i < n; i++ {
		dt := t[i+1] - t[i]
		dx := x[i+1] - x[i]
		dy := y[i+1] - y[i]
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		if dx+dy > dt {
			ok = false
			break
		}
		if (dx+dy)%2 != dt%2 {
			ok = false
			break
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
