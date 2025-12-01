package main

import (
	"math"
	"strconv"
	"strings"
)

func part02(input string) int {
	trInput := strings.Split(input, "\n")

	init := 50

	total := 0

	for _, line := range trInput {
		num, err := strconv.Atoi(line[1:])
		check(err)

		total += num / 100

		num = int(math.Mod(float64(num), 100))

		isZero := init == 0

		if line[0] == 'R' {
			init += num
		} else {
			init -= num
		}

		if !isZero && (init >= 100 || init <= 0) {
			total += 1
		}

		init += 100
		init = int(math.Mod(float64(init), 100))

	}
	return total
}
