package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	passports := readPassports()
	validPassports := 0
	for _, passport := range passports {
		if containsRequiredFields(passport) {
			validPassports++
		}
	}
	fmt.Printf("valid passports: %d", validPassports)
}

func containsRequiredFields(passport string) bool {
	for _, field := range requiredFields() {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}

func requiredFields() []string {
	return []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
}

func readPassports() []string {
	file, err := ioutil.ReadFile("04-passport-processing1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n\n")
}
