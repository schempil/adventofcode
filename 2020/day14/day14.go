package day14

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day14/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println("Solution Day 14 - Part 1:", solvePart1(inputs))
	fmt.Println("Solution Day 14 - Part 2:")
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
		binaryResult := getMaskedValue(parseIntTo32BitBinaryString(valueAssignment.value), mask)

		storage[valueAssignment.address] = parseBinaryStringToInt(binaryResult)
	}

	return sumFilledStorageIndices(storage)
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

func parseIntTo32BitBinaryString(value int) string {
	return fmt.Sprintf("%036v", strconv.FormatInt(int64(value), 2))
}

func parseBinaryStringToInt(binary string) int {
	value, _ := strconv.ParseInt(binary, 2, 64)
	return int(value)
}
