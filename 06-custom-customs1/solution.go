package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	groupAnswers := readGroupAnswers()
	totalSumOfGroupAnswers := 0
	for _, groupAnswer := range groupAnswers {
		visitedAnswers := make(map[string]bool)
		for _, answer := range groupAnswer.answer {
			for _, letter := range answer {
				visitedAnswers[string(letter)] = true
			}
		}
		totalSumOfGroupAnswers += len(visitedAnswers)
	}
	fmt.Printf("%v", totalSumOfGroupAnswers)
}

type Group struct {
	answer []string
}

func readGroupAnswers() []Group {
	file, err := ioutil.ReadFile("06-custom-customs1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var groups []Group
	groupsStrings := strings.Split(string(file), "\n\n")
	for _, groupsString := range groupsStrings {
		answers := strings.Split(groupsString, "\n")
		groups = append(groups, Group{answer: answers})
	}
	return groups
}
