package common

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadInputLines(location string) []string {
	file, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	return lines
}

func Contains(slice []int, element int) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}

func ContainsAll(slice []int, elements []int) bool {
	for _, e := range elements {
		if !Contains(slice, e) {
			return false
		}
	}
	return true
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
