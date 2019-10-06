package solution

import "fmt"

func Run0263() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(8))
	fmt.Println(isUgly(14))
	fmt.Println(isUgly(1))
	fmt.Println(isUgly(0))
	fmt.Println(isUgly(-1))
}

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for num != 1 {
		if (num & 1) == 0 {
			num /= 2
			continue
		}
		if num%3 == 0 {
			num /= 3
			continue
		}
		if num%5 == 0 {
			num /= 5
			continue
		}
		return false
	}
	return true
}
