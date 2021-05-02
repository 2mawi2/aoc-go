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
		distinctAnswers := getDistinctAnswers(groupAnswer)
		for answer, _ := range distinctAnswers {
			if allGroupAnswersContainAnswer(groupAnswer, answer) {
				totalSumOfGroupAnswers++
			}
		}
	}
	fmt.Printf("%v", totalSumOfGroupAnswers)
}

func allGroupAnswersContainAnswer(groupAnswers Group, answer string) bool {
	for _, line := range groupAnswers.answer {
		if !strings.Contains(line, answer) {
			return false
		}
	}
	return true
}

func getDistinctAnswers(groupAnswer Group) map[string]bool {
	visitedAnswers := make(map[string]bool)
	for _, answer := range groupAnswer.answer {
		for _, letter := range answer {
			visitedAnswers[string(letter)] = true
		}
	}
	return visitedAnswers
}

type Group struct {
	answer []string
}

func readGroupAnswers() []Group {
	file, err := ioutil.ReadFile("06-custom-customs2/input.txt")
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
