package base

// 1.熟悉go语言基本语法
// 2.使用hashmap来空间换时间
func twoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	for i, x := range nums {
		if p, ok := hashmap[target-x]; ok {
			return []int{p, i}
		}
		hashmap[x] = i
	}
	return nil
}
