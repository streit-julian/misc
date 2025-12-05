package main

import (
	"strconv"
	"strings"
)

// TODO: Use union find
func (iOld Interval) disjunctions(iNew Interval) []Interval {
	// No intersection
	if iOld.end < iNew.start || iOld.start > iNew.end {
		return []Interval{iNew}
	} else if iOld.start <= iNew.start && iOld.end >= iNew.end { // iOld contains INew
		return []Interval{}
	} else if iOld.start <= iNew.start { //iOld partially contains iNew
		return []Interval{{iOld.end + 1, iNew.end}}
	} else if iOld.end >= iNew.end { // iOld partially contains iNew
		return []Interval{{iNew.start, iOld.start - 1}}
	} else { // iNew contains iOld
		return []Interval{{iNew.start, iOld.start - 1}, {iOld.end + 1, iNew.end}}
	}

}

func part02(input string) int {
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

	disjunctIntervals := []Interval{intervals[0]}
	for _, interval := range intervals[1:] {
		checkIntervals := []Interval{interval}

		for _, disjInterval := range disjunctIntervals {
			newCheckIntervals := []Interval{}
			for _, checkInterval := range checkIntervals {
				newCheckIntervals = append(newCheckIntervals, disjInterval.disjunctions(checkInterval)...)
			}
			checkIntervals = newCheckIntervals
			if len(checkIntervals) == 0 {
				break
			}
		}

		disjunctIntervals = append(disjunctIntervals, checkIntervals...)

	}

	total = 0
	for _, disjInterval := range disjunctIntervals {
		total += disjInterval.end - disjInterval.start + 1
	}

	return total
}
