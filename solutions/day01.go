package solutions

import (
	"fmt"
	"strconv"
)

// Day01 Inverse Captcha
// http://adventofcode.com/2017/day/1
func Day01(input string) string {
	var sum int

	for i := 0; i < len(input); i++ {
		var next byte

		// Wrap around the next byte at the last character
		if i == len(input)-1 {
			next = input[0]
		} else {
			next = input[i+1]
		}

		// Increment sum if the next byte matches
		if next == input[i] {
			currInt, _ := strconv.Atoi(fmt.Sprintf("%c", input[i]))
			sum += currInt
		}
	}

	return fmt.Sprintf("%d", sum)
}
