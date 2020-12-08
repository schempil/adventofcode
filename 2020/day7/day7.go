package day7

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {
	absolutePath, _ := filepath.Abs("./day7/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	nameToBag := parseBags(inputs)
	shinyGoldBag := nameToBag["shiny gold"]

	fmt.Println("Solution Day 7 - Part 1:", countUniqueParentBags(shinyGoldBag))
	fmt.Println("Solution Day 7 - Part 2:", countChildrenBags(shinyGoldBag))
}

type Bag struct {
	name        string
	parents     []*Bag
	children    []*Bag
	bagToAmount map[string]int
}

func countUniqueParentBags(bag *Bag) int {

	fmt.Println("### bag", bag)

	parentBags := collectParentBags(bag, make(map[string]*Bag))

	return len(parentBags)
}

func collectParentBags(bag *Bag, parentBagMap map[string]*Bag) map[string]*Bag {
	for _, parentBag := range bag.parents {
		parentBagMap[parentBag.name] = parentBag
		parentBagMap = collectParentBags(parentBag, parentBagMap)
	}

	return parentBagMap
}

func countChildrenBags(bag *Bag) int {
	childrenCount := 0

	for _, childBag := range bag.children {
		childAmount := bag.bagToAmount[childBag.name]
		childrenCount += childAmount + countChildrenBags(childBag)*childAmount
	}

	return childrenCount
}

func parseBags(lines []string) (nameToBag map[string]*Bag) {
	nameToBag = make(map[string]*Bag)

	for _, line := range lines {
		line = strings.TrimSuffix(line, ".")
		parentParts := strings.Split(line, " bags contain ")
		parentBagName := parentParts[0]
		tail := parentParts[1]
		parentBag := getBag(nameToBag, parentBagName)
		containedBags := strings.Split(tail, ", ")

		for _, containedBag := range containedBags {
			if containedBag == "no other bags" {
				continue
			}

			childParts := strings.Fields(containedBag)
			amount, _ := strconv.Atoi(childParts[0])
			childBagName := strings.Join(childParts[1:], " ")
			childBagName = normalizeBagName(childBagName)
			childBag := getBag(nameToBag, childBagName)

			parentBag.children = append(parentBag.children, childBag)
			childBag.parents = append(childBag.parents, parentBag)

			parentBag.bagToAmount[childBagName] = amount
		}
	}

	return nameToBag
}

func getBag(nameToBag map[string]*Bag, bagName string) *Bag {
	bag, hasBag := nameToBag[bagName]

	if !hasBag {
		bag = &Bag{name: bagName, bagToAmount: make(map[string]int)}
		nameToBag[bagName] = bag
	}

	return bag
}

func normalizeBagName(bagName string) string {
	bagName = strings.TrimSuffix(bagName, " bags")
	bagName = strings.TrimSuffix(bagName, " bag")

	return bagName
}
