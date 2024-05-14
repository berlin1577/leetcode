package 代码框架

//二分查找的思路很简单，但是边界问题很细致
//计算时使用left + (right - left) / 2,有效防止了left 和 right太大,直接相加导致溢出的情况
//不管是左闭右开还是左闭右闭,当一次查找之后,nums[mid]不是target值,之后就要根据区间来排除已经判断过不是的元素.

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
