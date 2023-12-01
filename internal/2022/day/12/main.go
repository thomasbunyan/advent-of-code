package main

import (
	"fmt"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

const (
	CURRENT = "S"
	DEST    = "E"
)

func main() {
	input := common.ReadFast("./internal/2022/day/12/input.txt")

	heightMap := [][]string{}
	x, y := -1, -1

	for j, line := range input {
		lineSplit := (strings.Split(line, ""))

		if x == -1 && y == -1 {
			for i := range lineSplit {
				if lineSplit[i] == CURRENT {
					x, y = i, j
					lineSplit[i] = "a"
				}
			}
		}

		heightMap = append(heightMap, lineSplit)
	}

	steps := traverse(heightMap, x, y, make(map[string]bool))

	fmt.Println(steps)
}

// stop a path when too long

func traverse(heightMap [][]string, x, y int, path map[string]bool) int {
	nonMisses := -1
	max := heightMap[y][x][0] + 1
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			path[fmt.Sprint(y, x)] = true
			if (i == 0 || j == 0) && (i != 0 || j != 0) {
				newY, newX := y+i, x+j
				if newY >= 0 && newX >= 0 && newY < len(heightMap) && newX < len(heightMap[0]) {
					if heightMap[y][x] == "z" && heightMap[newY][newX] == DEST {
						return 1
					} else if !path[fmt.Sprint(newY, newX)] {
						if heightMap[newY][newX][0] <= max && heightMap[newY][newX] != DEST {
							res := traverse(heightMap, newX, newY, copyMap(path))
							if res == -1 {
							} else {
								if nonMisses == -1 || nonMisses >= res+1 {
									nonMisses = res + 1
								}
							}
						}
					}
				}
			}
		}
	}
	if nonMisses != -1 {
		return nonMisses
	}
	return -1
}

func copyMap(m map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}
