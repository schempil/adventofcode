package day6

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day6/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n\n")

	fmt.Println("Solution Day 6 - Part 1:", getAnsweredYesSum(inputs))
	fmt.Println("Solution Day 6 - Part 2:")
}

func getAnsweredYesSum(inputs []string) int {

	sum := 0

	for _, input := range inputs {
		sum += answeredYesPerGroup(input)
	}

	return sum
}

func answeredYesPerGroup(input string) int {

	persons := strings.Split(input, "\n")
	answeredQuestionsMap := make(map[string]bool)

	for _, person := range persons {
		answers := strings.Split(person, "")

		for _, answer := range answers {
			answeredQuestionsMap[answer] = true
		}
	}

	return len(answeredQuestionsMap)
}
