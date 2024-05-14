package 代码框架


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
