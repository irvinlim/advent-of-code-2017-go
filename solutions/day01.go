package solutions

import (
	"fmt"
	"strconv"
)

// Day01PartOne Inverse Captcha (Part One)
// http://adventofcode.com/2017/day/1
func Day01PartOne(input string) string {
	var sum int

	for i := 0; i < len(input); i++ {
		next := input[(i+1)%len(input)]

		if next == input[i] {
			currInt, _ := strconv.Atoi(fmt.Sprintf("%c", input[i]))
			sum += currInt
		}
	}

	return fmt.Sprintf("%d", sum)
}

// Day01PartTwo Inverse Captcha (Part Two)
// http://adventofcode.com/2017/day/1
func Day01PartTwo(input string) string {
	var sum int
	r := len(input) / 2

	for i := 0; i < len(input); i++ {
		next := input[(i+r)%len(input)]

		if next == input[i] {
			currInt, _ := strconv.Atoi(fmt.Sprintf("%c", input[i]))
			sum += currInt
		}
	}

	return fmt.Sprintf("%d", sum)
}
