package medium

func groupThePeople(groupSizes []int) [][]int {
	var res [][]int
	m := map[int]int{}
	for i, val := range groupSizes {
		idx, ok := m[val]
		if !ok {
			res = append(res, []int{})
			idx = len(res) - 1
			m[val] = idx
		}
		res[idx] = append(res[idx], i)
		if len(res[idx]) == val {
			delete(m, val)
		}
	}
	return res
}
