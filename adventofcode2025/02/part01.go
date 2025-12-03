package main

import (
	"math"
	"strconv"
	"strings"
)

func isEvenLength(n string) bool {
	return math.Mod(float64(len(n)), 2) == 0
}

func getInvalidNum(numA int, numB int) int {

	strLengthHalf := len(strconv.Itoa(numA)) / 2

	tot := 0

	for i := numA; i <= numB; i++ {

		strI := strconv.Itoa(i)

		strIA := strI[:strLengthHalf]
		strIB := strI[strLengthHalf:]

		if strIA == strIB {
			// println(strIA, strIB, strI, i, tot)
			tot += i
		}
	}
	return tot
}

func part01(input string) int {
	splitInput := strings.Split(input, ",")

	total := 0

	for _, line := range splitInput {

		nums := strings.Split(line, "-")

		num1, err := strconv.Atoi(nums[0])
		check(err)
		num2, err := strconv.Atoi(nums[1])
		check(err)

		if len(nums[0]) == len(nums[1]) {
			if isEvenLength(nums[0]) {
				total += getInvalidNum(num1, num2)
			}
		} else {
			if isEvenLength(nums[0]) {
				println(num1, int(math.Pow10(len(nums[0])))-1)
				total += getInvalidNum(num1, int(math.Pow10(len(nums[0])))-1)
			}
			if isEvenLength(nums[1]) {
				println(int(math.Pow10(len(nums[1])-1)), num2)
				total += getInvalidNum(int(math.Pow10(len(nums[1])-1)), num2)
			}

			if len(nums[1])-len(nums[0]) > 1 {
				println(line)
				panic(err)
			}
			// TODO: If length differs by more than 1 then we have to check ranges inbetween
		}

		// println(line, total)

	}
	return total
}
