package main

import (
	"strings"
)

func part01(input string) int {
	splitInput := strings.Split(input, "\n")

	total := 0

	for _, line := range splitInput {
		var n1 int
		var n2 int
		for i, ch := range line {

			n := int(ch - '0')
			if (i != len(line)-1) && (n > n1) {
				n1 = n
				n2 = 0
			} else if n > n2 {
				n2 = n
			}
		}

		println(n1*10+n2, n1, n2)

		total += n1*10 + n2
	}
	return total
}
