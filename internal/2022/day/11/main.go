package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

type MonkeyOperation func(int) int

type MonkeyTest func(int) string

type Monkey struct {
	Items     []string        `json:"items"`
	Operation MonkeyOperation `json:"-"`
	Test      MonkeyTest      `json:"-"`
	DivNum    int             `json:"-"`
}

func main() {
	input := common.ReadFast("./internal/2022/day/11/input.txt")

	monkeys := make(map[string]Monkey)

	for i := 0; i < len(input); i++ {
		command := strings.Split(input[i], " ")
		if command[0] == "Monkey" {
			id := strings.Split(command[1], ":")[0]
			i++

			startingItems := strings.Split(input[i], ":")
			i++

			operation := strings.Split(input[i], ":")
			i++

			test := strings.Split(input[i], " ")
			divNum := strToInt(test[len(test)-1])
			i++

			trueCondList := strings.Split(input[i], " ")
			trueCond := trueCondList[len(trueCondList)-1]
			i++

			falseCondList := strings.Split(input[i], " ")
			falseCond := falseCondList[len(falseCondList)-1]

			monkeys[id] = Monkey{
				Items: strings.Split(strings.ReplaceAll(startingItems[1], " ", ""), ","),
				Operation: func(i int) int {
					res := int(0)
					strOperation := strings.Split(strings.TrimSpace(strings.Split(operation[1], "=")[1]), " ")
					for index, op := range strOperation {
						if op == "old" {
							res = i
						} else if op == "+" {
							val := strOperation[index+1]
							if val == "old" {
								return res + i
							} else {
								return res + strToInt(val)
							}
						} else if op == "*" {
							val := strOperation[index+1]
							if val == "old" {
								return res * i
							} else {
								return res * strToInt(val)
							}
						}
					}
					return res
				},
				Test: func(i int) string {
					if i%divNum == 0 {
						return trueCond
					} else {
						return falseCond
					}
				},
				DivNum: divNum,
			}
		}
	}

	lcm := 1
	for _, monkey := range monkeys {
		lcm *= monkey.DivNum
	}

	inspections := make(map[int]int)

	const ROUNDS = 10000
	for round := 0; round < ROUNDS; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[fmt.Sprint(i)]
			for _, item := range monkey.Items {
				inspections[i]++

				itemVal := strToInt(item)
				// perform operation
				itemVal = monkey.Operation(itemVal)
				// divide by 3
				// itemVal /= 3
				itemVal %= lcm
				// run test
				testResult := monkey.Test(itemVal)
				// give item to monkey
				monkey.Items = monkey.Items[1:]
				monkeys[fmt.Sprint(i)] = monkey

				destMonkey := monkeys[testResult]
				destMonkey.Items = append(destMonkey.Items, fmt.Sprint(itemVal))
				monkeys[testResult] = destMonkey
			}
		}
	}

	printJSON(monkeys)
	fmt.Println(inspections)

	monkeyBusiness := []int{}
	for _, count := range inspections {
		monkeyBusiness = append(monkeyBusiness, count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(monkeyBusiness)))

	fmt.Println(monkeyBusiness)
	fmt.Printf("Ans: %d\n", monkeyBusiness[0]*monkeyBusiness[1])
}

func printJSON(input map[string]Monkey) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
