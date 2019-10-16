package math

func FactoredPrimeNumber(num int) map[int]int {
	return factor(num, 2)
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

func factor(num, pnum int) {
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
	return merge(result, factoring(num, pnum))
}
