package basic_algorithm

func findAnagrams(s string, p string) []int {
	ans := []int{}
	var need, window [26]int
	if len(s) < len(p) {
		return ans
	}
	for i, k := range p {
		need[k-'a']++
		window[s[i]-'a']++
	}
	if need == window {
		ans = append(ans, 0)
	}
	for i, k := range s[:len(s)-len(p)] {
		window[k-'a']--
		window[s[i+len(p)]-'a']++
		if need == window {
			ans = append(ans, i+1)
		}
	}
	return ans
}

// 滑动窗口解法，固定窗口大小的滑动窗口，
func findAnagrams1(s string, p string) []int {
	var ans []int
	need, window := map[byte]int{}, map[byte]int{}
	left, right, length := 0, 0, len(p)
	for i, _ := range p {
		need[p[i]]++
	}
	check := func() bool {
		for v, k := range need {
			if window[v] < k {
				return false
			}
		}
		return true
	}
	for right < len(s) {
		if right-left < length {
			window[s[right]]++
		}
		right++
		for right-left == length {
			if check() {
				ans = append(ans, left)
			}
			window[s[left]]--
			left++
		}
	}
	return ans
}
