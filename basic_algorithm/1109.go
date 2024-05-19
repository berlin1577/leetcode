package basic_algorithm

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	diff := make([]int, n)
	for _, record := range bookings {
		diff[record[0]-1] += record[2]
		if record[1] < n {
			diff[record[1]] -= record[2]
		}
	}
	res[0] = diff[0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res
}
