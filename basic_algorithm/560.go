package basic_algorithm

func subarraySum(nums []int, k int) int {
	res := 0
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	for left := 0; left < len(nums); left++ {
		for right := left; right < len(nums); right++ {
			if preSum[right+1]-preSum[left] == k {
				res++
			}
		}
	}
	return res
}
