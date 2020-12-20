package day20

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {

	absolutePath, _ := filepath.Abs("./day20/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println("Solution Day 20 - Part 1:", inputs[0])
	fmt.Println("Solution Day 20 - Part 2:")
}
