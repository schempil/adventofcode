package day2

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day2/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	validSledRentalPasswordCount := 0
	validTobogganCorporatePasswordCount := 0

	for _, input := range inputs {
		policy, password := getPolicyAndPasswordFromEntry(input)

		if checkPasswordComplianceAgainstPolicyAccordingToSledRental(password, policy) {
			validSledRentalPasswordCount++
		}

		if checkPasswordComplianceAgainstPolicyAccordingToTobogganCorporate(password, policy) {
			validTobogganCorporatePasswordCount++
		}
	}

	fmt.Println("Solution Day 2 - Part 1:", validSledRentalPasswordCount)
	fmt.Println("Solution Day 2 - Part 2:", validTobogganCorporatePasswordCount)
}

type policy struct {
	from   int
	to     int
	letter string
}

func getPolicyAndPasswordFromEntry(entry string) (policy, string) {
	s := strings.Split(entry, ":")
	policyString := s[0]
	password := strings.Join(strings.Fields(s[1]), "")

	policyParts := strings.Split(policyString, " ")
	fromTo := strings.Split(policyParts[0], "-")
	from, _ := strconv.Atoi(fromTo[0])
	to, _ := strconv.Atoi(fromTo[1])
	policy := policy{from, to, policyParts[1]}

	return policy, password
}

func checkPasswordComplianceAgainstPolicyAccordingToSledRental(password string, policy policy) bool {
	amountOfLetterInPassword := len(strings.Split(password, policy.letter)) - 1

	return amountOfLetterInPassword >= policy.from && amountOfLetterInPassword <= policy.to
}

func checkPasswordComplianceAgainstPolicyAccordingToTobogganCorporate(password string, policy policy) bool {

	passwordLetters := strings.Split(password, "")

	hasFirstPositionTheLetter := passwordLetters[policy.from-1] == policy.letter
	hasSecondPositionTheLetter := passwordLetters[policy.to-1] == policy.letter

	return hasFirstPositionTheLetter != hasSecondPositionTheLetter
}
