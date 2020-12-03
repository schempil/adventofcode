package day3

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type position struct {
	x int
	y int
}

type route struct {
	x int
	y int
}

func Solve() {
	absolutePath, _ := filepath.Abs("./day3/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	oneOneRoute := route{1, 1}
	threeOneRoute := route{3, 1}
	fiveOneRoute := route{5, 1}
	sevenOneRoute := route{7, 1}
	oneTwoRoute := route{1, 2}

	routes := []route{oneOneRoute, threeOneRoute, fiveOneRoute, sevenOneRoute, oneTwoRoute}

	fmt.Println("Solution Day 3 - Part 1:", travelRoute(inputs, threeOneRoute))
	fmt.Println("Solution Day 3 - Part 2:", travelRoutesAndMultiply(inputs, routes))
}

func travelRoute(inputs []string, route route) int {
	currentPosition := position{0, 0}
	treeCount := 0

	for range inputs {
		currentPosition = position{currentPosition.x + route.x, currentPosition.y + route.y}

		if currentPosition.y >= len(inputs)-1 {
			return treeCount
		}

		inputParts := strings.Split(inputs[currentPosition.y], "")

		if getSymbolAtPositionInRepeatingPattern(inputParts, currentPosition.x) == "#" {
			treeCount++
		}
	}

	return treeCount
}

func travelRoutesAndMultiply(inputs []string, routes []route) int {

	product := 0

	for index, route := range routes {
		treesInRoute := travelRoute(inputs, route)

		if index < 1 {
			product = treesInRoute
			continue
		}

		product = product * treesInRoute
	}

	return product
}

func getSymbolAtPositionInRepeatingPattern(inputParts []string, positionX int) string {
	return inputParts[positionX%len(inputParts)]
}
