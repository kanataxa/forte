package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	var count int
	for _, s := range n {
		if fmt.Sprintf("%c", s) == "2" {
			count++
		}
	}
	fmt.Println(count)
}
