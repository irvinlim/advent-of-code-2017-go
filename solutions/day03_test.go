package solutions

import (
	"github.com/irvinlim/advent-of-code-2017-go/types"
	"testing"
)

func TestDay03PartOne(t *testing.T) {
	var tests = []types.TestCase{
		{Input: "1", Expected: "0"},
		{Input: "12", Expected: "3"},
		{Input: "23", Expected: "2"},
		{Input: "1024", Expected: "31"},
	}

	RunTestCases(t, Day03PartOne, tests)
	RunInputFileTest(t, Day03PartOne, "input03", "371")
}
