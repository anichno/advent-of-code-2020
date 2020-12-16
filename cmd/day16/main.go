package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/set"
)

func resolveFields(fields []string, numFilledFields int, usedFields *set.Set, possibleFields map[int][]string) bool {

	if numFilledFields == len(fields) {
		return true
	}

	minPossible := 1 << 30
	minFieldPos := 0
	for i, field := range fields {
		if field != "" {
			continue
		}

		possible := 0
		for _, posField := range possibleFields[i] {
			if usedFields.Has(posField) {
				continue
			}
			possible++
		}

		if possible < minPossible {
			minPossible = possible
			minFieldPos = i
		}

	}

	for _, posField := range possibleFields[minFieldPos] {
		if usedFields.Has(posField) {
			continue
		}

		fields[minFieldPos] = posField
		usedFields.Insert(posField)
		if resolveFields(fields, numFilledFields+1, usedFields, possibleFields) {
			return true
		}
		usedFields.Remove(posField)
		fields[minFieldPos] = ""
	}

	return false
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ruleRe, err := regexp.Compile("^([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$")
	if err != nil {
		log.Fatal(err)
	}

	// Rules
	rules := make(map[string][4]int)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		matches := ruleRe.FindStringSubmatch(text)
		field, leftLower, leftUpper, rightLower, rightUpper := matches[1], matches[2], matches[3], matches[4], matches[5]
		leftLowerNum, _ := strconv.Atoi(leftLower)
		leftUpperNum, _ := strconv.Atoi(leftUpper)
		rightLowerNum, _ := strconv.Atoi(rightLower)
		rightUpperNum, _ := strconv.Atoi(rightUpper)

		boundsArr := [4]int{leftLowerNum, leftUpperNum, rightLowerNum, rightUpperNum}
		rules[field] = boundsArr
	}

	// My ticket
	scanner.Scan() // advance over "your ticket:"
	scanner.Scan()
	myTicket := make([]int, 0)
	for _, text := range strings.Split(scanner.Text(), ",") {
		val, _ := strconv.Atoi(text)
		myTicket = append(myTicket, val)
	}

	// Nearby tickets
	scanner.Scan() // blank line
	scanner.Scan() // "nearby tickets:"

	nearbyTickets := make([][]int, 0)
	for scanner.Scan() {
		ticket := make([]int, 0)
		for _, text := range strings.Split(scanner.Text(), ",") {
			val, _ := strconv.Atoi(text)
			ticket = append(ticket, val)
		}

		nearbyTickets = append(nearbyTickets, ticket)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ticketScanningErrorRate := 0
	validTickets := make([][]int, 0)
	for _, ticket := range nearbyTickets {
		validTicket := true
		for _, val := range ticket {
			validFieldFound := false
			for _, ruleBounds := range rules {
				leftLower, leftUpper, rightLower, rightUpper := ruleBounds[0], ruleBounds[1], ruleBounds[2], ruleBounds[3]
				if (val >= leftLower && val <= leftUpper) || (val >= rightLower && val <= rightUpper) {
					// valid val
					validFieldFound = true
					break
				} else {
					// invalid val
				}
			}

			if !validFieldFound {
				ticketScanningErrorRate += val
				validTicket = false
			}
		}

		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}

	fmt.Println("Part 1:", ticketScanningErrorRate)

	// Part 2
	validTickets = append(validTickets, myTicket)
	matchingFields := make(map[int][]string)

	for fieldName, fieldBounds := range rules {
		leftLower, leftUpper, rightLower, rightUpper := fieldBounds[0], fieldBounds[1], fieldBounds[2], fieldBounds[3]

		for fieldNum := 0; fieldNum < len(myTicket); fieldNum++ {
			ruleWorks := true
			for _, ticket := range validTickets {
				if (ticket[fieldNum] >= leftLower && ticket[fieldNum] <= leftUpper) || (ticket[fieldNum] >= rightLower && ticket[fieldNum] <= rightUpper) {
					// worked!
				} else {
					// rule failed
					ruleWorks = false
					break
				}
			}

			if ruleWorks {
				matchingFields[fieldNum] = append(matchingFields[fieldNum], fieldName)
			}
		}
	}

	fields := make([]string, len(matchingFields))
	usedFields := set.New()
	numFilled := 0
	for i := 0; i < len(matchingFields); i++ {
		if len(matchingFields[i]) == 1 {
			fields[i] = matchingFields[i][0]
			usedFields.Insert(fields[i])
			numFilled++
		}
	}

	resolveFields(fields, numFilled, usedFields, matchingFields)

	part2Answer := 1
	for i, fieldName := range fields {
		if strings.HasPrefix(fieldName, "departure") {
			part2Answer *= myTicket[i]
		}
	}

	fmt.Println("Part 2:", part2Answer)
}
