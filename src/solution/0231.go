package solution

import "fmt"

func Run0231() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(218))
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(5))
	fmt.Println(isPowerOfTwo(8))
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (-n)) == n
}
