package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]bool, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		gridRow := make([]bool, len(text))
		for i, char := range text {
			if char == '.' {
				gridRow[i] = false
			} else {
				gridRow[i] = true
			}
		}

		grid = append(grid, gridRow)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	posX, posY := 0, 0
	downX, downY := 3, 1
	numTrees := 0

	for _, row := range grid {
		if row[posX] {
			numTrees++
		}

		posX = (posX + downX) % len(row)
		posY = posY + downY
	}

	fmt.Println("Part 1:", numTrees)

	// Part 2
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2}}

	treesMultiplied := 0
	for i, slope := range slopes {
		posX, posY = 0, 0
		numTrees = 0

		for j, row := range grid {
			if j != posY {
				continue
			}
			if row[posX] {
				numTrees++
			}

			posX = (posX + slope[0]) % len(row)
			posY += slope[1]
		}
		if i == 0 {
			treesMultiplied = numTrees
		} else {
			treesMultiplied *= numTrees
		}
	}

	fmt.Println("Part 2:", treesMultiplied)
}
