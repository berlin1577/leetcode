package basic_algorithm

//128.最长连续序列
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

//第一个元素要开始进行第二层循环，一定没有i-1
//使用 for range循环结构时，如果不需要建或者值，必须要用_来忽略输出的结果

func longestConsecutive(nums []int) int {
	hash := map[int]bool{}
	for _, i := range nums {
		hash[i] = true
	}
	ans := 1
	for _, i := range nums {
		if hash[i-1] != true {
			t := i
			a := 1
			for hash[t+1] == true {
				a++
				t++
			}
			if a > ans {
				ans = a
			}
		}
	}
	return ans
}
