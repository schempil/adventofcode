package day4

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
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

	simpleValidatedPassportCount := 0
	strictValidatedPassportCount := 0

	for _, input := range inputs {
		passport := getPassportOfInput(input)
		if isPassportSimpleValidated(passport) {
			simpleValidatedPassportCount++

			if isPassportStrictValidated(passport) {
				strictValidatedPassportCount++
			}

		}
	}

	fmt.Println("Solution Day 4 - Part 1:", simpleValidatedPassportCount)
	fmt.Println("Solution Day 4 - Part 2:", strictValidatedPassportCount)
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

func isPassportSimpleValidated(passport passport) bool {

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

func isPassportStrictValidated(passport passport) bool {

	if len(strconv.Itoa(passport.birthYear)) != 4 || passport.birthYear < 1920 || passport.birthYear > 2002 {
		return false
	}
	if len(strconv.Itoa(passport.issueYear)) != 4 || passport.issueYear < 2010 || passport.issueYear > 2020 {
		return false
	}
	if len(strconv.Itoa(passport.expirationYear)) != 4 || passport.expirationYear < 2020 || passport.expirationYear > 2030 {
		return false
	}
	if !isHeightValid(passport.height) {
		return false
	}
	if !isHairColorValid(passport.hairColor) {
		return false
	}
	if !isEyeColorValid(passport.eyeColor) {
		return false
	}
	if !isPassportIdValid(passport.passportId) {
		return false
	}

	return true
}

func isHeightValid(height string) bool {
	matchedCm, _ := regexp.MatchString(`[0-9][0-9][0-9][c][m]`, height)
	matchedIn, _ := regexp.MatchString(`[0-9][0-9][i][n]`, height)

	if matchedCm == matchedIn {
		return false
	}

	if matchedCm {
		heightValue, _ := strconv.Atoi(height[0:3])

		if heightValue < 150 || heightValue > 193 {
			return false
		}

	}

	if matchedIn {
		heightValue, _ := strconv.Atoi(height[0:2])

		if heightValue < 59 || heightValue > 76 {
			return false
		}
	}

	return true
}

func isHairColorValid(color string) bool {
	matched, _ := regexp.MatchString(`#[a-fA-F0-9]{6}`, color)

	return matched
}

func isEyeColorValid(color string) bool {

	if color == "amb" || color == "blu" || color == "brn" || color == "gry" || color == "grn" || color == "hzl" || color == "oth" {
		return true
	}

	return false
}

func isPassportIdValid(id string) bool {
	matched, _ := regexp.MatchString(`^(\d){9}$`, id)

	return matched
}
