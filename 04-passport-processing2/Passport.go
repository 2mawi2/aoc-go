package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         Height
	HairColor      string
	EyeColor       string
	PassportId     string
}

func (p Passport) IsValid() bool {
	if p.BirthYear < 1920 || p.BirthYear > 2002 {
		return false
	}
	if p.IssueYear < 2010 || p.IssueYear > 2020 {
		return false
	}
	if p.ExpirationYear < 2020 || p.ExpirationYear > 2030 {
		return false
	}
	if !p.Height.IsValid() {
		return false
	}
	if !isValidHairColor(p.HairColor) {
		return false
	}
	if !isValidEyeColor(p.EyeColor) {
		return false
	}
	if !isValidEyeColor(p.EyeColor) {
		return false
	}
	if !isValidPassportId(p.PassportId) {
		return false
	}
	return true
}

func isValidPassportId(passportId string) bool {
	valid9DigitNumberRegex := regexp.MustCompile("^[0-9]{9}$")
	return valid9DigitNumberRegex.MatchString(passportId)
}

func isValidEyeColor(color string) bool {
	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, validColor := range validEyeColors {
		if validColor == color {
			return true
		}
	}
	return false
}

func isValidHairColor(hairColor string) bool {
	validHexColorRegex := regexp.MustCompile("^#([0-9a-f]{6})$")
	return validHexColorRegex.MatchString(hairColor)
}

func NewPassportFromInput(input string) Passport {
	fields := parseFieldsMap(input)
	birthdate, _ := strconv.Atoi(fields["byr"])
	issueDate, _ := strconv.Atoi(fields["iyr"])
	expirationYear, _ := strconv.Atoi(fields["eyr"])
	return Passport{
		BirthYear:      birthdate,
		IssueYear:      issueDate,
		ExpirationYear: expirationYear,
		Height:         NewHeightFromInput(fields["hgt"]),
		HairColor:      fields["hcl"],
		EyeColor:       fields["ecl"],
		PassportId:     fields["pid"],
	}
}

func parseFieldsMap(input string) map[string]string {
	fields := strings.Split(strings.ReplaceAll(input, "\n", " "), " ")
	fieldsMap := map[string]string{}
	for _, field := range fields {
		elements := strings.Split(field, ":")
		fieldsMap[elements[0]] = elements[1]
	}
	return fieldsMap
}
