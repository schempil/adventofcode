package day17

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {

	absolutePath, _ := filepath.Abs("./day17/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println("Solution Day 17 - Part 1:", startBootProcessCycles(inputs, 6))
	fmt.Println("Solution Day 17 - Part 2:")
}

type coordinate struct {
	x int
	y int
	z int
}

type coordinates []coordinate

type dimension map[coordinate]string

func startBootProcessCycles(inputs []string, cycles int) int {

	pocketDimension := make(dimension)
	pocketDimension.initializeFromInput(inputs)

	for i := 0; i < cycles; i++ {
		pocketDimension = pocketDimension.doCycle()
	}

	return pocketDimension.getActiveStateCount()
}

func (pocketDimension dimension) doCycle() dimension {

	newPocketDimension := make(dimension)

	for coordinate, value := range pocketDimension {
		newPocketDimension[coordinate] = value
	}

	for coord, _ := range pocketDimension {

		neighbors := coord.getNeighbors()
		pocketDimension.addMissingNeighbors(neighbors)
	}

	for coord, _ := range pocketDimension {

		neighbors := coord.getNeighbors()

		activeNeighbors := neighbors.getActiveNeighbors(pocketDimension)

		newPocketDimension[coord] = pocketDimension.getUpdatedStateOfCoordinate(coord, activeNeighbors)
	}

	return newPocketDimension
}

func (pocketDimension dimension) getActiveStateCount() int {

	activeStateCount := 0

	for _, coordValue := range pocketDimension {
		if coordValue == "#" {
			activeStateCount++
		}
	}

	return activeStateCount
}

func (pocketDimension dimension) addMissingNeighbors(neighbors coordinates) {

	for _, neighbor := range neighbors {
		if pocketDimension[neighbor] != "#" && pocketDimension[neighbor] != "." {
			pocketDimension[neighbor] = "."
		}
	}
}

func (pocketDimension dimension) getUpdatedStateOfCoordinate(coord coordinate, activeNeighbors coordinates) string {
	if coord.isActiveInDimension(pocketDimension) {
		if len(activeNeighbors) == 2 || len(activeNeighbors) == 3 {
			return "#"
		}

		return "."
	}

	if len(activeNeighbors) == 3 {
		return "#"
	}

	return "."
}

func (coords coordinates) getActiveNeighbors(pocketDimension dimension) coordinates {

	var activeNeighbors coordinates

	for _, neighbor := range coords {

		if neighbor.isActiveInDimension(pocketDimension) {
			activeNeighbors = append(activeNeighbors, neighbor)
		}
	}

	return activeNeighbors
}

func (coord coordinate) isActiveInDimension(pocketDimension dimension) bool {
	return pocketDimension[coord] == "#"
}

func (pocketDimension dimension) getStateOfCoordinate(coord coordinate) string {
	if pocketDimension[coord] == "#" {
		return "#"
	}

	return "."
}

func (coord coordinate) getNeighbors() coordinates {
	directions := []coordinate{
		{x: 0, y: 0, z: 1},
		{x: 0, y: 0, z: -1},
		{x: 0, y: 1, z: 0},
		{x: 0, y: 1, z: 1},
		{x: 0, y: 1, z: -1},
		{x: 0, y: -1, z: 0},
		{x: 0, y: -1, z: 1},
		{x: 0, y: -1, z: -1},
		{x: 1, y: 0, z: 0},
		{x: 1, y: 0, z: 1},
		{x: 1, y: 0, z: -1},
		{x: 1, y: 1, z: 0},
		{x: 1, y: 1, z: 1},
		{x: 1, y: 1, z: -1},
		{x: 1, y: -1, z: 0},
		{x: 1, y: -1, z: 1},
		{x: 1, y: -1, z: -1},
		{x: -1, y: 0, z: 0},
		{x: -1, y: 0, z: 1},
		{x: -1, y: 0, z: -1},
		{x: -1, y: 1, z: 0},
		{x: -1, y: 1, z: 1},
		{x: -1, y: 1, z: -1},
		{x: -1, y: -1, z: 0},
		{x: -1, y: -1, z: 1},
		{x: -1, y: -1, z: -1},
	}

	var neighbors []coordinate

	for _, direction := range directions {
		neighbors = append(neighbors, coordinate{
			x: coord.x + direction.x,
			y: coord.y + direction.y,
			z: coord.z + direction.z,
		})
	}

	return neighbors
}

func (pocketDimension dimension) initializeFromInput(inputs []string) {

	for rowIndex, row := range inputs {
		for colIndex, col := range row {

			pocketDimension[coordinate{
				x: colIndex,
				y: rowIndex,
				z: 0,
			}] = string(col)
		}
	}
}
