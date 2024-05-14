package 代码框架

//二分查找的思路很简单，但是边界问题很细致
//计算时使用left + (right - left) / 2,有效防止了left 和 right太大,直接相加导致溢出的情况
//不管是左闭右开还是左闭右闭,当一次查找之后,nums[mid]不是target值,之后就要根据区间来排除已经判断过不是的元素.

//最简单的二分查找
func binarysearch(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := (low + high)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	return -1
//	return nums[left] == target ? left : -1
}

//接下来是搜索target的左边结和右边界,只要记清搜索区域是左闭右开的就行
//最后要检查一下两种nums数组中没有target的情况,一种是查过len(nums),还有一种情况是nums[left] != target,这两种情况可以return -1


//搜索左边界
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			right = mid
		}else if nums[mid] < target {
			right = mid
		}else if nums[mid] > target {
			left = mid + 1
		}
	}
	if left >= len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	}else {
		return -1
	}
}

//搜索右边界
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			left = mid + 1
		}else if nums[mid] > target {
			right = mid
		}else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left - 1 >= len(nums) {
		return -1
	}
	if nums[left - 1] == target {
		return left - 1
	}else {
		return -1
	}
}
