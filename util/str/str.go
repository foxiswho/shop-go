package str

import "strings"

//字符串比较 是否相等  ，这里加入，防止以后忘记
func EqualFold(str1, str2 string) bool {
	return strings.EqualFold(str1, str2)
}
