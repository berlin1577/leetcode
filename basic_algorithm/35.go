package basic_algorithm

// 二分查找
// 本题在于查找目标值的插入位置条件是找到第一个第一个大于等于目标值的元素,该元素的位置就是目标值的插入位置
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			ans = mid
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}

	return ans
}
