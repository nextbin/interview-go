package easy

func runningSum(nums []int) []int {
	var res []int
	var last_val = 0
	for _, val := range nums {
		res = append(res, last_val+val)
		last_val += val
	}
	return res
}
