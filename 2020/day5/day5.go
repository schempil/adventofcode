package day5

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day5/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n\n")

	fmt.Println("Solution Day 4 - Part 1:", inputs)
	fmt.Println("Solution Day 4 - Part 2:")
}
