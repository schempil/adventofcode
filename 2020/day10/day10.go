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

	adapters := parseAndSort(inputs)

	fmt.Println("Solution Day 10 - Part 1:", getProductOfOneAndThreeDifferences(adapters))

	targetJolt := getTargetJolt(adapters)
	fmt.Println("### targetJolt", targetJolt)

	fmt.Println("Solution Day 10 - Part 2:")
}

func parseAndSort(inputs []string) []int {

	var sequence []int

	for _, input := range inputs {
		parsedInput, _ := strconv.Atoi(input)
		sequence = append(sequence, parsedInput)
	}

	sort.Ints(sequence)

	return sequence
}

func getTargetJolt(inputs []int) int {
	return inputs[len(inputs)-1] + 3
}

func getProductOfOneAndThreeDifferences(inputs []int) int {

	diffOneSum := 0
	diffThreeSum := 1 //One because of the builtIn Adapter

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
