package day16

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {

	absolutePath, _ := filepath.Abs("./day16/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)

	rules, myTicket, nearbyTickets := getRulesMyTicketAndNearbyTicketsFromInput(text)

	fmt.Println("Solution Day 16 - Part 1:", getErrorRateOfNearbyTickets(rules, nearbyTickets))
	fmt.Println("Solution Day 16 - Part 2:", getDepartureDataFromMyTicket(rules, myTicket, nearbyTickets))
}

type scope struct {
	from int
	to   int
}

type rule struct {
	name   string
	scopes []scope
}

type ticket struct {
	values []int
}

func getErrorRateOfNearbyTickets(rules []rule, nearbyTickets []ticket) int {

	invalidValues := getInvalidValuesOfNearbyTickets(rules, nearbyTickets)
	return getErrorRateOfInvalidValues(invalidValues)
}

func getDepartureDataFromMyTicket(rules []rule, myTicket ticket, nearbyTickets []ticket) int {

	validNearbyTickets := getValidNearbyTickets(rules, nearbyTickets)

	fmt.Println("### validNearbyTickets", validNearbyTickets)

	return -42
}

func getErrorRateOfInvalidValues(invalidValues []int) int {

	errorRate := 0

	for _, invalidValue := range invalidValues {
		errorRate += invalidValue
	}

	return errorRate
}

func getValidNearbyTickets(rules []rule, nearbyTickets []ticket) []ticket {
	var validTickets []ticket

	for _, nearbyTicket := range nearbyTickets {
		if isTicketValid(rules, nearbyTicket) {
			validTickets = append(validTickets, nearbyTicket)
		}
	}

	return validTickets
}

func isTicketValid(rules []rule, ticket ticket) bool {

	for _, value := range ticket.values {
		isValueValid := isValueValidToAnyRule(rules, value)

		if !isValueValid {
			return false
		}

	}

	return true
}

func getInvalidValuesOfNearbyTickets(rules []rule, nearbyTickets []ticket) []int {

	var invalidValues []int

	for _, nearbyTicket := range nearbyTickets {

		for _, ticketValue := range nearbyTicket.values {
			isValid := isValueValidToAnyRule(rules, ticketValue)

			if !isValid {
				invalidValues = append(invalidValues, ticketValue)
			}

		}
	}

	return invalidValues
}

func isValueValidToAnyRule(rules []rule, value int) bool {

	for _, rule := range rules {
		for _, scope := range rule.scopes {
			if value >= scope.from && value <= scope.to {
				return true
			}
		}
	}

	return false
}

func getRulesMyTicketAndNearbyTicketsFromInput(text string) ([]rule, ticket, []ticket) {

	inputs := strings.Split(text, "\n\nyour ticket:")
	inputs2 := strings.Split(inputs[1], "\n\nnearby tickets:")

	ruleStrings := strings.Split(inputs[0], "\n")
	myTicketString := inputs2[0]
	nearbyTicketStrings := strings.Split(inputs2[1], "\n")

	rules := getRulesFromRuleStrings(ruleStrings)
	myTicket := getTicketFromTicketString(myTicketString)
	nearbyTickets := getTicketsFromTicketStrings(nearbyTicketStrings)

	return rules, myTicket, nearbyTickets
}

func getRulesFromRuleStrings(ruleStrings []string) []rule {

	var rules []rule

	for _, ruleString := range ruleStrings {

		newRule := rule{}

		splitAfterName := strings.Split(ruleString, ": ")
		newRule.name = splitAfterName[0]

		scopes := strings.Split(splitAfterName[1], " or ")

		for _, scp := range scopes {
			scopeSplit := strings.Split(scp, "-")

			from, _ := strconv.Atoi(scopeSplit[0])
			to, _ := strconv.Atoi(scopeSplit[1])

			newScope := scope{
				from: from,
				to:   to,
			}

			newRule.scopes = append(newRule.scopes, newScope)
		}

		rules = append(rules, newRule)
	}

	return rules
}

func getTicketFromTicketString(myTicketString string) ticket {

	var myTicket ticket

	for _, valueString := range strings.Split(myTicketString, ",") {

		newValue, _ := strconv.Atoi(strings.Trim(valueString, "\n"))
		myTicket.values = append(myTicket.values, newValue)
	}

	return myTicket
}

func getTicketsFromTicketStrings(nearbyTicketStrings []string) []ticket {

	var nearbyTickets []ticket

	for _, nearbyTicketString := range nearbyTicketStrings {

		if len(nearbyTicketString) < 1 {
			continue
		}

		nearbyTickets = append(nearbyTickets, getTicketFromTicketString(nearbyTicketString))
	}

	return nearbyTickets
}
