package solution

import "fmt"

func Run0242() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("abc", "acb"))
}

func isAnagram(s string, t string) bool {
	var cntMap = make(map[int32]int)
	for _, char := range s {
		var _, ok = cntMap[char]
		if !ok {
			cntMap[char] = 1
		} else {
			cntMap[char]++
		}
	}
	for _, char := range t {
		var _, ok = cntMap[char]
		if !ok {
			return false
		} else {
			cntMap[char]--
		}
	}
	for _, v := range cntMap {
		if v != 0 {
			return false
		}
	}
	return true
}
