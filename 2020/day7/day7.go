package day7

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day7/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, ".")

	fmt.Println("Solution Day 7 - Part 1:", inputs[0])
	fmt.Println("Solution Day 7 - Part 2:")
}
