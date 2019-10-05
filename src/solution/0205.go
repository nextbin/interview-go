package solution

import "fmt"

func Run0205() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}

func isIsomorphic(s string, t string) bool {
	var firstExistS = make(map[int32]int)
	var mapS = make(map[int]int)
	for idx, val := range s {
		var firstIdx, ok = firstExistS[val]
		if ok {
			mapS[idx] = firstIdx
		} else {
			mapS[idx] = idx
			firstExistS[val] = idx
		}
	}
	var firstExistT = make(map[int32]int)
	for idx, val := range t {
		var firstIdx, ok = firstExistT[val]
		if ok {
			if firstIdx != mapS[idx] {
				return false
			}
		} else {
			if idx != mapS[idx] {
				return false
			}
			firstExistT[val] = idx
		}
	}
	return true
}
