package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solvePart1(line []byte) (int, int) {
	var leftBuilder strings.Builder
	var rightBuilder strings.Builder
	leftParsed := false
	rightParsed := false

	left := 0
	right := 0
	operator := byte('+')
	i := 0
	for end := false; i < len(line) && !end; i++ {
		char := line[i]

		if char >= '0' && char <= '9' {
			if !leftParsed {
				leftBuilder.WriteByte(char)
			} else {
				rightBuilder.WriteByte(char)
			}

		} else {
			if !leftParsed && leftBuilder.Len() > 0 {
				newLeft, err := strconv.Atoi(leftBuilder.String())
				if err == nil {
					left = newLeft
				}
				leftParsed = true
			} else if rightBuilder.Len() > 0 {
				newRight, err := strconv.Atoi(rightBuilder.String())
				if err == nil {
					right = newRight
				}
				rightParsed = true
				rightBuilder.Reset()
			}
		}
		if char == ' ' || char == ')' {

			if char == ')' {
				end = true
			}
		} else if char == '+' || char == '*' {
			operator = char
		} else if char == '(' {
			val, numParsed := solvePart1(line[i+1:])
			if !leftParsed {
				left = val
				leftParsed = true
			} else {
				right = val
				rightParsed = true
			}
			i += numParsed
		}

		if leftParsed && rightParsed {
			rightParsed = false
			if operator == '+' {
				left = left + right
			} else if operator == '*' {
				left = left * right
			} else {
				log.Fatal(left, operator, right, line)
			}
		}
	}

	return left, i
}

func addParenthesis(line string) string {
	for i := 0; i < len(line); i++ {
		char := line[i]
		if char == '+' {
			var depth int
			// find left token
			leftEnd := i - 2
			var leftStart int
			for leftStart, depth = leftEnd, 0; leftStart >= 0; leftStart-- {
				if line[leftStart] == ')' {
					depth++
				} else if line[leftStart] == '(' {
					depth--
				}

				if depth <= 0 && (line[leftStart] == ' ' || line[leftStart] == '(') {
					if line[leftStart] == ' ' {
						leftStart++
					}
					break
				}
			}
			if leftStart == -1 {
				leftStart = 0
			}

			left := line[leftStart : leftEnd+1]

			// find right token
			rightStart := i + 2
			var rightEnd int
			for rightEnd, depth = rightStart, 0; rightEnd < len(line); rightEnd++ {
				if line[rightEnd] == '(' {
					depth++
				} else if line[rightEnd] == ')' {
					depth--
				}

				if (line[rightEnd] == ' ' && depth <= 0) || (line[rightEnd] == ')' && depth <= 0) {
					if line[rightEnd] == ' ' {
						rightEnd--
					}
					break
				}
			}
			if rightEnd >= len(line) {
				rightEnd = len(line) - 1
			}

			right := line[rightStart : rightEnd+1]

			var newLine strings.Builder
			newLine.WriteString(line[:leftStart])
			newLine.WriteByte('(')
			newLine.WriteString(left)
			newLine.WriteString(" + ")
			newLine.WriteString(right)
			newLine.WriteByte(')')
			newLine.WriteString(line[rightEnd+1:])

			line = newLine.String()
			i++
		}
	}

	return line
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumPart1 := 0
	sumPart2 := 0
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		problem := []byte(text)
		problem = append(problem, ' ')
		answer, _ := solvePart1(problem)
		sumPart1 += answer

		part2Problem := []byte(addParenthesis(text))
		part2Problem = append(part2Problem, ' ')
		answer, _ = solvePart1(part2Problem)
		sumPart2 += answer
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", sumPart1)
	fmt.Println("Part 2:", sumPart2)
}
