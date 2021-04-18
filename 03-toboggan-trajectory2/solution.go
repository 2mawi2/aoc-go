package main

import (
	"aoc-go/common"
	"fmt"
)

type slope struct {
	deltaX int
	deltaY int
}

func main() {
	lines := common.ReadInputLines("03-toboggan-trajectory2/input.txt")

	treeProduct := 1
	for _, slope := range []slope{
		{deltaX: 1, deltaY: 1},
		{deltaX: 3, deltaY: 1},
		{deltaX: 5, deltaY: 1},
		{deltaX: 7, deltaY: 1},
		{deltaX: 1, deltaY: 2},
	} {
		treeProduct *= traverseMap(lines, slope.deltaX, slope.deltaY)
	}

	fmt.Printf("Found tree product of all slopes: %d", treeProduct)
}

func traverseMap(lines []string, deltaX, deltaY int) int {
	trees := 0
	x := 0
	for y := deltaY; y < len(lines); y += deltaY {
		x += deltaX
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
