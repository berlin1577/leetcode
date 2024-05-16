package basic_algorithm

func checkInclusion(s1 string, s2 string) bool {
	need, window := map[byte]int{}, map[byte]int{}
	length, left, right := len(s1), 0, 0
	for h, _ := range s1 {
		need[s1[h]]++
	}
	check := func() bool {
		for i, h := range need {
			if window[i] < h {
				return false
			}
		}
		return true
	}
	for right < len(s2) {
		if right-left < length {
			window[s2[right]]++
			right++
		}
		for ; right-left == length; left++ {
			if check() {
				return true
			} else {
				window[s2[left]]--
			}
		}
	}
	return false
}
