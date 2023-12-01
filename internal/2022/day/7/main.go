package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

type File struct {
	dir  bool
	name string
	size int
}

type Dir struct {
	size  int
	files []File
}

func main() {
	lines := common.ReadFast("./internal/2022/day/7/input.txt")

	tree := make(map[string]Dir)
	dirPath := []string{}

	// Build tree
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == "$" {
			switch lineSplit[1] {
			case "cd":
				location := lineSplit[2]
				if location == ".." {
					dirPath = dirPath[:len(dirPath)-1]
				} else {
					dirPath = append(dirPath, lineSplit[2])
				}
			case "ls":
				prefix := strings.Join(dirPath, "/")
				advance, dir := ls(prefix, lines, i)
				tree[prefix] = dir
				i += advance
			default:
				panic("Command not supported")
			}
		}
	}

	// Calculate result
	sizeTree, _ := calculate(tree, "/")

	resultA := 0
	for _, dirSize := range sizeTree {
		if dirSize < 100000 {
			resultA += dirSize
		}
	}

	fmt.Printf("Part 1: %d\n", resultA)

	required := 30000000 - (70000000 - sizeTree["/"])

	min := 70000000
	for _, dirSize := range sizeTree {
		if dirSize >= required && dirSize < min {
			min = dirSize
		}
	}

	fmt.Printf("Part 2: %d\n", min)
}

func calculate(tree map[string]Dir, current string) (map[string]int, int) {
	sizeTree := make(map[string]int)
	contentsSize := tree[current].size
	for _, file := range tree[current].files {
		if file.dir {
			dirSizeTree, size := calculate(tree, file.name)
			for k, v := range dirSizeTree {
				sizeTree[k] = v
			}
			contentsSize += size
		}
	}
	sizeTree[current] = contentsSize
	return sizeTree, contentsSize
}

func ls(prefix string, lines []string, i int) (int, Dir) {
	files := []File{}
	size := 0
	for count, line := range lines[i+1:] {
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == "$" {
			return count, Dir{size: size, files: files}
		}
		if lineSplit[0] == "dir" {
			files = append(files, File{dir: true, name: prefix + "/" + lineSplit[1]})
		} else {
			files = append(files, File{dir: false, name: prefix + "/" + lineSplit[1], size: toInt(lineSplit[0])})
			size += toInt(lineSplit[0])
		}
	}
	return 0, Dir{size: size, files: files}
}

func toInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
