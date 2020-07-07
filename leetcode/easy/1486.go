package easy

func xorOperation(n int, start int) int {
	res := 0
	val := start
	for i := 0; i < n; i++ {
		res = res ^ val
		val += 2
	}
	return res
}
