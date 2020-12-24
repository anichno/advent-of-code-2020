package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

func numNeighbors(p point, grid map[point]bool) int {
	offsets := [][]int{{2, 0}, {1, 1}, {-1, 1}, {-2, 0}, {-1, -1}, {1, -1}}

	num := 0
	for _, offset := range offsets {
		xOffset, yOffset := offset[0], offset[1]
		if black, ok := grid[point{p.x + xOffset, p.y + yOffset}]; ok {
			if black {
				num++
			}
		}
	}

	return num
}

func allAdjacent(p point) []point {
	offsets := [][]int{{2, 0}, {1, 1}, {-1, 1}, {-2, 0}, {-1, -1}, {1, -1}}

	adj := make([]point, len(offsets))
	for i, offset := range offsets {
		x, y := p.x+offset[0], p.y+offset[1]
		adj[i] = point{x, y}
	}

	return adj
}

func step(grid map[point]bool) map[point]bool {
	newGrid := make(map[point]bool)
	for key := range grid {
		neighbors := allAdjacent(key)
		neighbors = append(neighbors, key)
		for _, neigh := range neighbors {
			numBlack := numNeighbors(neigh, grid)
			isBlack := grid[neigh]
			if isBlack && (numBlack > 0 && numBlack <= 2) {
				newGrid[neigh] = true
			} else if !isBlack && numBlack == 2 {
				newGrid[neigh] = true
			}
		}
	}

	return newGrid
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tiles := make(map[point]bool)
	for scanner.Scan() {
		text := scanner.Text()
		input := text
		x, y := 0, 0
		for len(input) > 0 {
			if strings.HasPrefix(input, "e") {
				// east
				x += 2
				input = input[1:]
			} else if strings.HasPrefix(input, "se") {
				// south-east
				x++
				y++
				input = input[2:]
			} else if strings.HasPrefix(input, "sw") {
				// south-west
				x--
				y++
				input = input[2:]
			} else if strings.HasPrefix(input, "w") {
				// west
				x -= 2
				input = input[1:]
			} else if strings.HasPrefix(input, "nw") {
				// north-west
				x--
				y--
				input = input[2:]
			} else if strings.HasPrefix(input, "ne") {
				// north-east
				x++
				y--
				input = input[2:]
			}
		}

		newPoint := point{x, y}
		tiles[newPoint] = !tiles[newPoint]

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	numBlack := 0
	for _, val := range tiles {
		if val {
			numBlack++
		}
	}

	fmt.Println("Part 1:", numBlack)

	// Part 2

	for i := 0; i < 100; i++ {
		tiles = step(tiles)
	}

	numBlack = 0
	for _, val := range tiles {
		if val {
			numBlack++
		}
	}

	fmt.Println("Part 2:", numBlack)
}
