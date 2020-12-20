package day20

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {

	absolutePath, _ := filepath.Abs("./day20/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "Tile ")

	fmt.Println("Solution Day 20 - Part 1:", getProductOfRearrangedCornerFragments(inputs))
	fmt.Println("Solution Day 20 - Part 2:")
}

func getProductOfRearrangedCornerFragments(inputs []string) int {

	fragments := parseFragments(inputs)

	var product int

	for _, fragment := range fragments {
		neighbors := fragment.findNeighbors(fragments)

		if len(neighbors) == 2 {

			if product == 0 {
				product = fragment.id
				continue
			}

			product *= fragment.id
		}

	}

	return product
}

func (frgmnt fragment) findNeighbors(fragments []fragment) []fragment {

	var neighbors []fragment

	for _, otherFragment := range fragments {
		if frgmnt.id == otherFragment.id {
			continue
		}

		for _, border := range frgmnt.borders {
			for _, otherBorder := range otherFragment.borders {
				if border == otherBorder || border == reverseString(otherBorder) || otherBorder == reverseString(border) {
					neighbors = append(neighbors, otherFragment)
				}
			}
		}

	}

	return neighbors
}

type fragment struct {
	id      int
	borders []string
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func parseFragments(inputs []string) []fragment {

	var fragments []fragment

	for _, input := range inputs {
		inputLines := strings.Split(input, "\n")

		newFragment := fragment{}
		var leftBorder bytes.Buffer
		var rightBorder bytes.Buffer

		if len(inputLines) == 1 {
			continue
		}

		for lineIndex, inputLine := range inputLines {

			if len(inputLine) < 1 {
				continue
			}

			if lineIndex == 0 {
				newFragment.id, _ = strconv.Atoi(inputLine[0:4])
				continue
			}

			if lineIndex == 1 || lineIndex == 10 {
				newFragment.borders = append(newFragment.borders, inputLine)
			}

			leftBorder.WriteRune(rune(inputLine[0]))
			rightBorder.WriteRune(rune(inputLine[9]))

		}

		newFragment.borders = append(newFragment.borders, leftBorder.String())
		newFragment.borders = append(newFragment.borders, rightBorder.String())

		fragments = append(fragments, newFragment)

	}

	return fragments
}
