package day21

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {

	absolutePath, _ := filepath.Abs("./day21/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "Tile ")

	fmt.Println("Solution Day 21 - Part 1:", inputs[0])
	fmt.Println("Solution Day 21 - Part 2:")
}
