package basic_algorithm

import "math"

//最小覆盖子串

// 使用valid来记录t中的某个字符是否已经在滑动窗口中全部集齐
func MinWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i, _ := range t {
		need[t[i]]++
	}
	left, right, valid, length, start := 0, 0, 0, math.MaxInt32, 0
	for right < len(s) {
		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				valid++
			}
		}
		right++
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			if _, ok := need[s[left]]; ok {
				if window[s[left]] == need[s[left]] {
					valid--
				}
				window[s[left]]--
			}
			left++
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

// 官解
//func minWindow(s string, t string) string {
//	ori, cnt := map[byte]int{}, map[byte]int{}
//	for i := 0; i < len(t); i++ {
//		ori[t[i]]++
//	}
//
//	sLen := len(s)
//	len := math.MaxInt32
//	ansL, ansR := -1, -1
//
//	check := func() bool {
//		for k, v := range ori {
//			if cnt[k] < v {
//				return false
//			}
//		}
//		return true
//	}
//	for l, r := 0, 0; r < sLen; r++ {
//		if r < sLen && ori[s[r]] > 0 {
//			cnt[s[r]]++
//		}
//		for check() && l <= r {
//			if r-l+1 < len {
//				len = r - l + 1
//				ansL, ansR = l, l+len
//			}
//			if _, ok := ori[s[l]]; ok {
//				cnt[s[l]] -= 1
//			}
//			l++
//		}
//	}
//	if ansL == -1 {
//		return ""
//	}
//	return s[ansL:ansR]
//}

// 超时
//func minWindow1(s string, t string) string {
//	need, window := map[byte]int{}, map[byte]int{}
//	slen := len(s)
//	left, right, length, ansl, ansr := 0, 0, math.MaxInt32, -1, -1
//	for i, _ := range t {
//		need[t[i]]++
//	}
//	check := func() bool {
//		for k, v := range need {
//			if window[k] < v {
//				return false
//			}
//		}
//		return true
//	}
//	for ; right < slen; right++ {
//		k := s[right]
//		if right < slen && need[k] > 0 {
//			window[k]++
//		}
//		for check() && left <= right {
//			l := s[left]
//			if right-left+1 < length {
//				length = right - left + 1
//				ansl, ansr = left, left+length
//			}
//			if _, ok := need[l]; ok {
//				window[l] -= 1
//			}
//			left++
//		}
//	}
//	if ansl == -1 {
//		return ""
//	}
//	return s[ansl:ansr]
//}

//func MinWindow(s string, t string) string {
//	left, right := 0, 0
//	need, windows := map[byte]int{}, map[byte]int{}
//	ans := ""
//	for _, i := range t {
//		need[byte(i)]++
//	}
//	check := func() bool {
//		for i, _ := range need {
//			if windows[i] < need[i] {
//				return false
//			}
//		}
//		return true
//	}
//	for ; right < len(s); right++ {
//		for ; check(); left++ {
//			if ans == "" {
//				ans = s[left:right]
//			} else if ans != "" && len(ans) < len(s[left:right]) {
//				ans = s[left:right]
//			}
//			windows[s[left]]--
//		}
//		windows[s[right]]++
//	}
//	return ans
//}

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
