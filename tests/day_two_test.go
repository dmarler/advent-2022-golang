package tests

import (
	"testing"

	"github.com/dmarler/advent-2022-golang/advent"
)

func TestDayTwo(t *testing.T) {
	advent.DayTwoPartTwo = false
	expected := 15
	actual := advent.DayTwo("day_two_test_data.txt")
	if actual != expected {
		t.Errorf("Day two part one test failed, expected '%d', got: '%d'", expected, actual)
	}

	advent.DayTwoPartTwo = true
	expected = 12
	actual = advent.DayTwo("day_two_test_data.txt")
	if actual != expected {
		t.Errorf("Day two part two test failed, expected '%d', got: '%d'", expected, actual)
	}
}
