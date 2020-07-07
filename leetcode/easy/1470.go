package easy

func shuffle(nums []int, n int) []int {
	var res []int
	var i int
	for i = 0; i < n; i++ {
		res = append(res, nums[i])
		res = append(res, nums[i+n])
	}
	return res
}
