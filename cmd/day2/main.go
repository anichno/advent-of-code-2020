package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1ValidatePassword(minCount, maxCount int, char, password string) bool {
	numOccurrence := strings.Count(password, char)

	return numOccurrence >= minCount && numOccurrence <= maxCount
}

func part2ValidatePassword(leftIndex, rightIndex int, char, password string) bool {
	leftChar := string(password[leftIndex-1])
	rightChar := string(password[rightIndex-1])

	return (leftChar != rightChar) && (leftChar == char || rightChar == char)
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re, err := regexp.Compile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	if err != nil {
		log.Fatal(err)
	}

	numValidPart1 := 0
	numValidPart2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result := re.FindStringSubmatch(scanner.Text())
		if len(result) != 5 {
			log.Fatal("invalid input:", scanner.Text())
		}

		minCountStr, maxCountStr, char, password := result[1], result[2], result[3], result[4]
		minCount, _ := strconv.Atoi(minCountStr)
		maxCount, _ := strconv.Atoi(maxCountStr)
		if part1ValidatePassword(minCount, maxCount, char, password) {
			numValidPart1++
		}

		if part2ValidatePassword(minCount, maxCount, char, password) {
			numValidPart2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", numValidPart1)
	fmt.Println("Part 2:", numValidPart2)
}
