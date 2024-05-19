package basic_algorithm

func getWinner(arr []int, k int) int {
	if k == 1 {
		return max(arr[0], arr[1])
	}
	flag := max(arr[0], arr[1])
	for i, j := 1, 2; i <= k; i++ {
		if flag != max(arr[0], arr[1]) {
			i = 1
			flag = max(arr[0], arr[1])
		}
		arr[0], arr[1], arr[j] = max(arr[0], arr[1]), arr[j], min(arr[0], arr[1])
		if j == len(arr)-1 {
			j = 2
		} else {
			j++
		}
	}
	return arr[0]
}

// 超时
func getWinner1(arr []int, k int) int {
	if len(arr) < 2 {
		return -1
	}
	flag := max(arr[0], arr[1])
	for i := 1; i <= k; i++ {
		if flag != max(arr[0], arr[1]) {
			i = 1
			flag = max(arr[0], arr[1])
		}
		temp := min(arr[0], arr[1])
		arr[0] = max(arr[0], arr[1])
		for j := 1; j < len(arr)-1; j++ {
			arr[j] = arr[j+1]
		}
		arr[len(arr)-1] = temp
	}
	return arr[0]
}
