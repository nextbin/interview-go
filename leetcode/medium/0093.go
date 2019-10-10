package medium

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}
	var ret = []string{string(s[0])}
	for _, char := range s[1:] {
		var val = char - '0'
		var toAppend = []string{}
		for _, str := range ret {
			if strings.HasSuffix(str, ".") {
				toAppend = append(toAppend, str+string(char))
			} else {
				var i, _ = strconv.Atoi(str[strings.LastIndex(str, ".")+1:])
				if int32(i)*10+val <= 255 && i != 0 {
					toAppend = append(toAppend, str+string(char))
				}
				if len(strings.Split(str, ".")) < 4 {
					toAppend = append(toAppend, str+"."+string(char))
				}
			}
		}
		ret = toAppend
	}
	var valid = []string{}
	for _, str := range ret {
		if len(strings.Split(str, ".")) == 4 {
			valid = append(valid, str)
		}
	}
	return valid
}
