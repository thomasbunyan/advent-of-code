package common

import "strconv"

func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func StrToUInt64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
