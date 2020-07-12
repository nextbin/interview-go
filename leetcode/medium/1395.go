package medium

func numTeams(rating []int) int {
	gtSlice := []int{}
	ltSlice := []int{}
	for i, v1 := range rating {
		gt := 0
		lt := 0
		for j, v2 := range rating {
			if i < j {
				if v1 > v2 {
					gt += 1
				} else if v1 < v2 {
					lt += 1
				}
			}
		}
		gtSlice = append(gtSlice, gt)
		ltSlice = append(ltSlice, lt)
	}
	res := 0
	for i, v1 := range rating {
		for j, v2 := range rating {
			if j <= i {
				continue
			}
			if v1 < v2 {
				res += ltSlice[j]
			} else if v1 > v2 {
				res += gtSlice[j]
			}
		}
	}
	return res
}
