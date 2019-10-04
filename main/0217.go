package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	var exist = make(map[int]int32)
	for _, val := range nums {
		var _, ok = exist[val]
		if ok {
			return true
		}
		exist[val] = 1
	}
	return false
}
