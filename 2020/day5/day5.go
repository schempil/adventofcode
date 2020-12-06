package day5

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day5/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println("Solution Day 5 - Part 1:", getHighestSeatId(inputs))
	fmt.Println("Solution Day 5 - Part 2:")
}

type seat struct {
	row    int
	column int
}

func getHighestSeatId(inputs []string) int {
	maxSeatId := 0

	for _, input := range inputs {
		seat := getSeat(input)
		seatId := getSeatId(seat)

		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}

	return maxSeatId

}

func getSeatId(seat seat) int {
	return seat.row*8 + seat.column
}

func getSeat(input string) seat {
	row := getRow(input[0:7], 0, 127)
	column := getColumn(input[7:10], 0, 7)

	return seat{row, column}
}

func getRow(input string, low int, high int) int {

	if len(input) == 0 {
		return low
	}

	key := input[0:1]
	leftover := input[1:len(input)]

	if key == "F" {
		high -= int(math.Round((float64(high) - float64(low)) / 2))
	}

	if key == "B" {
		low += int(math.Round((float64(high) - float64(low)) / 2))
	}

	return getRow(leftover, low, high)
}

func getColumn(input string, left int, right int) int {

	if len(input) == 0 {
		return left
	}

	key := input[0:1]
	leftover := input[1:len(input)]

	if key == "L" {
		right -= int(math.Round((float64(right) - float64(left)) / 2))
	}

	if key == "R" {
		left += int(math.Round((float64(right) - float64(left)) / 2))
	}

	return getColumn(leftover, left, right)
}
