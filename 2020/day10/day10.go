package day10

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day10/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println("Solution Day 10 - Part 1:", getProductOfOneAndThreeDifferences(parseInputsToSequence(inputs)))
	fmt.Println("Solution Day 10 - Part 2:")
}

func parseInputsToSequence(inputs []string) []int {

	var sequence []int

	for _, input := range inputs {
		parsedInput, _ := strconv.Atoi(input)
		sequence = append(sequence, parsedInput)
	}

	return sequence
}

func getProductOfOneAndThreeDifferences(inputs []int) int {

	diffOneSum := 0
	diffThreeSum := 1 //One because of the builtIn Adapter

	sort.Ints(inputs)

	for index, _ := range inputs {

		var diff int

		if index == 0 {
			diff = inputs[index]
		}

		if index > 0 {
			diff = inputs[index] - inputs[index-1]
		}

		if index+1 <= len(inputs) {
			if diff == 1 {
				diffOneSum++
			}

			if diff == 3 {
				diffThreeSum++
			}
		}
	}

	return diffOneSum * diffThreeSum
}
