package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	text  byte
	list1 []int
	list2 []int
}

func validate(input string, curRule rule, rules map[int]rule) (bool, int) {
	inputPos := 0
	match := false
	for _, subrule := range curRule.list1 {
		subRule := rules[subrule]
		if subRule.text > 0 {
			match = input[inputPos] == subRule.text
			inputPos++
		} else {
			var numValidated int
			if len(input[inputPos:]) > 0 {
				match, numValidated = validate(input[inputPos:], subRule, rules)
				inputPos += numValidated
			} else {
				match = false
			}
		}

		if !match {
			break
		}
	}

	if !match {
		inputPos = 0
		for _, subrule := range curRule.list2 {
			subRule := rules[subrule]
			if subRule.text > 0 {
				match = input[inputPos] == subRule.text
				inputPos++
			} else {
				var numValidated int
				if len(input[inputPos:]) > 0 {
					match, numValidated = validate(input[inputPos:], subRule, rules)
					inputPos += numValidated
				} else {
					match = false
				}
			}

			if !match {
				break
			}
		}
	}

	return match, inputPos
}

func validate2(input string, ruleList []int, rules map[int]rule) bool {
	if len(input) == 0 && len(ruleList) == 0 {
		return true
	}

	if len(ruleList) == 0 || len(ruleList) > len(input) {
		return false
	}

	match := false
	rule := rules[ruleList[0]]

	if rule.text > 0 {
		if rule.text == input[0] {
			// recurse
			match = validate2(input[1:], ruleList[1:], rules)
		} else {
			return false
		}
	} else {
		//left side
		match = validate2(input, append(rule.list1, ruleList[1:]...), rules)

		//right side
		if !match && len(rule.list2) > 0 {
			match = validate2(input, append(rule.list2, ruleList[1:]...), rules)
		}
	}

	return match
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(map[int]rule)

	// get rules
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		parts := strings.Split(text, " ")
		ruleNum, _ := strconv.Atoi(strings.TrimSuffix(parts[0], ":"))

		list1 := true
		newRule := rule{}
		for _, part := range parts[1:] {
			if strings.Contains(part, "\"") {
				newRule.text = part[1]
				break
			}

			if part == "|" {
				list1 = false
			} else if list1 {
				subrule, _ := strconv.Atoi(part)
				newRule.list1 = append(newRule.list1, subrule)
			} else {
				subrule, _ := strconv.Atoi(part)
				newRule.list2 = append(newRule.list2, subrule)
			}
		}

		rules[ruleNum] = newRule
	}

	// get inputs
	var inputs []string
	for scanner.Scan() {
		text := scanner.Text()
		inputs = append(inputs, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	numMatch := 0
	for _, text := range inputs {
		valid, numValidated := validate(text, rules[0], rules)
		if numValidated < len(text) {
			valid = false
		}
		if valid {
			numMatch++
		}
	}

	fmt.Println("Part 1:", numMatch)

	// Part 2
	rules[8] = rule{0, []int{42}, []int{42, 8}}
	rules[11] = rule{0, []int{42, 31}, []int{42, 11, 31}}

	numMatch = 0
	for _, text := range inputs {
		valid := validate2(text, []int{0}, rules)
		if valid {
			numMatch++
		}
	}

	fmt.Println("Part 2:", numMatch)
}
