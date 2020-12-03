package day3

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day3/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println(inputs)

	fmt.Println("Solution Day 3 - Part 1:")
	fmt.Println("Solution Day 3 - Part 2:")
}
