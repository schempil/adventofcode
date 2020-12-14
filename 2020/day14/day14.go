package day14

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day14/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println("Solution Day 14 - Part 1:", solvePart1(inputs))
	fmt.Println("Solution Day 14 - Part 2:", solvePart2(inputs))
}

func solvePart1(inputs []string) int {

	mask := parseBitMaskFromInput(inputs[0])
	storage := make(map[int]int)

	for _, input := range inputs {

		if isInputMaskUpdate(input) {
			mask = parseBitMaskFromInput(input)
			continue
		}

		valueAssignment := parseValueAssignment(input)
		binaryResult := getMaskedValue(parseIntTo32BitBinaryString(valueAssignment.value, 36), mask)

		storage[valueAssignment.address] = parseBinaryStringToInt(binaryResult)
	}

	return sumFilledStorageIndices(storage)
}

func solvePart2(inputs []string) int {

	mask := parseBitMaskFromInput(inputs[0])
	storage := make(map[int]int)

	for _, input := range inputs {

		if isInputMaskUpdate(input) {
			mask = parseBitMaskFromInput(input)
			continue
		}

		valueAssignment := parseValueAssignment(input)
		binaryAddress := parseIntTo32BitBinaryString(valueAssignment.address, 36)

		maskedAddress := getMaskedAddress(binaryAddress, mask)

		addressCombinations := getCombinationsFromFloatingValues(maskedAddress)

		for _, combination := range addressCombinations {
			storage[parseBinaryStringToInt(combination)] = valueAssignment.value

		}

	}

	return sumFilledStorageIndices(storage)
}

func getCombinationsFromFloatingValues(maskedAddress string) (combinations []string) {

	floatingPositions := regexp.MustCompile("X").FindAllStringIndex(maskedAddress, -1)

	for _, valueForFloating := range createValuesForFloatings(len(floatingPositions)) {

		floated := []rune(maskedAddress)

		for i, floatingPosition := range floatingPositions {
			floated[floatingPosition[0]] = []rune(valueForFloating[i : i+1])[0]
		}

		combinations = append(combinations, string(floated))
	}

	return combinations
}

func createValuesForFloatings(floatingCount int) (binaryOptions []string) {
	target := int(math.Pow(2, float64(floatingCount)))

	for i := 0; i < target; i++ {
		binaryOptions = append(binaryOptions, parseIntTo32BitBinaryString(i, floatingCount))
	}

	return binaryOptions
}

func sumFilledStorageIndices(storage map[int]int) int {
	sum := 0

	for _, value := range storage {
		sum += value
	}

	return sum
}

func getMaskedValue(input string, mask string) (result string) {

	for index, _ := range mask {
		char := mask[index : index+1]

		if char == "X" {
			result += input[index : index+1]
			continue
		}

		result += char
	}

	return result
}

func getMaskedAddress(binaryAddress string, mask string) (result string) {

	for index, _ := range mask {
		char := mask[index : index+1]

		if char == "0" {
			result += binaryAddress[index : index+1]
			continue
		}

		result += char
	}

	return result
}

func isInputMaskUpdate(input string) bool {
	return strings.Contains(input, "mask")
}

func parseBitMaskFromInput(input string) string {
	return strings.Split(input, "mask = ")[1]
}

type valueAssignment struct {
	address int
	value   int
}

func parseValueAssignment(input string) valueAssignment {
	start := strings.Split(input, "mem[")[1]
	memoryAddressString := strings.Split(start, "]")[0]
	memoryAddress, _ := strconv.Atoi(memoryAddressString)

	value, _ := strconv.Atoi(strings.Split(input, "] = ")[1])

	return valueAssignment{
		address: memoryAddress,
		value:   value,
	}
}

func parseIntTo32BitBinaryString(value int, padTo int) string {
	binary := strconv.FormatInt(int64(value), 2)
	return fmt.Sprintf("%0"+strconv.Itoa(padTo)+"v", binary)
}

func parseBinaryStringToInt(binary string) int {
	value, _ := strconv.ParseInt(binary, 2, 64)
	return int(value)
}
