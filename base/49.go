package base

import "sort"

//import fmt

// 进一步了解hashmap，可以new一个哈希数组，一个值与一个string数组对应,键也可以是int数组。
//哈希表的题中带你就是找到唯一的特点，用空间换时间。一般时间复杂度是O(n)

func groupAnagrams(strs []string) [][]string {
	hashmap := map[string][]string{}
	for _, t := range strs {
		str := []rune(t)
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		hashmap[string(str)] = append(hashmap[string(str)], t)
	}
	ans := [][]string{}
	for _, v := range hashmap {
		ans = append(ans, v)
	}
	return ans
}
