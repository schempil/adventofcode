package day13

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day13/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	earliest, _ := strconv.Atoi(inputs[0])
	availableLines := strings.Split(inputs[1], ",")

	fmt.Println("Solution Day 13 - Part 1:", getLineWithLeastWaitingTime(earliest, availableLines))
	fmt.Println("Solution Day 13 - Part 2:", getEarliestTimestampOfSequenceDepartures(availableLines))
}

func getLineWithLeastWaitingTime(earliest int, lines []string) int {

	var waitingTime int
	var lineWithLeastWaitingTime int

	for index, line := range lines {

		if line == "x" {
			continue
		}

		currentLine, _ := strconv.Atoi(line)

		currentWaitingTime := currentLine - (earliest % currentLine)

		if index == 0 || currentWaitingTime < waitingTime {
			waitingTime = currentWaitingTime
			lineWithLeastWaitingTime = currentLine
		}

	}

	return waitingTime * lineWithLeastWaitingTime
}

func getEarliestTimestampOfSequenceDepartures(lines []string) int {

	timestamp := 0
	diffToNextTimestamp := 1

	for index, line := range lines {

		if line == "x" {
			continue
		}

		currentLine, _ := strconv.Atoi(line)

		for (timestamp+index)%currentLine > 0 {
			timestamp += diffToNextTimestamp
			fmt.Println("### timestamp", timestamp)
		}

		diffToNextTimestamp *= currentLine
	}

	return timestamp
}
