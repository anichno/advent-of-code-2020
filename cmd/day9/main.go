package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/ericaro/ringbuffer"
)

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufSize := 25

	scanner := bufio.NewScanner(file)
	buffer := ringbuffer.New(bufSize)
	fullList := make([]int, 0)
	var part1Num int
	var part1Idx int
	for i := 0; ; scanner.Scan() {
		text := scanner.Text()
		curNum, _ := strconv.Atoi(text)
		fullList = append(fullList, curNum)
		if i > bufSize {
			validNum := false
			for j := 0; j < bufSize; j++ {
				for k := 0; k < bufSize; k++ {
					if j == k {
						continue
					}
					val1iface, _ := buffer.Get(j)
					val1, _ := val1iface.(int)
					val2iface, _ := buffer.Get(k)
					val2, _ := val2iface.(int)
					if val1+val2 == curNum {
						validNum = true
						break
					}
				}
				if validNum {
					break
				}
			}

			if !validNum {
				fmt.Println("Part 1:", curNum)
				part1Num = curNum
				part1Idx = i
				break
			}
		}
		if i < bufSize {
			buffer.Add(curNum)
		} else {
			buffer.Push(curNum)
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

outer:
	for startIdx := 0; startIdx < part1Idx; startIdx++ {
		for windowLen := 2; windowLen < part1Idx-startIdx; windowLen++ {
			sum := 0
			min := int(math.Pow(2, 32) - 1)
			max := -1
			for i := startIdx; i < windowLen+startIdx; i++ {
				sum += fullList[i]
				if fullList[i] < min {
					min = fullList[i]
				} else if fullList[i] > max {
					max = fullList[i]
				}
				if sum == part1Num {
					fmt.Println("Part 2:", min+max)
					break outer
				}
			}
		}
	}

}
