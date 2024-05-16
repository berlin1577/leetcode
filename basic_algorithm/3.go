package basic_algorithm

// 注意：go 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码不保证正确性，仅供参考。如有疑惑，可以参照我写的 cpp 代码对比查看。

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)

	left, right := 0, 0
	res := 0 // 记录结果
	for right < len(s) {
		c := s[right]
		right++
		// 进行窗口内数据的一系列更新
		window[c]++
		// 判断左侧窗口是否要收缩
		for window[c] > 1 {
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			window[d]--
		}
		// 在这里更新答案
		res = max(res, right-left)
	}
	return res
}

// 思路错误
func lengthOfLongestSubstring1(s string) int {
	var t byte
	left, right, ans, tag := 0, 0, 0, false
	window := map[byte]bool{}
	for right < len(s) {
		if !window[s[right]] {
			ans++
			window[s[right]] = true
			right++
			continue
		}
		tag = false
		t = s[right]
		right++
		for tag {
			window[s[left]] = false
			ans--
			if s[left] == t {
				tag = true
				t = '1'
			}
			left++
		}
	}
	return ans
}
