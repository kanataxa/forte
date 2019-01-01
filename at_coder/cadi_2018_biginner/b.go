package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	var n, h, w, count int
	fmt.Fscanf(in, "%d %d %d\n", &n, &h, &w)
	for i := 0; i < n; i++ {
		var height, weight int
		fmt.Fscanf(in, "%d %d\n", &height, &weight)
		if height >= h && weight >= w {
			count++
		}
	}
	fmt.Println(count)
}
