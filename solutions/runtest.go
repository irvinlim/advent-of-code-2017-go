package solutions

import (
	"fmt"
	"github.com/irvinlim/advent-of-code-2017-go/types"
	"io/ioutil"
	"strings"
	"testing"
)

// RunTestCases is a generic test runner for running a given set of TestCases.
func RunTestCases(t *testing.T, f func(string) string, testcases []types.TestCase) {
	for _, testcase := range testcases {
		runTestCase(t, f, testcase)
	}
}

// RunInputFileTest is a generic test runner for running a test case from a file.
func RunInputFileTest(t *testing.T, f func(string) string, filename string, expected string) {
	data, err := ioutil.ReadFile(fmt.Sprintf("../inputs/%s", filename))

	if err != nil {
		t.Errorf("Could not open file: %s", filename)
		return
	}

	input := strings.TrimSpace(string(data))

	testcase := types.TestCase{Input: input, Expected: expected}
	runTestCase(t, f, testcase)
}

func runTestCase(t *testing.T, f func(string) string, testcase types.TestCase) {
	actual := f(testcase.Input)

	if actual != testcase.Expected {
		t.Errorf("Input %s: Expected %s, got %s instead", testcase.Input, testcase.Expected, actual)
	}
}
