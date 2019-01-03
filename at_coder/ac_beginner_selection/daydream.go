package main

import (
	"fmt"
)

const (
	dream   = "dream"
	dreamer = "dreamer"
	erase   = "erase"
	eraser  = "eraser"
)

func main() {
	var (
		s string
	)
	fmt.Scan(&s)
	if dfs(s, "") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func dfs(str, split string) bool {
	if len(str) < len(split) {
		return false
	}
	if len(split) > 0 {
		if split != str[:len(split)] {
			return false
		}
		str = str[len(split):]
	}
	if len(str) == 0 {
		return true
	}
	if fmt.Sprintf("%c", str[0]) == "d" {
		return dfs(str, dream) || dfs(str, dreamer)
	} else if fmt.Sprintf("%c", str[0]) == "e" {
		return dfs(str, erase) || dfs(str, eraser)
	}
	return false
}
