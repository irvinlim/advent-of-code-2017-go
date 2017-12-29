package solutions

import (
	"github.com/irvinlim/advent-of-code-2017-go/types"
	"testing"
)

func TestDay02PartOne(t *testing.T) {
	var tests = []types.TestCase{
		{Input: "5 1 9 5\n7 5 3\n2 4 6 8", Expected: "18"},
	}

	RunTestCases(t, Day02PartOne, tests)
	RunInputFileTest(t, Day02PartOne, "input02", "44670")
}

func TestDay02PartTwo(t *testing.T) {
	var tests = []types.TestCase{
		{Input: "5 9 2 8\n9 4 7 3\n3 8 6 5", Expected: "9"},
	}

	RunTestCases(t, Day02PartTwo, tests)
	RunInputFileTest(t, Day02PartTwo, "input02", "285")
}
