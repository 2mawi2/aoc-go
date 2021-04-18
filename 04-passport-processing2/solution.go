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
		if passport.IsValid() {
			validPassports++
		}
	}

	fmt.Printf("valid passports: %d", validPassports)
}

func readPassports() []Passport {
	file, err := ioutil.ReadFile("04-passport-processing2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	passportInputs := strings.Split(string(file), "\n\n")
	return mapPassportsFromInput(passportInputs)
}

func mapPassportsFromInput(passportInputs []string) []Passport {
	var passports []Passport
	for _, input := range passportInputs {
		passport := NewPassportFromInput(input)
		passports = append(passports, passport)
	}
	return passports
}
