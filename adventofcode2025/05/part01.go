package main

import (
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func (i Interval) contains(x int) bool {
	return i.start <= x && x <= i.end
}

func part01(input string) int {
	splitInput := strings.Split(input, "\n")

	total := 0

	var intervals = []Interval{}

	curLineIx := 0
	for _, line := range splitInput {
		curLineIx += 1
		if len(line) == 0 {
			break
		}

		spLine := strings.Split(line, "-")

		left, err := strconv.Atoi(spLine[0])
		check(err)
		right, err := strconv.Atoi(spLine[1])
		check(err)

		intervals = append(intervals, Interval{left, right})
	}

	for _, line := range splitInput[curLineIx:] {
		if len(line) == 0 {
			break
		}
		lineNum, err := strconv.Atoi(line)
		check(err)

		for _, interval := range intervals {
			if interval.contains(lineNum) {
				total += 1
				break
			}
		}

	}

	return total
}
