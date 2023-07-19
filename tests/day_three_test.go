package tests

import (
	"testing"

	"github.com/dmarler/advent-2022-golang/advent"
)

func TestDayThree(t *testing.T) {
	expected := 157
	actual := advent.DayThree("day_three_test_data.txt")
	if actual != expected {
		t.Errorf("Day three test failed, expected: '%d', got: '%d'", expected, actual)
	}
}
