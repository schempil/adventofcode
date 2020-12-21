package day21

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Solve() {

	absolutePath, _ := filepath.Abs("./day21/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println("Solution Day 21 - Part 1:", getOccurrencesOfIngredientsWithoutAnyAllergens(inputs))
	fmt.Println("Solution Day 21 - Part 2:")
}

func getOccurrencesOfIngredientsWithoutAnyAllergens(inputs []string) int {

	foods := parseFoods(inputs)

	allergenIngredientMap := make(map[string][]string)

	initializeAllergenIngredientMap(allergenIngredientMap, foods)
	fillAllergenIngredientMap(allergenIngredientMap, foods)

	return countIngredientsWithNoAllergens(getListOfIngredientsWithAllergens(allergenIngredientMap), foods)
}

type food struct {
	ingredients []string
	allergens   []string
}

func countIngredientsWithNoAllergens(listOfIngredientsWithAllergens []string, foods []food) int {

	ingredientsWithNoAllergensCount := 0

	for _, food := range foods {
		for _, ingredient := range food.ingredients {

			if !isStringInList(listOfIngredientsWithAllergens, ingredient) {
				ingredientsWithNoAllergensCount++
			}

		}
	}

	return ingredientsWithNoAllergensCount
}

func isStringInList(list []string, search string) bool {

	for _, stringInList := range list {
		if search == stringInList {
			return true
		}
	}

	return false
}

func getListOfIngredientsWithAllergens(allergenIngredientMap map[string][]string) []string {

	var listOfIngredientsWithAllergens []string

	for allergen, _ := range allergenIngredientMap {
		listOfIngredientsWithAllergens = append(listOfIngredientsWithAllergens, allergenIngredientMap[allergen][0])
	}

	return listOfIngredientsWithAllergens
}

func fillAllergenIngredientMap(allergenIngredientMap map[string][]string, foods []food) {
	for allergen, _ := range allergenIngredientMap {

		for _, food := range foods {
			if food.hasAllergen(allergen) {
				if len(allergenIngredientMap[allergen]) == 0 {
					allergenIngredientMap[allergen] = append(allergenIngredientMap[allergen], food.ingredients...)
					continue
				}

				allergenIngredientMap[allergen] = findSameIngredients(allergenIngredientMap[allergen], food.ingredients)
			}
		}

	}

	for !isEverEntryInMapUnique(allergenIngredientMap) {
		cleanUpAllergenIngredientMap(allergenIngredientMap)
	}
}

func isEverEntryInMapUnique(allergenIngredientMap map[string][]string) bool {
	for allergen, _ := range allergenIngredientMap {
		if len(allergenIngredientMap[allergen]) > 1 {
			return false
		}
	}

	return true
}

func cleanUpAllergenIngredientMap(allergenIngredientMap map[string][]string) {

	for allergen, _ := range allergenIngredientMap {
		if len(allergenIngredientMap[allergen]) == 1 {
			removeIngredientFromMap(allergenIngredientMap[allergen][0], allergenIngredientMap)

		}
	}
}

func removeIngredientFromMap(ingredient string, allergenIngredientMap map[string][]string) {
	for allergen, _ := range allergenIngredientMap {
		if len(allergenIngredientMap[allergen]) != 1 {
			allergenIngredientMap[allergen] = removeStringFromArray(allergenIngredientMap[allergen], ingredient)
		}
	}
}

func removeStringFromArray(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func findSameIngredients(existingIngredients []string, newIngredients []string) []string {

	ingredientsForAllergen := make(map[string]string)

	for _, existingIngredient := range existingIngredients {
		for _, newIngredient := range newIngredients {
			if existingIngredient == newIngredient {
				ingredientsForAllergen[newIngredient] = ""
			}
		}
	}

	var ingredients []string

	for ingredient, _ := range ingredientsForAllergen {
		ingredients = append(ingredients, ingredient)
	}

	return ingredients

}

func (food food) hasAllergen(allergen string) bool {
	for _, foodAllergen := range food.allergens {
		if foodAllergen == allergen {
			return true
		}
	}

	return false
}

func initializeAllergenIngredientMap(allergenIngredientMap map[string][]string, foods []food) {
	for _, food := range foods {
		for _, allergen := range food.allergens {
			allergenIngredientMap[allergen] = []string{}
		}
	}
}

func parseFoods(inputs []string) []food {
	var foods []food

	for _, input := range inputs {

		newFood := food{}

		splitInput := strings.Split(input, " (contains ")

		ingredientsString := splitInput[0]
		allergensString := splitInput[1]

		for _, ingredient := range strings.Split(ingredientsString, " ") {
			newFood.ingredients = append(newFood.ingredients, ingredient)
		}

		allergensStrings := strings.Split(allergensString, ", ")

		for index, allergen := range allergensStrings {

			if index == len(allergensStrings)-1 {
				newFood.allergens = append(newFood.allergens, allergen[0:len(allergen)-1])
				continue
			}

			newFood.allergens = append(newFood.allergens, allergen)
		}

		foods = append(foods, newFood)
	}

	return foods
}
