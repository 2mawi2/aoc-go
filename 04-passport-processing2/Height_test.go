package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHeightFromInput(t *testing.T) {
	// arrange
	input := "183cm"
	// act
	height := NewHeightFromInput(input)
	// assert
	assert.Equal(t, centimeters, height.Unit)
	assert.Equal(t, 183, height.Value)
}

func TestHeight_IsValid(t *testing.T) {
	type fields struct {
		Value int
		Unit  Unit
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"true when valid height",
			fields{Value: 150, Unit: centimeters},
			true,
		},
		{
			"false when cm to low",
			fields{Value: 149, Unit: centimeters},
			false,
		},
		{
			"false when cm to high",
			fields{Value: 194, Unit: centimeters},
			false,
		},
		{
			"false when inch to low",
			fields{Value: 58, Unit: inch},
			false,
		},
		{
			"false when inch to high",
			fields{Value: 77, Unit: inch},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Height{
				Value: tt.fields.Value,
				Unit:  tt.fields.Unit,
			}
			if got := h.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
