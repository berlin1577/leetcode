package basic_algorithm

func maxDivScore(nums []int, divisors []int) int {
	score, ans := -1, 0
	for i := 0; i < len(divisors); i++ {
		t := 0
		for j := 0; j < len(nums); j++ {
			if nums[j]%divisors[i] == 0 {
				t++
			}
		}
		if score < t {
			score = t
			ans = divisors[i]
		} else if score == t {
			ans = min(ans, divisors[i])
		}
	}
	return ans
}
