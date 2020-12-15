package main

import "fmt"

func main() {
	input := []int{13, 0, 10, 12, 1, 5, 8}

	inputPart1 := make([]int, len(input))
	copy(inputPart1, input)
	for len(inputPart1) < 2020 {
		lastNum := inputPart1[len(inputPart1)-1]

		foundPrev := false
		for i := len(inputPart1) - 2; i >= 0; i-- {
			if inputPart1[i] == lastNum {
				foundPrev = true
				inputPart1 = append(inputPart1, len(inputPart1)-1-i)
				break
			}
		}

		if !foundPrev {
			inputPart1 = append(inputPart1, 0)
		}
	}

	fmt.Println("Part 1:", inputPart1[len(inputPart1)-1])

	// Part 2

	recentIndex := make(map[int]int)
	for i, val := range input {
		recentIndex[val] = i
	}

	nextVal := 0
	for curIndex := len(input); curIndex < 30000000-1; curIndex++ {
		lastSpokenDiff := 0
		if lastIdx, ok := recentIndex[nextVal]; ok {
			lastSpokenDiff = curIndex - lastIdx
		}

		recentIndex[nextVal] = curIndex
		nextVal = lastSpokenDiff
	}

	fmt.Println("Part 2:", nextVal)
}
