package main

import (
	"strconv"
	"strings"
)

func add(x int, y int) int {
	return x + y
}
func mult(x int, y int) int {
	return x * y
}

func part01(input string) int {
	rows := strings.Split(input, "\n")

	arr := make([][]int, len(rows)-1)

	total := 0

	for y, row := range rows {
		rowSplit := strings.Split(row, " ")

		if y != len(rows)-1 {

			x := 0
			for _, col := range rowSplit {
				if len(col) == 0 {
					continue
				}

				colVal, err := strconv.Atoi(col)
				check(err)

				arr[y] = append(arr[y], colVal)
				x++
			}
		} else {

			x := 0
			for _, col := range rowSplit {
				if len(col) == 0 {
					continue
				}

				var op func(x int, y int) int

				var colTotal int
				switch col {
				case "*":
					op = mult
					colTotal = 1
				case "+":
					op = add
					colTotal = 0
				default:
					panic(col)
				}

				for arry := range len(arr) {
					colTotal = op(colTotal, arr[arry][x])
				}
				total += colTotal
				x += 1

			}
		}
	}
	return total
}
