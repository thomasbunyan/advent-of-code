package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	lines := common.ReadFast("./internal/2023/day/1/input.txt")

	numberMap := getNumberMap()

	sum := 0

	for _, line := range lines {

		var first, last string

		firstIndex, lastIndex := len(line), -1
		for numberKey, numberVal := range numberMap {
			iF := strings.Index(line, numberKey)
			iL := strings.LastIndex(line, numberKey)
			if iF >= 0 && iF < firstIndex {
				firstIndex = iF
				first = numberVal
			}
			if iL >= 0 && iL > lastIndex {
				lastIndex = iL
				last = numberVal
			}
		}

		for i := 0; i < len(line); i++ {
			if isInt(string(line[i])) {
				if i < firstIndex {
					first = string(line[i])
				}
				break
			}
		}

		for j := len(line) - 1; j >= 0; j-- {
			if isInt(string(line[j])) {
				if j > lastIndex {
					last = string(line[j])
				}
				break
			}
		}

		if first == "" || last == "" {
			panic("missing line value")
		}

		sum += common.StrToInt(fmt.Sprintf("%s%s", first, last))
	}

	fmt.Printf("Ans: %d\n", sum)
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	} else {
		return true
	}
}

func getNumberMap() map[string]string {
	return map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
}
