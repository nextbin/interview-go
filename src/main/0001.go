package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 1}, 4))
}

func twoSum(nums []int, target int) []int {
	var exist = make(map[int]int)
	for idx, val := range nums {
		var idx2, ok = exist[target-val]
		if ok {
			return []int{idx2, idx}
		}
		exist[val] = idx
	}
	return []int{}
}
