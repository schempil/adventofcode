package day9

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day9/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	sequence := parseInputsToSequence(inputs)
	invalidNumber, invalidIndex := checkSequenceWithPreambleUntilError(sequence, 25)

	fmt.Println("Solution Day 9 - Part 1:", invalidNumber)
	fmt.Println("Solution Day 9 - Part 2:", sumInvalidTargetFromSequence(sequence, invalidNumber, invalidIndex))
}

func parseInputsToSequence(inputs []string) []int {

	var sequence []int

	for _, input := range inputs {
		parsedInput, _ := strconv.Atoi(input)
		sequence = append(sequence, parsedInput)
	}

	return sequence
}

func sumInvalidTargetFromSequence(sequence []int, invalidTarget int, invalidIndex int) int {

	partialSequence := sequence[0:invalidIndex]
	valuesThatSumUpTarget := canBeSummedUpOutOfAll(partialSequence, invalidTarget)

	return getSumOfMinAndMax(valuesThatSumUpTarget)
}

func getSumOfMinAndMax(values []int) int {
	var min, max int

	for _, value := range values {
		if value > max {
			max = value
		}

		if min == 0 || value < min {
			min = value
		}
	}

	return min + max
}

func checkSequenceWithPreambleUntilError(sequence []int, preambleCount int) (invalidNumber int, invalidIndex int) {

	for index, _ := range sequence {
		if index+preambleCount >= len(sequence) {
			break
		}

		preamble := sequence[index : index+preambleCount]
		target := sequence[index+preambleCount]

		canBe := canBeSummedUpOutOfTwo(preamble, target)

		if canBe == false {
			return target, index + preambleCount
		}
	}

	return -1, -1
}

func canBeSummedUpOutOfTwo(sequence []int, target int) bool {

	canBeSummedUp := false

	for index, valueOne := range sequence {
		for _, valueTwo := range sequence[index+1:] {
			if target == valueOne+valueTwo {
				canBeSummedUp = true
			}
		}
	}

	return canBeSummedUp
}

func canBeSummedUpOutOfAll(sequence []int, target int) []int {

	var sum int
	var addedValues []int

	for startIndex, _ := range sequence {
		addedValues = []int{}
		sum = 0

		partialSequence := sequence[startIndex:]

		for index := 0; sum < target; index++ {
			addedValues = append(addedValues, partialSequence[index])
			sum += partialSequence[index]

			if sum == target {
				return addedValues
			}
		}
	}

	return []int{}
}
