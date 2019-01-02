package main

import "fmt"

func main() {
	var (
		a, b, c int
		str     string
	)
	fmt.Scanf("%d\n%d %d\n%s", &a, &b, &c, &str)
	fmt.Printf("%d %s\n", a+b+c, str)
}
