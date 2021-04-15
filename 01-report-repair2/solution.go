package main

import (
	"aoc-go/common"
	"fmt"
	"strconv"
)

func main() {
	reportNumbers := readReportNumbers()
	for _, first := range reportNumbers {
		for _, second := range reportNumbers {
			for _, third := range reportNumbers {
				if first == second || first == third || third == second{
					continue
				} else if first+second+third == 2020 {
					fmt.Printf("Solution is: %v", first*second*third)
					return
				}
			}
		}
	}
}


func readReportNumbers() []int {
	lines := common.ReadInputLines("01-report-repair2/input.txt")
	var reportNumbers []int
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		reportNumbers = append(reportNumbers, number)
	}
	return reportNumbers
}
