package main

import (
	"testing"
)

func TestGenderMultiplier(t *testing.T) {
	for _, test := range []struct {
		ageGroup AgeGroup
		gender   Gender
		expect   float64
	}{
		{LTE35, Male, 0.9},
		{LTE35, Female, 0.85},
		{GT35, Male, 0.95},
	} {
		t.Run("", func(t *testing.T) {
			received := genderMultiplier(test.ageGroup, test.gender)

			if test.expect != received {
				t.Errorf("expected %v, received %v", test.expect, received)
			}
		})
	}
}

func TestClassify(t *testing.T) {
	for _, test := range []struct {
		input  float64
		expect string
	}{
		{0, "thin"},
		{84.9, "thin"},
		{9999999999999.9, "fat"},
	} {
		t.Run("", func(t *testing.T) {
			received := classify(test.input)

			if test.expect != received {
				t.Errorf("expected %v, received %v", test.expect, received)
			}
		})
	}
}
