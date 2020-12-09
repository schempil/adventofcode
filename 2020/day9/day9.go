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

	fmt.Println("Solution Day 9 - Part 1:", checkSequenceWithPreambleUntilError(sequence, 25))
	fmt.Println("Solution Day 9 - Part 2:")
}

func parseInputsToSequence(inputs []string) []int {

	var sequence []int

	for _, input := range inputs {
		parsedInput, _ := strconv.Atoi(input)
		sequence = append(sequence, parsedInput)
	}

	return sequence
}

func checkSequenceWithPreambleUntilError(sequence []int, preambleCount int) int {

	for index, _ := range sequence {
		if index+preambleCount >= len(sequence) {
			break
		}

		preamble := sequence[index : index+preambleCount]
		target := sequence[index+preambleCount]

		canBe := canBeSummedUp(preamble, target)

		if canBe == false {
			return target
		}
	}

	return -1
}

func canBeSummedUp(sequence []int, target int) bool {

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
