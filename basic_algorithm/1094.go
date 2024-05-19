package basic_algorithm

func carPooling(trips [][]int, capacity int) bool {
	res := make([]int, 1001)
	diff := make([]int, 1001)
	diff[0] = res[0]
	for _, record := range trips {
		diff[record[1]] += record[0]
		if record[2] < 1001 {
			diff[record[2]] -= record[0]
		}
	}
	res[0] = diff[0]
	if res[0] > capacity {
		return false
	}
	for i := 1; i < 1001; i++ {
		res[i] = res[i-1] + diff[i]
		if res[i] > capacity {
			return false
		}
	}
	return true
}
