package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	if s == "5" || s == "7" || s == "3" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
