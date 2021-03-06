package solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Day02PartOne Corruption Checksum (Part One)
// http://adventofcode.com/2017/day/2
func Day02PartOne(input string) string {
	spreadsheet := parseSpreadsheet(input)
	checksum := 0

	for _, row := range spreadsheet {
		max := 0
		min := math.MaxInt32

		for _, col := range row {
			max = int(math.Max(float64(col), float64(max)))
			min = int(math.Min(float64(col), float64(min)))
		}

		checksum += max - min
	}

	return fmt.Sprintf("%d", checksum)
}

// Day02PartTwo Corruption Checksum (Part Two)
// http://adventofcode.com/2017/day/2
func Day02PartTwo(input string) string {
	spreadsheet := parseSpreadsheet(input)
	sum := 0

	for _, row := range spreadsheet {
		for _, x := range row {
			for _, y := range row {
				if x > y && x%y == 0 {
					sum += x / y
				}
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}

func parseSpreadsheet(input string) [][]int {
	var matrix [][]int

	rows := strings.Split(input, "\n")
	matrix = make([][]int, len(rows))

	for i, row := range rows {
		columns := strings.Fields(row)
		matrix[i] = make([]int, len(columns))

		for j, col := range columns {
			parsed, _ := strconv.Atoi(col)
			matrix[i][j] = parsed
		}
	}

	return matrix
}
