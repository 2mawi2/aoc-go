package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewPassportFromInput(t *testing.T) {
	// arrange
	input := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm"
	// act
	passport := NewPassportFromInput(input)
	// assert
	assert.Equal(t, 1937, passport.BirthYear)
	assert.Equal(t, 2017, passport.IssueYear)
	assert.Equal(t, 2020, passport.ExpirationYear)
	assert.Equal(t, centimeters, passport.Height.Unit)
	assert.Equal(t, 183, passport.Height.Value)
	assert.Equal(t, "#fffffd", passport.HairColor)
	assert.Equal(t, "gry", passport.EyeColor)
	assert.Equal(t, "860033327", passport.PassportId)
}

func TestPassport_IsValid(t *testing.T) {
	type fields struct {
		BirthYear      int
		IssueYear      int
		ExpirationYear int
		Height         Height
		HairColor      string
		EyeColor       string
		PassportId     string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"true for valid passport",
			fields{
				BirthYear:      2000,
				IssueYear:      2011,
				ExpirationYear: 2021,
				Height: Height{
					Value: 155,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "060033327",
			},
			true,
		},
		{
			"false with birth year to low",
			fields{
				BirthYear:      1919,
				IssueYear:      2011,
				ExpirationYear: 2021,
				Height: Height{
					Value: 155,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "860033327",
			},
			false,
		},
		{
			"false with birth year to high",
			fields{
				BirthYear:      2003,
				IssueYear:      2011,
				ExpirationYear: 2021,
				Height: Height{
					Value: 155,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "860033327",
			},
			false,
		},
		{
			"false with IssueYear to low",
			fields{
				BirthYear:      2000,
				IssueYear:      2009,
				ExpirationYear: 2021,
				Height: Height{
					Value: 155,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "860033327",
			},
			false,
		},
		{
			"false with IssueYear to high",
			fields{
				BirthYear:      2000,
				IssueYear:      2021,
				ExpirationYear: 2021,
				Height: Height{
					Value: 155,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "860033327",
			},
			false,
		},
		{
			"false for expirationYear to low",
			fields{
				BirthYear:      2000,
				IssueYear:      2011,
				ExpirationYear: 2019,
				Height: Height{
					Value: 155,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "860033327",
			},
			false,
		},
		{
			"false for expirationYear to low",
			fields{
				BirthYear:      2000,
				IssueYear:      2011,
				ExpirationYear: 2031,
				Height: Height{
					Value: 155,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "860033327",
			},
			false,
		},
		{
			"false if invalid height",
			fields{
				BirthYear:      2000,
				IssueYear:      2011,
				ExpirationYear: 2021,
				Height: Height{
					Value: 0,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "860033327",
			},
			false,
		},
		{
			"false if invalid color",
			fields{
				BirthYear:      2000,
				IssueYear:      2011,
				ExpirationYear: 2021,
				Height: Height{
					Value: 150,
					Unit:  centimeters,
				},
				HairColor:  "#fff$fd",
				EyeColor:   "amb",
				PassportId: "860033327",
			},
			false,
		},
		{
			"false if invalid eye color",
			fields{
				BirthYear:      2000,
				IssueYear:      2011,
				ExpirationYear: 2021,
				Height: Height{
					Value: 150,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "invalid",
				PassportId: "860033327",
			},
			false,
		},
		{
			"false if passport id too short",
			fields{
				BirthYear:      2000,
				IssueYear:      2011,
				ExpirationYear: 2021,
				Height: Height{
					Value: 150,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "86003332",
			},
			false,
		},
		{
			"false if passport id too long",
			fields{
				BirthYear:      2000,
				IssueYear:      2011,
				ExpirationYear: 2021,
				Height: Height{
					Value: 150,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "8600333233",
			},
			false,
		},
		{
			"false if passport id not a number",
			fields{
				BirthYear:      2000,
				IssueYear:      2011,
				ExpirationYear: 2021,
				Height: Height{
					Value: 150,
					Unit:  centimeters,
				},
				HairColor:  "#fffffd",
				EyeColor:   "amb",
				PassportId: "860033A23",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Passport{
				BirthYear:      tt.fields.BirthYear,
				IssueYear:      tt.fields.IssueYear,
				ExpirationYear: tt.fields.ExpirationYear,
				Height:         tt.fields.Height,
				HairColor:      tt.fields.HairColor,
				EyeColor:       tt.fields.EyeColor,
				PassportId:     tt.fields.PassportId,
			}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
