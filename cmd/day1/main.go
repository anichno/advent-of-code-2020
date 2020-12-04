package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	entries := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		entries = append(entries, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

part1loop:
	for i, val1 := range entries {
		for j, val2 := range entries {
			if i == j {
				continue
			}

			if val1+val2 == 2020 {
				fmt.Println("Part 1:", val1, "+", val2, "=", val1*val2)
				break part1loop
			}
		}
	}

part2loop:
	for i, val1 := range entries {
		for j, val2 := range entries {
			if i == j {
				continue
			}
			for k, val3 := range entries {
				if i == k || j == k {
					continue
				}

				if val1+val2+val3 == 2020 {
					fmt.Println("Part 2:", val1, "+", val2, "+", val3, "=", val1*val2*val3)
					break part2loop
				}
			}
		}
	}

}
