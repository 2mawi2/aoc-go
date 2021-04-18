package main

import (
	"aoc-go/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	passwordPolicies := readInput()
	var validPasswords = 0
	for _, policy := range *passwordPolicies {
		if policy.isValid() {
			validPasswords++
		}
	}
	fmt.Printf("Valid passwords: %v", validPasswords)
}

type PasswordPolicy struct {
	upper    int
	lower    int
	letter   string
	password string
}

func (p *PasswordPolicy) isValid() bool {
	lowerContained := letterAtPosition(p.password, p.lower) == p.letter
	upperContained := letterAtPosition(p.password, p.upper) == p.letter
	return (lowerContained || upperContained) && !(lowerContained && upperContained)
}

func letterAtPosition(str string, position int) string {
	return string([]rune(str)[position-1])
}

func fromLine(line string) PasswordPolicy {
	policyElements := strings.Split(line, " ")
	boundaries := strings.Split(policyElements[0], "-")
	lower, _ := strconv.Atoi(boundaries[0])
	upper, _ := strconv.Atoi(boundaries[1])
	return PasswordPolicy{
		upper:    upper,
		lower:    lower,
		letter:   strings.Replace(policyElements[1], ":", "", -1),
		password: policyElements[2],
	}
}

func readInput() *[]PasswordPolicy {
	lines := common.ReadInputLines("02-password-philosophy2/input.txt")
	var passwordPolicies []PasswordPolicy
	for _, line := range lines {
		passwordPolicies = append(passwordPolicies, fromLine(line))
	}
	return &passwordPolicies
}
