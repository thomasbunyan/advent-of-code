package main

import (
	"fmt"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	lines := common.ReadFast("./internal/2023/day/8/input.txt")

	directions := lines[0]

	network := make(map[string][]string)

	for _, line := range lines[2:] {
		fields := strings.Fields(line)
		network[fields[0]] = []string{string(fields[2][1:4]), string(fields[3][:3])}
	}

	count := 0

	for position := "AAA"; position != "ZZZ"; count++ {
		dir := getDir(directions[count%len(directions)])
		position = network[position][dir]
	}

	fmt.Printf("Ans 1: %d\n", count)

	positions := []string{}
	for node := range network {
		if node[2] == 'A' {
			positions = append(positions, node)
		}
	}

	counts := make([]int, len(positions))

	for index, position := range positions {
		count = 0
		for ; position[2] != 'Z'; count++ {
			dir := getDir(directions[count%len(directions)])
			position = network[position][dir]
		}
		counts[index] = count
	}

	fmt.Printf("Ans 2: %d\n", LCM(counts[0], counts[1], counts[2:]...))
}

func getDir(dir byte) int {
	if dir == 'L' {
		return 0
	} else if dir == 'R' {
		return 1
	} else {
		panic("bad dir")
	}
}

// https://go.dev/play/p/SmzvkDjYlb
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}
