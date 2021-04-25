package main

import (
	common "aoc-go/common"
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
	existingSeats := getExistingSeats(seatNumbers)
	allSeats := common.MakeRange(0, 127*8*7)
	findMySeat(allSeats, existingSeats)
}

func findMySeat(allSeats []int, existingSeats []int) {
	for _, seat := range allSeats {
		row := seat / 8
		if row == 0 || row == 127 {
			continue
		}
		previousRowBooked := common.ContainsAll(existingSeats, getRow(row-1))
		nextRowBooked := common.ContainsAll(existingSeats, getRow(row+1))
		isSeatFree := !common.Contains(existingSeats, seat)
		if previousRowBooked && nextRowBooked && isSeatFree {
			fmt.Printf("Found my seat: %d", seat)
			return
		}
	}
}

func getRow(row int) (seatIds []int) {
	for _, column := range common.MakeRange(0, 7) {
		seatIds = append(seatIds, calcSeatId(row, column))
	}
	return
}

func getExistingSeats(seatNumbers []string) (seats []int) {
	for _, seatNumber := range seatNumbers {
		boardingPass := parseBoardingPass(seatNumber)
		seats = append(seats, boardingPass.SeatId)
	}
	return
}

func parseBoardingPass(seatNumber string) BoardingPass {
	row := binarySearchRow(seatNumber)
	column := binarySearchColumn(seatNumber)
	return BoardingPass{
		Row:    row,
		Column: column,
		SeatId: calcSeatId(row, column),
	}
}

func calcSeatId(row int, column int) int {
	return row*8 + column
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
