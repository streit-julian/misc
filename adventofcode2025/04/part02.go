package main

import (
	"fmt"
	"strings"
)

func checkDirsAndRemove(pos cord, shape [2]int, arr [][]int) bool {
	n := 0
	for _, dir := range directions {
		checkPos := cord{pos.x + dir.x, pos.y + dir.y}

		if onGrid(checkPos, shape) {
			if arr[checkPos.y][checkPos.x] == 1 {
				n += 1
				if n == 4 {
					return false
				}
			}
		}
	}
	arr[pos.y][pos.x] = 0
	return true
}

func part02(input string) int {
	splitInput := strings.Split(input, "\n")

	total := 0

	arr := make([][]int, len(splitInput))

	for i, line := range splitInput {
		arr[i] = make([]int, len(line))

		for j, ch := range line {
			if ch == '@' {
				arr[i][j] = 1
			} else {
				arr[i][j] = 0
			}
		}
	}

	shape := [2]int{len(arr), len(arr[0])}

	didRemove := true

	for didRemove {
		didRemove = false
		for y, row := range arr {
			for x, val := range row {
				if val == 1 {
					pos := cord{x, y}
					if checkDirsAndRemove(pos, shape, arr) {
						total += 1
						didRemove = true
					}
				}
			}
		}
	}

	fmt.Println(arr)
	return total
}
