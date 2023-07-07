package tests

import (
	"testing"

	"github.com/dmarler/advent-2022-golang/advent"
)

func TestDayOne(t *testing.T) {
	expected := 24000
	actual := advent.DayOne("day_one_test_data.txt")
	if actual != expected {
		t.Errorf("Day one test failed, expected: '%d', got: '%d'", expected, actual)
	}
}
