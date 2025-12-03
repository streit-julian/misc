package main

import (
	"strconv"
	"strings"
)

func getInvalidNum2(numA int, numB int) int {

	tot := 0

	for i := numA; i <= numB; i++ {

		strI := strconv.Itoa(i)

		for step := 1; step <= len(strI)/2; step++ {

			nRep := len(strI) / step

			if nRep*step != len(strI) {
				continue
			}

			key := strI[:step]

			var s strings.Builder

			for range nRep {
				s.WriteString(key)
			}

			if s.String() == strI {
				tot += i
				break
			}

		}
	}
	return tot
}

func part02(input string) int {
	splitInput := strings.Split(input, ",")

	total := 0

	for _, line := range splitInput {

		nums := strings.Split(line, "-")

		num1, err := strconv.Atoi(nums[0])
		check(err)
		num2, err := strconv.Atoi(nums[1])
		check(err)

		total += getInvalidNum2(num1, num2)

	}
	return total
}
