package main

import (
	"math"
	"strconv"
	"strings"
)

func part01(input string) int {
	trInput := strings.Split(input, "\n")

	init := 50

	total := 0

	for _, line := range trInput {
		num, err := strconv.Atoi(line[1:])
		check(err)

		if line[0] == 'R' {
			init += num
		} else {
			init -= num
		}

		init += 100
		init = int(math.Mod(float64(init), 100))

		if init == 0 {
			total += 1
		}
	}
	return total
}
