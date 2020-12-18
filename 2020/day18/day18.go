package day18

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {

	absolutePath, _ := filepath.Abs("./day18/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println("Solution Day 18 - Part 1:", getSumOfSolvedExpressions(inputs, solveExpressionLeftToRight))
	fmt.Println("Solution Day 18 - Part 2:", getSumOfSolvedExpressions(inputs, solveExpressionWithPlusPrecedence))
}

type operation struct {
	operator string
	value    string
}

type solveFunction func(input string) int

func getSumOfSolvedExpressions(inputs []string, solve solveFunction) int {

	sum := 0

	for _, input := range inputs {
		sum += int(solveExpression(input, solve))
	}

	return sum
}

func getMatchesInExpression(input string, regex string) (matches []string, expression string) {
	expression = strings.Replace(input, " ", "", -1)
	r := regexp.MustCompile(regex)
	return r.FindAllString(expression, -1), expression
}

func solveExpression(input string, solve solveFunction) int {

	matches, expression := getMatchesInExpression(input, `[(]\d+((\*|\+)\d+)*[)]`)

	if len(matches) > 0 {
		match := matches[0]
		matchResult := solve(match[1 : len(match)-1])
		updatedExpression := strings.Replace(expression, match, strconv.Itoa(matchResult), -1)

		return solveExpression(updatedExpression, solve)
	}

	return solve(expression)
}

func solveExpressionWithPlusPrecedence(input string) int {

	matches, expression := getMatchesInExpression(input, `\d+\+\d+`)

	if len(matches) > 0 {
		match := matches[0]
		matchResult := solveExpressionLeftToRight(match)
		updatedExpression := strings.Replace(expression, match, strconv.Itoa(matchResult), -1)

		return solveExpression(updatedExpression, solveExpressionWithPlusPrecedence)
	}

	return solveExpressionLeftToRight(expression)
}

func solveExpressionLeftToRight(input string) int {

	expression := strings.Replace(input, " ", "", -1)

	r := regexp.MustCompile(`(\*|\+)\d+`)
	splits := r.Split(string(expression), -1)

	result, _ := strconv.Atoi(splits[0])

	matches := r.FindAllString(expression, -1)

	for _, match := range matches {
		operation := operation{
			operator: match[0:1],
			value:    match[1:],
		}

		result = calculate(result, operation)
	}

	return result
}

func calculate(result int, op operation) int {

	parsedValue, _ := strconv.Atoi(op.value)

	if op.operator == "+" {
		return result + parsedValue
	}

	return result * parsedValue
}
