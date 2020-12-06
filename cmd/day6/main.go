package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
	groupSet := set.New()
	var groupSetPart2 *set.Set
	sumAnswers := 0
	sumAnswersPart2 := 0
	firstInGroup := true
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			sumAnswers += groupSet.Len()
			sumAnswersPart2 += groupSetPart2.Len()
			groupSet = set.New()
			firstInGroup = true
			continue
		}
		personSet := set.New()
		for _, c := range text {
			groupSet.Insert(c)
			personSet.Insert(c)
		}

		if firstInGroup {
			groupSetPart2 = personSet
			firstInGroup = false
		} else {
			groupSetPart2 = groupSetPart2.Intersection(personSet)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", sumAnswers)
	fmt.Println("Part 2:", sumAnswersPart2)
}
