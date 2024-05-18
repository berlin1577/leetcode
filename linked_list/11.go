package linked_list

func maxArea(height []int) int {
	i, j, ans := 0, len(height)-1, 0
	for i < j {
		if ans < (j-i)*min(height[i], height[j]) {
			ans = (j - i) * min(height[i], height[j])
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}

// 暴力超时，使用左右指针
func maxArea1(height []int) int {
	ans := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			if (j-i)*min(height[i], height[j]) > ans {
				ans = (j - i) * min(height[i], height[j])
			}
		}
	}
	return ans
}
