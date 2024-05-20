package basic_algorithm

import "sort"

func merge(intervals [][]int) (res [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for _, y := range intervals[1:] {
		if res[len(res)-1][1] < y[0] {
			res = append(res, y)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], y[1])
		}
	}
	return
}
