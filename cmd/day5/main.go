package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getRec(leftChar byte, input []byte, left int, right int) int {
	var middle int
	middle = (left + right) / 2

	if left == right {
		return left
	}

	if input[0] == leftChar {
		return getRec(leftChar, input[1:], left, middle)
	}
	return getRec(leftChar, input[1:], middle+1, right)

}

func getRow(input string) int {
	return getRec('F', []byte(input)[:7], 0, 127)
}

func getCol(input string) int {
	return getRec('L', []byte(input)[7:], 0, 7)
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxSeatID := 0
	seatIDs := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		seatID := getRow(text)*8 + getCol(text)
		seatIDs = append(seatIDs, seatID)

		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", maxSeatID)

	sort.Ints(seatIDs)
	for i, id := range seatIDs {
		if i == 0 {
			continue
		}
		if id-seatIDs[i-1] == 2 {
			fmt.Println("Part 2:", id-1)
		}
	}
}
