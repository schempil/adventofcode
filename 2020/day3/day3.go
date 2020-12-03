package day3

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day3/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	trees := travelRouteWithRationalNumbers(inputs)

	fmt.Println("Solution Day 3 - Part 1:", trees)
	fmt.Println("Solution Day 3 - Part 2:", "NOT SOLVED YET")
}

type position struct {
	x int
	y int
}

func travelRouteWithRationalNumbers(inputs []string) int {

	currentPosition := position{0, 0}
	treeCount := 0

	for index, _ := range inputs {
		currentPosition = position{currentPosition.x + 3, currentPosition.y + 1}

		if index+1 >= len(inputs)-1 {
			break
		}

		inputParts := strings.Split(inputs[currentPosition.y], "")

		if getSymbolAtPositionInRepeatingPattern(inputParts, currentPosition.x) == "#" {
			treeCount++
		}
	}

	return treeCount
}

func getSymbolAtPositionInRepeatingPattern(inputParts []string, positionX int) string {
	return inputParts[positionX%len(inputParts)]
}
