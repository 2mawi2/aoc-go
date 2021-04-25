package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type BoardingPass struct {
	Row    int
	Column int
	SeatId int
}

func main() {
	seatNumbers := readSeatNumbers()
	max := 0
	for _, seatNumber := range seatNumbers {
		boardingPass := parseBoardingPass(seatNumber)
		if boardingPass.SeatId >= max {
			max = boardingPass.SeatId
		}
	}
	fmt.Println(max)
}

func parseBoardingPass(seatNumber string) BoardingPass {
	row := binarySearchRow(seatNumber)
	column := binarySearchColumn(seatNumber)
	return BoardingPass{
		Row:    row,
		Column: column,
		SeatId: row*8 + column,
	}
}

func binarySearchColumn(seatNumber string) int {
	lower := 0
	upper := 7
	for i := 7; i < 10; i++ {
		indicator := seatNumber[i]
		center := (upper-lower)/2 + lower
		if indicator == 'L' {
			upper = center
		} else {
			lower = center + 1
		}
	}
	if upper != lower {
		panic("Column could not be found!")
	}
	return upper
}

func binarySearchRow(seatNumber string) int {
	lower := 0
	upper := 127
	for i := 0; i < 7; i++ {
		indicator := seatNumber[i]
		center := (upper-lower)/2 + lower
		if indicator == 'F' {
			upper = center
		} else {
			lower = center + 1
		}
	}
	if upper != lower {
		panic("Row could not be found!")
	}
	return upper
}

func readSeatNumbers() []string {
	file, err := ioutil.ReadFile("05-binary-boarding1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}
