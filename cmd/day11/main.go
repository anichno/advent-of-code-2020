package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const FLOOR = 0
const EMPTY = 1
const OCCUPIED = 2

func numAdj(state [][]byte, x, y int) int {
	type offset struct {
		xdiff int
		ydiff int
	}
	pairs := []offset{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	tot := 0
	for _, pair := range pairs {
		if pair.xdiff == -1 && x == 0 {
			continue
		} else if pair.xdiff == 1 && x == len(state[0])-1 {
			continue
		} else if pair.ydiff == -1 && y == 0 {
			continue
		} else if pair.ydiff == 1 && y == len(state)-1 {
			continue
		}

		if state[y+pair.ydiff][x+pair.xdiff] == OCCUPIED {
			tot++
		}
	}

	return tot
}

func numAdjPart2(state [][]byte, x, y int) int {
	type offset struct {
		xdiff int
		ydiff int
	}
	pairs := []offset{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	tot := 0
	for _, pair := range pairs {
		testx, testy := x+pair.xdiff, y+pair.ydiff
		for testx >= 0 && testx < len(state[0]) && testy >= 0 && testy < len(state) {
			if state[testy][testx] == OCCUPIED {
				tot++
				break
			} else if state[testy][testx] == EMPTY {
				break
			}
			testx += pair.xdiff
			testy += pair.ydiff
		}

	}

	return tot
}

func tick(state *[][]byte, part1 bool, numNeighbors int) bool {
	newGrid := make([][]byte, len(*state))
	changesMade := false

	for y, row := range *state {
		newRow := make([]byte, len(row))
		for x, col := range row {
			if col == FLOOR {
				continue
			}
			var n int
			if part1 {
				n = numAdj(*state, x, y)
			} else {
				n = numAdjPart2(*state, x, y)
			}
			if col == EMPTY && n == 0 {
				changesMade = true
				newRow[x] = OCCUPIED
			} else if col == OCCUPIED && n >= numNeighbors {
				changesMade = true
				newRow[x] = EMPTY
			} else {
				newRow[x] = col
			}
		}
		newGrid[y] = newRow
	}

	*state = newGrid
	return changesMade
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		row := make([]byte, len(text))
		for i, char := range text {
			if char == 'L' {
				row[i] = 1
			}
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	initialGrid := make([][]byte, len(grid))
	for y, row := range grid {
		rowCopy := make([]byte, len(row))
		for x, col := range row {
			rowCopy[x] = col
		}
		initialGrid[y] = rowCopy
	}

	rounds := 0
	for ; tick(&grid, true, 4); rounds++ {
	}

	numOccupied := 0
	for _, row := range grid {
		for _, col := range row {
			if col == OCCUPIED {
				numOccupied++
			}
		}
	}

	fmt.Println("Part 1:", numOccupied)

	for round := 0; tick(&initialGrid, false, 5); round++ {
	}

	numOccupied = 0
	for _, row := range initialGrid {
		for _, col := range row {
			if col == OCCUPIED {
				numOccupied++
			}
		}
	}

	fmt.Println("Part 2:", numOccupied)
}
