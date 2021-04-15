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