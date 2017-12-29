package solutions

import (
	"fmt"
	"math"
	"strconv"
)

// Day03PartOne Spiral Memory (Part One)
// http://adventofcode.com/2017/day/3
func Day03PartOne(input string) string {
	target, _ := strconv.Atoi(input)

	// Calculate grid size based on the nearest odd square.
	gridsize := calculateRoundedUpOddSquare(target)
	d := gridsize - 1
	r := d / 2

	// Rotate the outermost ring until data is at the right side.
	smallerSquare := (gridsize - 2) * (gridsize - 2)
	for target-d > smallerSquare {
		target -= d
	}

	// Calculate the distance away from the center piece.
	centerPiece := smallerSquare + r
	dist := int(math.Abs(float64(target - centerPiece)))

	// Number of steps is r (towards 1) and dist (towards centerPiece)
	steps := r + dist

	return fmt.Sprintf("%d", steps)
}

func calculateRoundedUpOddSquare(x int) int {
	root := 1

	for root*root < x {
		root += 2
	}

	return root
}
