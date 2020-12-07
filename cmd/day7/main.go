package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/set"
)

type subBag struct {
	count int
	child string
}

func findParents(searchColor string, parents map[string][]string, visitedParents *set.Set) {
	for _, parent := range parents[searchColor] {
		if !visitedParents.Has(parent) {
			visitedParents.Insert(parent)
			findParents(parent, parents, visitedParents)
		}
	}
}

func countChildren(curColor string, children map[string][]subBag, bagCounter *int) {
	for _, child := range children[curColor] {
		for x := 0; x < child.count; x++ {
			*bagCounter++
			countChildren(child.child, children, bagCounter)
		}

	}
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bagCategory := make(map[string][]subBag)
	bagParents := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " contain ")
		left, right := split[0], split[1]
		left = strings.TrimSuffix(left, " bags")
		rightParts := strings.Split(right, ", ")
		for _, part := range rightParts {
			entry := strings.SplitN(strings.Split(part, " bag")[0], " ", 2)
			count, _ := strconv.Atoi(entry[0])
			bagCategory[left] = append(bagCategory[left], subBag{count, entry[1]})
			bagParents[entry[1]] = append(bagParents[entry[1]], left)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	visitedParents := set.New()
	findParents("shiny gold", bagParents, visitedParents)
	fmt.Println("Part 1:", visitedParents.Len())

	totalBags := 0
	countChildren("shiny gold", bagCategory, &totalBags)
	fmt.Println("Part 2:", totalBags)
}
