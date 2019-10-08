package easy

import (
	"fmt"
)

func Run0007() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(100000012345))
	//fmt.Println(reverse(10000001234111111115))
}

func reverse(x int) int {
	var sign = 1
	if x < 0 {
		sign = -1
		x = -x
	}
	var ret = 0
	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	if sign < 0 {
		ret = -ret
	}
	if ret > (1<<31-1) || ret < -(1<<31) {
		return 0
	}
	return ret
}
