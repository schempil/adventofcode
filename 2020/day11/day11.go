package day11

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day11/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	seatLayout := make([][]string, 0)

	for _, input := range inputs {
		seats := strings.Split(input, "")
		row := make([]string, 0)

		for _, seat := range seats {
			row = append(row, seat)
		}

		seatLayout = append(seatLayout, row)
	}

	seatLayout, isEqual := makePeopleMoveOneStep(seatLayout)

	for !isEqual {
		seatLayout, isEqual = makePeopleMoveOneStep(seatLayout)
	}

	fmt.Println("Solution Day 11 - Part 1:", countOccupiedSeats(seatLayout))
	fmt.Println("Solution Day 11 - Part 2:")
}

func countOccupiedSeats(seatLayout [][]string) int {

	occupied := 0

	for _, row := range seatLayout {
		for _, seat := range row {
			if seat == "#" {
				occupied++
			}
		}
	}

	return occupied
}

func makePeopleMoveOneStep(seatLayout [][]string) (newLayout [][]string, isEqual bool) {

	prediction := make([][]string, 0)

	for rowIndex, _ := range seatLayout {

		var rowPrediction []string

		for seatIndex, _ := range seatLayout[rowIndex] {

			seat := seatLayout[rowIndex][seatIndex]
			seatPrediction := seat

			occupiedSeats := countAdjacentOccupiedSeats(seatLayout, rowIndex, seatIndex)

			if seat == "L" && occupiedSeats == 0 {
				seatPrediction = toggleSeat(seatPrediction)
			}

			if seat == "#" && occupiedSeats >= 4 {
				seatPrediction = toggleSeat(seatPrediction)
			}

			rowPrediction = append(rowPrediction, seatPrediction)

		}

		prediction = append(prediction, rowPrediction)

	}

	return prediction, reflect.DeepEqual(seatLayout, prediction)
}

func toggleSeat(seat string) string {
	if seat == "#" {
		return "L"
	}

	if seat == "L" {
		return "#"
	}

	return "."
}

func countAdjacentOccupiedSeats(seatLayout [][]string, row int, seat int) int {

	adjacentSeats := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	occupiedCount := 0

	for _, adjacentSeat := range adjacentSeats {
		seatValue, isValid := getAdjacentSeat(seatLayout, row+adjacentSeat[0], seat+adjacentSeat[1])

		if !isValid {
			continue
		}

		if seatValue == "#" {
			occupiedCount++
		}

	}

	return occupiedCount
}

func getAdjacentSeat(seatLayout [][]string, row int, seat int) (seatValue string, isValid bool) {

	if row < 0 || row >= len(seatLayout) {
		return "", false
	}

	if seat < 0 || seat >= len(seatLayout[row]) {
		return "", false
	}

	return seatLayout[row][seat], true
}
