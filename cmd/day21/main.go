package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/golang-collections/collections/set"
)

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var recipes [][]string
	allergenSet := make(map[string][]*set.Set)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, "(")
		ingredientString, allergenString := strings.Trim(parts[0], " "), parts[1]
		allergenString = allergenString[9 : len(allergenString)-1]
		ingredients := strings.Split(ingredientString, " ")
		recipes = append(recipes, ingredients)

		for _, allergen := range strings.Split(allergenString, ", ") {
			foodIngredients := set.New()
			for _, ingredient := range ingredients {
				foodIngredients.Insert(ingredient)
			}
			allergenSet[allergen] = append(allergenSet[allergen], foodIngredients)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	allergenIngredientMap := make(map[string]string)
	// reduce ingredients
	for atLeastTwo := true; atLeastTwo; {
		atLeastTwo = false
		for allergen, foods := range allergenSet {
			if _, ok := allergenIngredientMap[allergen]; ok {
				continue
			}
			finalSet := foods[0]
			for _, food := range foods[1:] {
				finalSet = finalSet.Intersection(food)
			}
			finalSet.Do(func(ingredient interface{}) {
				if _, ok := allergenIngredientMap[ingredient.(string)]; ok {
					finalSet.Remove(ingredient)
				}
			})

			if finalSet.Len() == 1 {
				finalSet.Do(func(ingredient interface{}) {
					allergenIngredientMap[allergen] = ingredient.(string)
					allergenIngredientMap[ingredient.(string)] = allergen
				})
			} else {
				atLeastTwo = true
			}
		}

	}

	// Part 1
	tot := 0
	for _, recipe := range recipes {
		for _, ingredient := range recipe {
			if _, ok := allergenIngredientMap[ingredient]; !ok {
				tot++
			}
		}
	}

	fmt.Println("Part 1:", tot)

	// Part 2
	var allergens []string
	for key := range allergenSet {
		allergens = append(allergens, key)
	}

	sort.Strings(allergens)

	var dangerousList strings.Builder
	for i, allergen := range allergens {
		dangerousList.WriteString(allergenIngredientMap[allergen])
		if i != len(allergens)-1 {
			dangerousList.WriteByte(',')
		}
	}

	fmt.Println("Part 2:", dangerousList.String())
}
