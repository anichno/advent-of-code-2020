package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	instruction string
	arg         int
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	program := make([]instruction, 0)
	visitedLines := make([]bool, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		instrParts := strings.Split(text, " ")
		instr, arg := instrParts[0], instrParts[1]
		argVal, _ := strconv.Atoi(arg)

		program = append(program, instruction{instr, argVal})
		visitedLines = append(visitedLines, false)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	acc := 0
	for ip := 0; ; {
		if visitedLines[ip] {
			break
		}
		visitedLines[ip] = true
		if program[ip].instruction == "acc" {
			acc += program[ip].arg
		} else if program[ip].instruction == "jmp" {
			ip += program[ip].arg
			continue
		}

		ip++
	}

	fmt.Println("Part 1:", acc)

	for i := 0; i < len(program); i++ {
		if program[i].instruction != "nop" && program[i].instruction != "jmp" {
			continue
		}

		prevInstr := program[i].instruction
		if program[i].instruction == "nop" {
			program[i].instruction = "jmp"
		} else {
			program[i].instruction = "nop"
		}

		visitedLines = make([]bool, len(program))
		acc := 0
		success := false
		for ip := 0; ; {
			if ip == len(program) {
				success = true
				break
			}

			if ip > len(program) || visitedLines[ip] {
				break
			}
			visitedLines[ip] = true
			if program[ip].instruction == "acc" {
				acc += program[ip].arg
			} else if program[ip].instruction == "jmp" {
				ip += program[ip].arg
				continue
			}

			ip++
		}

		if success {
			fmt.Println("Part 2:", acc)
			break
		}

		program[i].instruction = prevInstr
	}
}
