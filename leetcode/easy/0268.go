package easy

import "fmt"

func Run0268() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{}))
}

func missingNumber(nums []int) int {
	var size = len(nums)
	var res = 0
	for i := 0; i < size; i++ {
		res += i - nums[i]
	}
	res += size
	return res
}
