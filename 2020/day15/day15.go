package day15

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day15/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, ",")

	fmt.Println("Solution Day 15 - Part 1:", get2020thSpokenNumber(inputs))
	fmt.Println("Solution Day 15 - Part 2:")
}

func get2020thSpokenNumber(inputs []string) string {

	var spoken []int

	return "Not solved yet :("
}
