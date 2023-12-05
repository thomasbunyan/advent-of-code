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

func StrToIntSlice(strList []string) []int {
	out := make([]int, len(strList))
	for index, str := range strList {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		out[index] = i
	}
	return out
}
