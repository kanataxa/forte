package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scanf("%d\n%d\n%d\n%d\n", &a, &b, &c, &x)
	var cnt int
	for i := 0; i <= a; i++ {
		changeA := x
		result := 500 * i
		if changeA == result {
			cnt++
			continue
		}
		if changeA > result {
			changeA -= result
			for i := 0; i <= b; i++ {
				changeB := changeA
				result := 100 * i
				if changeB == result {
					cnt++
					continue
				}
				if changeB > result {
					changeB -= result
					for i := 0; i <= c; i++ {
						changeC := changeB
						result := 50 * i
						if changeC == result {
							cnt++
							continue
						}
					}
				}
			}
		}
	}
	fmt.Println(cnt)
}
