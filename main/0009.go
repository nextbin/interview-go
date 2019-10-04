package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(100))
}

func isPalindrome(x int) bool {
	var str = fmt.Sprintf("%d", x)
	var len = len(str)
	for idx, _ := range str {
		if str[idx] != str[len-1-idx] {
			return false
		}
	}
	return true
}
