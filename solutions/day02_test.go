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
