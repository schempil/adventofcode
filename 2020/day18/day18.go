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

	fmt.Println("Solution Day 18 - Part 1:", getSumOfSolvedExpressions(inputs))
	fmt.Println("Solution Day 18 - Part 2:")
}

type resultType int
type symbolType string
type expressionType string

func getSumOfSolvedExpressions(inputs []string) int {

	sum := 0

	for _, input := range inputs {
		sum += int(solveExpression(input))
	}

	return sum
}

func solveExpression(input string) resultType {

	expression := strings.Replace(input, " ", "", -1)

	r := regexp.MustCompile(`[(]\d+((\*|\+)\d+)*[)]`)

	matches := r.FindAllString(expression, -1)

	if len(matches) > 0 {
		match := matches[0]

		matchResult := solveExpressionWithoutParentheses(match[1 : len(match)-1])
		updatedExpression := strings.Replace(expression, match, strconv.Itoa(int(matchResult)), -1)

		return solveExpression(updatedExpression)
	}

	return solveExpressionWithoutParentheses(expression)
}

type operation struct {
	operator string
	value    string
}

func solveExpressionWithoutParentheses(input string) resultType {

	expression := expressionType(strings.Replace(input, " ", "", -1))

	r := regexp.MustCompile(`(\*|\+)\d+`)
	splits := r.Split(string(expression), -1)

	result, _ := strconv.Atoi(splits[0])

	matches := r.FindAllString(string(expression), -1)

	for _, match := range matches {
		operation := operation{
			operator: match[0:1],
			value:    match[1:len(match)],
		}

		result = resultType(result).calculate(operation)

	}

	return resultType(result)
}

func (result resultType) calculate(op operation) int {

	parsedChar, _ := strconv.Atoi(op.value)
	operator := symbolType(op.operator)

	if operator.isAddition() {
		return int(result + resultType(parsedChar))
	}

	return int(result * resultType(parsedChar))
}

func (char symbolType) isOperator() bool {
	if char == "+" || char == "*" {
		return true
	}

	return false
}

func (char symbolType) isAddition() bool {
	if char == "+" {
		return true
	}

	return false
}

func (char symbolType) isMultiplication() bool {
	if char == "*" {
		return true
	}

	return false
}
