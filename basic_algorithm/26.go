package basic_algorithm

//典型的快慢指针和哈希表结合的题目，用一个快指针和一个慢指针维护两个序列，慢指针维护已排序的序列，快指针维护未探寻的序列，可以简化思路
//哈希表用来判断check一个元素的唯一特征，在此题中元素的唯一特征是只出现一次，也就是用值为bool类型，判断是否出现过。

func removeDuplicates(nums []int) int {
	i, j, k := 0, 0, 0
	hp := map[int]bool{}
	for j < len(nums) {
		if hp[nums[j]] != true {
			nums[i] = nums[j]
			k++
			i++
		}
		hp[nums[j]] = true
		j++
	}
	return k
}
func test() {

}
