package 代码框架

//滑动窗口算法框架,时间复杂度为O(n)
//滑动窗口主要用来解决子数组问题，比如寻找某个符合条件的最长/最短子数组

func slidingWindow1(s string, t string) string {
	
}






















func slidingWindow(s string, t string) string {
	//用来记录窗口中的字符出现的次数
	need, window := make(map[byte]int), make(map[byte]int)
	//初始化need
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	//初始化窗口的左右边界
	left, right := 0, 0
	//用来记录符合条件的子串的起始位置和长度
	start, length := 0, 0
	//用来记录窗口中满足条件的字符个数
	valid := 0
	//开始滑动窗口
	for right < len(s) {
		//c是将移入窗口的字符
		c := s[right]
		//右移窗口
		right++
		//进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		//判断左侧窗口是否要收缩
		for right - left >= len(t) {
			//更新最小覆盖子串
			if valid == len(need) {
				return s[left:right]
			}
			//d是将移出窗口的字符
			d := s[left]
			//左移窗口
			left++
			//进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	//返回最小覆盖子串
	return ""
}
