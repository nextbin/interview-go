package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(3))
	fmt.Println(countPrimes(10000))
}

func countPrimes(n int) int {
	var flag []int
	var i = 0
	for i = 0; i < n; i++ {
		flag = append(flag, 0)
	}
	for idx, val := range flag {
		if val == 1 || idx < 2 {
			continue
		}
		var mult = 2
		for mult*idx < n {
			flag[mult*idx] = 1
			mult++
		}
	}
	var ret = 0
	for idx, val := range flag {
		if idx >= 2 && val == 0 {
			ret++
		}
	}
	return ret
}
