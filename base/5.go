package base

// 回文串是快慢指针中比较特殊的一种，使用左右指针可以解决，
//与普通的左右指针不同，求回文串是从中间往外扩散来寻找回文串
//再使用一个for循环，枚举出所有可能是中心节点的位置，需要注意回文串有偶数串和奇数串两种。
func longestPalindrome(s string) string {
	ans := ""
	for t, _ := range s {
		te := FindPalindrome(s, t, t)
		fe := FindPalindrome(s, t, t + 1)
		if len(te) > len(ans) {
			ans = te
		}
		if len(fe) > len(ans) {
			ans = fe
		}
	}
	return ans
}

func FindPalindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left + 1: right]
//	go语言中的切片slice是左闭右开的，最左边的元素能取到，最右边的元素取不到，所以left要加一
}
