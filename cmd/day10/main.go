package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/golang-collections/collections/set"
)

func solve1(chain []int, used *set.Set, adaptors []int) bool {
	if used.Len() == len(adaptors) {
		chain = append(chain, chain[len(chain)-1]+3)
		num1diff := 0
		num3diff := 0
		prevVal := chain[0]
		for _, val := range chain[1:] {
			if val-prevVal == 1 {
				num1diff++
			} else if val-prevVal == 3 {
				num3diff++
			}
			prevVal = val
		}
		fmt.Println("Part 1:", num1diff*num3diff)
		return true
	}

	lastElem := chain[len(chain)-1]
	for _, adapt := range adaptors {
		if adapt > lastElem && adapt-lastElem <= 3 && !used.Has(adapt) {
			used.Insert(adapt)
			chain = append(chain, adapt)
			if solve1(chain, used, adaptors) {
				return true
			}
			chain = chain[:len(chain)-1]
			used.Remove(adapt)
		}
	}

	return false
}

func solve2(curVal int, target int, adaptors []int, numFollowing map[int]int) int {
	if target-curVal == 3 {
		return 1
	}

	tot := 0
	for i, adapt := range adaptors {
		if adapt-curVal <= 3 {
			if val, ok := numFollowing[adapt]; ok {
				tot += val
			} else {
				totFollowing := solve2(adapt, target, adaptors[i+1:], numFollowing)
				numFollowing[adapt] = totFollowing
				tot += totFollowing
			}

		} else {
			break
		}
	}

	return tot
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	adaptors := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if val, err := strconv.Atoi(text); err == nil {
			adaptors = append(adaptors, val)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(adaptors)

	chain := []int{0}
	used := set.New()
	solve1(chain, used, adaptors)

	target := adaptors[len(adaptors)-1] + 3
	numFollowing := make(map[int]int)
	fmt.Println("Part 2:", solve2(0, target, adaptors, numFollowing))
}
