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

	fmt.Println("Solution Day 17 - Part 1:", inputs[0])
	fmt.Println("Solution Day 17 - Part 2:")
}
