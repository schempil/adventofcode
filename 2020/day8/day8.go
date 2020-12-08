package day8

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day8/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	acc := 0
	commands := parseCommands(inputs)

	executeCommands(commands, &acc)
	fmt.Println("Solution Day 8 - Part 1:", acc)

	resetAccumulatorAndCommands(&acc, &commands, inputs)
	executeCommandsAndFixEarlyTermination(&commands, &acc, inputs)
	fmt.Println("Solution Day 8 - Part 2:", acc)
}

type command struct {
	id       int
	name     string
	operator string
	value    int
}

func resetAccumulatorAndCommands(acc *int, commands *[]command, inputs []string) {
	*acc = 0
	*commands = parseCommands(inputs)
}

func executeCommandsAndFixEarlyTermination(commands *[]command, acc *int, inputs []string) int {

	for index, command := range *commands {

		resetAccumulatorAndCommands(acc, commands, inputs)
		commandList := *commands

		if command.name == "acc" {
			continue
		}

		if command.name == "jmp" || command.name == "nop" {
			commandList[index] = switchCommand(command)
		}

		earlyTermination := executeCommands(commandList, acc)

		if earlyTermination == false {
			return *acc
		}

	}

	return -10021995
}

func switchCommand(command command) command {
	if command.name == "jmp" {
		command.name = "nop"
		return command
	}

	if command.name == "nop" {
		command.name = "jmp"
		return command
	}

	return command
}

func executeCommands(commands []command, acc *int) (earlyTermination bool) {

	runMap := make(map[int]bool)

	for index := 0; index < len(commands); {
		actualCommand := commands[index]

		if runMap[actualCommand.id] == true {
			return true
		}

		runMap[actualCommand.id] = true

		if actualCommand.name == "nop" {
			index++
		}

		if actualCommand.name == "acc" {
			updateAccumulator(acc, actualCommand.operator, actualCommand.value)
			index++
		}

		if actualCommand.name == "jmp" {
			jump(&index, actualCommand.operator, actualCommand.value)
		}

	}

	return false
}

func jump(currentIndex *int, operator string, value int) {
	if operator == "+" {
		*currentIndex = *currentIndex + value
	}

	if operator == "-" {
		*currentIndex = *currentIndex - value
	}
}

func updateAccumulator(acc *int, operator string, value int) {
	if operator == "+" {
		*acc = *acc + value
	}

	if operator == "-" {
		*acc = *acc - value
	}
}

func parseCommands(inputs []string) []command {

	var commands []command

	for index, input := range inputs {

		value, _ := strconv.Atoi(input[5:])

		commands = append(commands, command{
			id:       index,
			name:     input[0:3],
			operator: input[4:5],
			value:    value,
		})
	}

	return commands
}
