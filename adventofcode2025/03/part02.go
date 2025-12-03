package main

import (
	"math"
	"strings"
)

func part02(input string) int {
	splitInput := strings.Split(input, "\n")

	total := 0

	for _, line := range splitInput {
		arr := make([]int, 12)

		// println(len(arr))

		for i, ch := range line {

			lenRemaining := len(line) - i - 1

			n := int(ch - '0')

			for j := max(0, len(arr)-1-lenRemaining); j < len(arr); j++ {

				if n > arr[j] {
					arr[j] = n
					for l := j + 1; l < len(arr); l++ {
						arr[l] = 0
					}
					break
				}
			}
		}

		tot_line := 0

		for i, n := range arr {
			tot_line += (int(math.Pow10(len(arr)-i-1)) * n)
		}

		total += tot_line
	}
	return total
}
