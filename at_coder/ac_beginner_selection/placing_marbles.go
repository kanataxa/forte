package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var cnt int
	for _, r := range s {
		if fmt.Sprintf("%c", r) == "1" {
			cnt++
		}
	}
	fmt.Println(cnt)
}
