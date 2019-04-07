package mygo

import "strconv"

func ItoA(num int64) string {
	return strconv.FormatInt(num, 10)
}

func GetHello() string {
	return "Hi"
}
