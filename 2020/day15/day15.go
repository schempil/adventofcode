package day15

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day15/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, ",")

	fmt.Println("Solution Day 15 - Part 1:", getNthSpokenNumber(inputs, 2020))
	fmt.Println("Solution Day 15 - Part 2:", getNthSpokenNumber(inputs, 30000000))
}

func getNthSpokenNumber(inputs []string, nth int) int {

	lastSpokenMap := make(map[int][]int)
	lastSpokenNumber := 0
	roundCount := 1

	for _, input := range inputs {
		number, _ := strconv.Atoi(input)

		if len(lastSpokenMap[number]) == 0 {
			lastSpokenMap[number] = append(lastSpokenMap[number], roundCount)
		}

		lastSpokenNumber = number
		roundCount++
	}

	for roundCount <= nth {

		if len(lastSpokenMap[lastSpokenNumber]) == 1 {
			lastSpokenNumber = 0

			if len(lastSpokenMap[lastSpokenNumber]) < 2 {
				lastSpokenMap[lastSpokenNumber] = append(lastSpokenMap[lastSpokenNumber], roundCount)
			} else {
				lastSpokenMap[lastSpokenNumber][0] = lastSpokenMap[lastSpokenNumber][1]
				lastSpokenMap[lastSpokenNumber][1] = roundCount
			}

			roundCount++
			continue
		}

		lastSpokenNumber = lastSpokenMap[lastSpokenNumber][1] - lastSpokenMap[lastSpokenNumber][0]

		if len(lastSpokenMap[lastSpokenNumber]) < 2 {
			lastSpokenMap[lastSpokenNumber] = append(lastSpokenMap[lastSpokenNumber], roundCount)
		} else {
			lastSpokenMap[lastSpokenNumber][0] = lastSpokenMap[lastSpokenNumber][1]
			lastSpokenMap[lastSpokenNumber][1] = roundCount
		}
		roundCount++
		continue

	}

	return lastSpokenNumber
}
