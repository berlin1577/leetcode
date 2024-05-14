package 代码框架

//二分查找的思路很简单，但是边界问题很细致
//计算时使用left + (right - left) / 2,有效防止了left 和 right太大,直接相加导致溢出的情况

func binarysearch(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := (low + high)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			high--
		}else {
			low++
		}
	}
	return -1
}
