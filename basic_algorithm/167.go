package basic_algorithm

//本题主要是双指针中的左右指针，用来求子数组问题。
//二分查找

func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left, right := i+1, len(numbers)-1
		for left <= right {
			mid := (left + right) / 2
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				left = mid - 1
			} else {
				right = mid + 1
			}
		}
	}
	return []int{-1, -1}
}

//双指针算法，左右指针相向而行，同时逼近。

func twoSum3(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}
