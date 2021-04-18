package main

import (
	"aoc-go/common"
	"fmt"
)

func main() {
	lines := common.ReadInputLines("03-toboggan-trajectory1/input.txt")

	trees := traverseMap(lines)

	fmt.Printf("Trees found: %d", trees)
}

func traverseMap(lines []string) int {
	trees := 0
	x := 0
	for y := 1; y < len(lines); y++ {
		x += 3
		field := getAtRepeatedPosition(lines[y], x)
		if field == "#" {
			trees++
		}
	}
	return trees
}

func getAtRepeatedPosition(line string, position int) string {
	return string([]rune(line)[position%len(line)])
}
