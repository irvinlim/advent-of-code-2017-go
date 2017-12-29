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
	for _, testcase := range testcases {
		actual := Day01(testcase.Input)

		if actual != testcase.Expected {
			t.Errorf("Input %s: Expected %s, got %s instead", testcase.Input, testcase.Expected, actual)
		}
	}
}
