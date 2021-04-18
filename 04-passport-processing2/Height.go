package main

import (
	"strconv"
	"strings"
)

type Unit int

const (
	centimeters Unit = iota
	inch
)

type Height struct {
	Value int
	Unit  Unit
}

func (h Height) IsValid() bool {
	if h.Unit == centimeters && (h.Value < 150 || h.Value > 193) {
		return false
	}
	if h.Unit == inch && (h.Value < 59 || h.Value > 76) {
		return false
	}
	return true
}

func NewHeightFromInput(input string) Height {
	if strings.Contains(input, "cm") {
		value, _ := strconv.Atoi(strings.ReplaceAll(input, "cm", ""))
		return Height{
			Value: value,
			Unit:  centimeters,
		}
	}
	if strings.Contains(input, "in") {
		value, _ := strconv.Atoi(strings.ReplaceAll(input, "in", ""))
		return Height{
			Value: value,
			Unit:  inch,
		}
	}
	return Height{}
}
