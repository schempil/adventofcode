package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {

	fmt.Println("Hello AventOfCode 2020 - Day #1")

	filepath, _ := filepath.Abs("./day1/input.txt")

	content, readFileError := ioutil.ReadFile(filepath)

	if readFileError != nil {
		log.Fatal(readFileError)
	}

	text := string(content)

	array := strings.Split(text, "\n")

	for _, element := range array {

		value1, _ := strconv.Atoi(element)

		for _, element2 := range array {

			value2, _ := strconv.Atoi(element2)

			if value1+value2 == 2020 {
				fmt.Println("Solution Part1:", value1*value2)
			}

			for _, element3 := range array {

				value3, _ := strconv.Atoi(element3)

				if value1+value2+value3 == 2020 {
					fmt.Println("Solution Part2:", value1*value2*value3)
				}
			}
		}

	}

}
