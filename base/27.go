package base

func removeElement(nums []int, val int) int {
	i, j, k := 0, 0, 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
			k++
		}
		j++
	}
	return k
}
