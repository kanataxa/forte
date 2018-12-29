package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := make(map[int]int)
	for i := 1; i <= n; i ++ {
		result = merge(result, factoring(i, 2))
	}
	// 75
	// 3 * 5 * 5
	// 15 * 5
	// 25 * 3
	n_75 := num(result, 75-1)
	n_25 :=num(result,25-1)
	n_15 := num(result,15-1)
	n_5 := num(result, 5-1)
	n_3 := num(result, 3-1)
	fmt.Println(n_75+n_5*(n_5-1)*(n_3-2)/2+n_15*(n_5-1)+n_25*(n_3-1))
}
func factoring(num , pnum int) map[int]int{
	result := make(map[int]int)
	if pnum * pnum > num {
		if num != 1 {
			result[num] += 1
		}
		return result
	}

	if num % pnum == 0 {
		num /= pnum
		result[pnum] ++
	} else {
		pnum ++
	}
	 return merge(result , factoring(num, pnum))
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

func num(t map[int]int, n int) int {
	var num int
	for _, value := range t {
		if value >= n {
			num ++
		}
	}
	return num
}