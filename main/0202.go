package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(100))
	fmt.Println(isHappy(1))
	fmt.Println(isHappy(0))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	var exist = make(map[int]int)
	for n != 1 {
		var sum = 0
		var tmp = n
		for tmp != 0 {
			var tmp2 = tmp % 10
			sum += tmp2 * tmp2
			tmp /= 10
		}
		var _, ok = exist[sum]
		if ok {
			return false
		}
		exist[sum] = 1
		n = sum
	}
	return true
}
