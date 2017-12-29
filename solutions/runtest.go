package solutions

import (
	"github.com/irvinlim/advent-of-code-2017-go/types"
	"testing"
)

// RunTest is a generic test runner for the project.
func RunTest(t *testing.T, f func(string) string, testcases []types.TestCase) {
	for _, testcase := range testcases {
		actual := f(testcase.Input)

		if actual != testcase.Expected {
			t.Errorf("Input %s: Expected %s, got %s instead", testcase.Input, testcase.Expected, actual)
		}
	}
}
