package easy

import "fmt"

func Run0009() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(100))
}

func isPalindrome(x int) bool {
	var str = fmt.Sprintf("%d", x)
	var size = len(str)
	for idx := range str {
		if str[idx] != str[size-1-idx] {
			return false
		}
	}
	return true
}
