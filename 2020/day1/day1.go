package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func Solve() {

	fmt.Println("Hello AventOfCode 2020 - Day #1")

	filepath, pathError := filepath.Abs("./day1/input.txt")

	if pathError != nil {
		log.Fatal(pathError)
	}

	content, readFileError := ioutil.ReadFile(filepath)

	if readFileError != nil {
		log.Fatal(readFileError)
	}

	text := string(content)

	array := strings.Split(text, "\n")

	fmt.Println(len(array))

}
