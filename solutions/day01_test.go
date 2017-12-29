package solutions

import (
	"github.com/irvinlim/advent-of-code-2017-go/types"
	"testing"
)

var partOneTests = []types.TestCase{
	{Input: "1122", Expected: "3"},
	{Input: "1111", Expected: "4"},
	{Input: "1234", Expected: "0"},
	{Input: "91212129", Expected: "9"},
}

var partTwoTests = []types.TestCase{
	{Input: "1212", Expected: "6"},
	{Input: "1221", Expected: "0"},
	{Input: "123425", Expected: "4"},
	{Input: "123123", Expected: "12"},
	{Input: "12131415", Expected: "4"},
}

func TestDay01PartOne(t *testing.T) {
	RunTestCases(t, Day01PartOne, partOneTests)
	RunInputFileTest(t, Day01PartOne, "input01", "1049")
}

func TestDay01PartTwo(t *testing.T) {
	RunTestCases(t, Day01PartTwo, partTwoTests)
	RunInputFileTest(t, Day01PartTwo, "input01", "1508")
}
