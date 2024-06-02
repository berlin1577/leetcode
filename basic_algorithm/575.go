package basic_algorithm

func distributeCandies(candyType []int) int {
	ans := 0
	hp := map[int]bool{}
	for i := 0; i < len(candyType) && ans < len(candyType)/2; i++ {
		if !hp[candyType[i]] {
			ans++
			hp[candyType[i]] = true
		}
	}
	return ans
}
