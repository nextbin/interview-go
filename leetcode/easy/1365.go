package easy

func smallerNumbersThanCurrent(nums []int) []int {
	var res []int
	// 方案一：类似冒泡穷举。时间O(n*n)，空间O(n)
	for idx1, val1 := range nums {
		var cnt = 0
		for idx2, val2 := range nums {
			if idx1 != idx2 && val1 > val2 {
				cnt += 1
			}
		}
		res = append(res, cnt)
	}
	// 方案二：排序。时间(n*logn)，空间O(n)
	return res
}
