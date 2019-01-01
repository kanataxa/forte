package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, p int
	fmt.Fscanf(in, "%d %d\n", &n, &p)

	if n == 1 {
		fmt.Println(p)
		return
	}
	result := 1
	for value, count := range Factoring(p, 2) {
		result *= int(math.Pow(float64(value), float64(count/n)))
	}
	fmt.Println(result)
}

func Factoring(num, pnum int) map[int]int {
	result := make(map[int]int)
	if pnum*pnum > num {
		if num != 1 {
			result[num] += 1
		}
		return result
	}

	if num%pnum == 0 {
		num /= pnum
		result[pnum]++
	} else {
		pnum++
	}
	return merge(result, Factoring(num, pnum))
}

func merge(m1, m2 map[int]int) map[int]int {
	ans := map[int]int{}
	for k, v := range m1 {
		ans[k] += v
	}
	for k, v := range m2 {
		ans[k] += v
	}
	return ans
}
