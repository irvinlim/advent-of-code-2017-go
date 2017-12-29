package solutions

import (
	"github.com/irvinlim/advent-of-code-2017-go/types"
	"testing"
)

var testcases = []types.TestCase{
	{Input: "1122", Expected: "3"},
	{Input: "1111", Expected: "4"},
	{Input: "1234", Expected: "0"},
	{Input: "91212129", Expected: "9"},
}

func TestDay01(t *testing.T) {
	RunTestCases(t, Day01, testcases)
	RunInputFileTest(t, Day01, "input01", "1049")
}
