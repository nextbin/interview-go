package easy

func numberOfSteps(num int) int {
	var res = 0
	for num != 0 {
		if num&1 == 0 {
			num >>= 1
		} else {
			num -= 1
		}
		res += 1
	}
	return res
}
