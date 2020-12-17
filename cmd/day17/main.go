package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/golang-collections/collections/set"
)

type coord struct {
	x int
	y int
	z int
	w int
}

func numActiveNeighbors(pos coord, space map[coord]bool) int {
	neighborsActive := 0
	for _, neighbor := range neighbors(pos) {
		if _, ok := space[neighbor]; ok {
			neighborsActive++
		}
	}

	return neighborsActive
}

func neighbors(pos coord) []coord {
	diffs := [3]int{-1, 0, 1}

	neighbors := make([]coord, 0)
	for _, xdiff := range diffs {
		for _, ydiff := range diffs {
			for _, zdiff := range diffs {
				if xdiff == 0 && ydiff == 0 && zdiff == 0 {
					continue
				}
				neighbors = append(neighbors, coord{pos.x + xdiff, pos.y + ydiff, pos.z + zdiff, 0})
			}
		}
	}

	return neighbors
}

// Part 2 versions ///////////

func numActiveNeighborsPart2(pos coord, space map[coord]bool) int {
	neighborsActive := 0
	for _, neighbor := range neighborsPart2(pos) {
		if _, ok := space[neighbor]; ok {
			neighborsActive++
		}
	}

	return neighborsActive
}

func neighborsPart2(pos coord) []coord {
	diffs := [3]int{-1, 0, 1}

	neighbors := make([]coord, 0)
	for _, xdiff := range diffs {
		for _, ydiff := range diffs {
			for _, zdiff := range diffs {
				for _, wdiff := range diffs {
					if xdiff == 0 && ydiff == 0 && zdiff == 0 && wdiff == 0 {
						continue
					}
					neighbors = append(neighbors, coord{pos.x + xdiff, pos.y + ydiff, pos.z + zdiff, pos.w + wdiff})
				}

			}
		}
	}

	return neighbors
}

//////////////////////////////

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	initialSpace := make(map[coord]bool)

	for y := 0; scanner.Scan(); y++ {
		text := scanner.Text()
		for x, char := range text {
			newCoord := coord{x, y, 0, 0}
			if char == '#' {
				initialSpace[newCoord] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	space := make(map[coord]bool)
	for key, val := range initialSpace {
		space[key] = val
	}

	for cycle := 0; cycle < 6; cycle++ {
		possibleCoords := set.New()
		for pos := range space {
			possibleCoords.Insert(pos)
			for _, neighbor := range neighbors(pos) {
				possibleCoords.Insert(neighbor)
			}
		}

		newSpace := make(map[coord]bool)
		possibleCoords.Do(func(pos interface{}) {
			cube := pos.(coord)
			activeNeighbors := numActiveNeighbors(cube, space)
			if _, ok := space[cube]; ok {
				if activeNeighbors == 2 || activeNeighbors == 3 {
					newSpace[cube] = true
				}
			} else if activeNeighbors == 3 {
				newSpace[cube] = true
			}
		})

		space = newSpace
	}

	fmt.Println("Part 1:", len(space))

	// Part 2

	space = make(map[coord]bool)
	for key, val := range initialSpace {
		space[key] = val
	}

	for cycle := 0; cycle < 6; cycle++ {
		possibleCoords := set.New()
		for pos := range space {
			possibleCoords.Insert(pos)
			for _, neighbor := range neighborsPart2(pos) {
				possibleCoords.Insert(neighbor)
			}
		}

		newSpace := make(map[coord]bool)
		possibleCoords.Do(func(pos interface{}) {
			cube := pos.(coord)
			activeNeighbors := numActiveNeighborsPart2(cube, space)
			if _, ok := space[cube]; ok {
				if activeNeighbors == 2 || activeNeighbors == 3 {
					newSpace[cube] = true
				}
			} else if activeNeighbors == 3 {
				newSpace[cube] = true
			}
		})

		space = newSpace
	}

	fmt.Println("Part 2:", len(space))
}
