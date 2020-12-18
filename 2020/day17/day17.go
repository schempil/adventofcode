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

	fmt.Println("Solution Day 17 - Part 1:", startBootProcessCycles(inputs, 6, false))
	fmt.Println("Solution Day 17 - Part 2:", startBootProcessCycles(inputs, 6, true))
}

type coordinate struct {
	x int
	y int
	z int
	w int
}

type coordinates []coordinate

type dimension map[coordinate]string

func startBootProcessCycles(inputs []string, cycles int, fourthDimension bool) int {

	pocketDimension := make(dimension)
	pocketDimension.initializeFromInput(inputs, fourthDimension)

	for i := 0; i < cycles; i++ {
		pocketDimension = pocketDimension.doCycle(fourthDimension)
	}

	return pocketDimension.getActiveStateCount()
}

func (pocketDimension dimension) doCycle(fourthDimension bool) dimension {

	newPocketDimension := make(dimension)

	for coord, value := range pocketDimension {
		newPocketDimension[coord] = value
		neighbors := coord.getNeighbors(fourthDimension)
		pocketDimension.addMissingNeighbors(neighbors)
	}

	for coord, _ := range pocketDimension {

		neighbors := coord.getNeighbors(fourthDimension)
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

func (coord coordinate) getNeighbors(fourthDimension bool) coordinates {

	var fromW int
	var toW int

	if fourthDimension {
		fromW = -1
		toW = 1
	}

	var neighbors []coordinate

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := fromW; w <= toW; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}

					direction := coordinate{
						x: x, y: y, z: z,
					}

					if fourthDimension {
						direction.w = w
					}

					newCoordinate := coordinate{
						x: coord.x + direction.x,
						y: coord.y + direction.y,
						z: coord.z + direction.z,
					}

					if fourthDimension {
						newCoordinate.w = coord.w + direction.w
					}

					neighbors = append(neighbors, newCoordinate)

				}
			}
		}
	}

	return neighbors
}

func (pocketDimension dimension) initializeFromInput(inputs []string, fourthDimension bool) {

	for rowIndex, row := range inputs {
		for colIndex, col := range row {

			newCoordinate := coordinate{
				x: colIndex,
				y: rowIndex,
				z: 0,
			}

			if fourthDimension {
				newCoordinate.w = 0
			}

			pocketDimension[newCoordinate] = string(col)
		}
	}
}
