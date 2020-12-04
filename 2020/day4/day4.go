package day4

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type passport struct {
	birthYear      int
	issueYear      int
	expirationYear int
	height         string
	hairColor      string
	eyeColor       string
	passportId     string
	countryId      string
}

func Solve() {
	absolutePath, _ := filepath.Abs("./day4/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n\n")

	validPassportCount := 0

	for _, input := range inputs {
		passport := getPassportOfInput(input)
		if isPassportValid(passport) {
			validPassportCount++
		}
	}

	fmt.Println("Solution Day 4 - Part 1:", validPassportCount)
	fmt.Println("Solution Day 4 - Part 2:")
}

func getPassportOfInput(input string) passport {
	passportDataStrings := strings.Split(strings.ReplaceAll(input, "\n", " "), " ")

	passport := passport{}

	for _, dataString := range passportDataStrings {
		passport = evaluateDataIntoPassport(dataString, passport)
	}

	return passport
}

func evaluateDataIntoPassport(data string, passport passport) passport {

	key := data[0:3]
	value := data[4:len(data)]

	if key == "byr" {
		passport.birthYear, _ = strconv.Atoi(value)
	}
	if key == "iyr" {
		passport.issueYear, _ = strconv.Atoi(value)
	}
	if key == "eyr" {
		passport.expirationYear, _ = strconv.Atoi(value)
	}
	if key == "hgt" {
		passport.height = value
	}
	if key == "hcl" {
		passport.hairColor = value
	}
	if key == "ecl" {
		passport.eyeColor = value
	}
	if key == "pid" {
		passport.passportId = value
	}
	if key == "cid" {
		passport.countryId = value
	}

	return passport
}

func isPassportValid(passport passport) bool {

	if passport.birthYear == 0 {
		return false
	}
	if passport.issueYear == 0 {
		return false
	}
	if passport.expirationYear == 0 {
		return false
	}
	if passport.height == "" {
		return false
	}
	if passport.hairColor == "" {
		return false
	}
	if passport.eyeColor == "" {
		return false
	}
	if passport.passportId == "" {
		return false
	}

	return true
}
