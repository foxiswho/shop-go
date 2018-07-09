package conv

import "strconv"

func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
