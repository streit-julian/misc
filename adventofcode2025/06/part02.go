package main

import (
	"strconv"
	"strings"
)

func part02(input string) int {
	rows := strings.Split(input, "\n")

	arr := make([][]string, len(rows))

	total := 0

	for y, row := range rows {
		for _, col := range row {
			arr[y] = append(arr[y], string(col))
		}
	}

	colSlices := make([][]string, 1)
	colOp := make([]string, 1)

	colSubIx := 0
	colIx := 0
	for x := range len(arr[0]) {
		colEmpty := true

		if len(colSlices) == colIx {
			colSlices = append(colSlices, []string{})
			colOp = append(colOp, "")
		}

		for y := range len(arr) {
			val := arr[y][x]
			if val != " " {
				if val == "*" || val == "+" {
					colOp[colIx] = val
				} else {
					colEmpty = false
					if len(colSlices[colIx]) == colSubIx {
						colSlices[colIx] = append(colSlices[colIx], val)
					} else {
						colSlices[colIx][colSubIx] += val
					}
				}

			}
		}
		colSubIx++

		if colEmpty {
			colIx++
			colSubIx = 0
		}
	}

	for i, colSlice := range colSlices {
		opVal := colOp[i]
		var op func(x int, y int) int
		var opTotal int

		if opVal == "*" {
			op = mult
			opTotal = 1
		} else {
			op = add
			opTotal = 0
		}

		for _, colVal := range colSlice {
			colInt, err := strconv.Atoi(colVal)
			check(err)

			opTotal = op(opTotal, colInt)
		}
		total += opTotal

	}

	return total
}
