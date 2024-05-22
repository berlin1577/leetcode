package basic_algorithm

import "sort"

func findWinners(matches [][]int) [][]int {
	res := map[int]int{}
	ans := make([][]int, 2)
	for _, match := range matches {
		if _, ok := res[match[0]]; !ok {
			res[match[0]] = 0
		}
		res[match[1]]++
	}
	for key, value := range res {
		if value < 2 {
			ans[value] = append(ans[value], key)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	return ans
}
