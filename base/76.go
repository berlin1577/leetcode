package base

//最小覆盖子串

func minWindow(s string, t string) string {
	left, right := 0, 0
	need, windows := map[byte]int{}, map[byte]int{}
	ans := ""
	for _, i := range t {
		need[byte(i)]++
	}
	check := func() bool {
		for i, _ := range need {
			if windows[i] < need[i] {
				return false
			}
		}
		return true
	}
	for ; right < len(s) ; right++{
		for ; check(); left++ {
			if ans == "" {
				ans = s[left : right]
			}else if ans != "" && len(ans) < len(s[left : right]) {
				ans = s[left : right]
			}
			windows[s[left]]--
		}
		windows[s[right]]++
	}
	return ans
}




























//func minWindow(s string, t string) string {
//	left, right := 0, 0
//	ans := ""
//	hp := map[string][26]int{}
//	for ; right < len(s); right++ {
//		if SubString(s[left : right], t, hp){
//			for ; SubString(s[left : right], t, hp); left++ {
//			}
//			ans = s[left - 1 : right]
//		}
//	}
//	return ans
//}
//func SubString(s, t string, hp map[string][26]int) bool {
//	hpcp := hp[s]
//	for _, i := range t {
//		j := i - 'a'
//		hpcp[j]--
//		if hpcp[j] <= 0 {
//			return false
//		}
//	}
//	return true
//}
