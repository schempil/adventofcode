package day1

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day1/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	input := strings.Split(text, "\n")

	fmt.Println("Solution Day 1 - Part 1:", findTuple(input))
	fmt.Println("Solution Day 1 - Part 2:", findThirdValueToMatchTupleTo2020(input))
}

func findTuple(input []string) int {
	value1, _ := strconv.Atoi(input[0])
	reducedInput := input[1:]

	for _, element := range reducedInput {
		value2, _ := strconv.Atoi(element)

		if value1+value2 == 2020 {
			return value1 * value2
		}

	}

	return findTuple(reducedInput)
}

type tuple struct {
	val1 int
	val2 int
}

func findTuplesSumLowerThan2020(input []string) []tuple {
	var possibleStarts []tuple

	if len(input) > 1 {
		value1, _ := strconv.Atoi(input[0])
		reducedInput := input[1:]

		for _, element := range reducedInput {
			value2, _ := strconv.Atoi(element)

			if value1+value2 < 2020 {
				possibleStarts = append(possibleStarts, tuple{value1, value2})
			}

		}

		possibleStarts = append(possibleStarts, findTuplesSumLowerThan2020(reducedInput)...)
	}

	return possibleStarts
}

func findThirdValueToMatchTupleTo2020(input []string) int {
	value1, _ := strconv.Atoi(input[0])
	reducedInput := input[1:]

	tuplesSumLowerThan2020 := findTuplesSumLowerThan2020(input)

	for _, tuple := range tuplesSumLowerThan2020 {
		value2 := tuple.val1 + tuple.val2

		if value1+value2 == 2020 {
			return value1 * tuple.val1 * tuple.val2
		}

	}

	return findThirdValueToMatchTupleTo2020(reducedInput)
}
