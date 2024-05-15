package basic_algorithm

//本题是双指针类型的题，双指针类型有快慢指针

//left指针指向已排序序列的末尾，right指针指向未排序序列的开始，right指针一直向前，当有一个不为0的元素将它与left位置交换，然后left向前，直到最后

func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
		}
		right++
	}
}
