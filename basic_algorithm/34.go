package basic_algorithm

func searchRange(nums []int, target int) []int {
	left1, left2, right1, right2 := 0, 0, len(nums), len(nums)
	for left1 < right1 || left2 < right2 {
		if left1 < right1 {
			mid1 := left1 + (right1-left1)/2
			if nums[mid1] == target {
				right1 = mid1
			} else if nums[mid1] < target {
				left1 = mid1 + 1
			} else if nums[mid1] > target {
				right1 = mid1
			}
		}
		if left2 < right2 {
			mid2 := left2 + (right2-left2)/2
			if nums[mid2] == target {
				left2 = mid2 + 1
			} else if nums[mid2] < target {
				left2 = mid2 + 1
			} else if nums[mid2] > target {
				right2 = mid2
			}
		}
	}
	if left1 >= len(nums) || left2-1 >= len(nums) {
		return []int{-1, -1}
	}
	if nums[left1] == target && nums[left2-1] == target {
		return []int{left1, left2 - 1}
	} else {
		return []int{-1, -1}
	}
}
