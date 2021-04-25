package main

import (
	"reflect"
	"testing"
)

func Test_parseBoardingPass(t *testing.T) {
	type args struct {
		seatNumber string
	}
	tests := []struct {
		name string
		args args
		want BoardingPass
	}{
		{
			"parses sample 1",
			args{seatNumber: "FBFBBFFRLR"},
			BoardingPass{Row: 44, Column: 5, SeatId: 357},
		},
		{
			"parses sample 2",
			args{seatNumber: "BFFFBBFRRR"},
			BoardingPass{Row: 70, Column: 7, SeatId: 567},
		},
		{
			"parses sample 3",
			args{seatNumber: "FFFBBBFRRR"},
			BoardingPass{Row: 14, Column: 7, SeatId: 119},
		},
		{
			"parses sample 4",
			args{seatNumber: "BBFFBBFRLL"},
			BoardingPass{Row: 102, Column: 4, SeatId: 820},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBoardingPass(tt.args.seatNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBoardingPass() = %v, want %v", got, tt.want)
			}
		})
	}
}
