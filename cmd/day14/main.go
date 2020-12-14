package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func resoveAndWrite(memory map[int64]int64, addr [36]byte, val int64) {
	noX := true
	for _, c := range addr {
		if c == 'X' {
			noX = false
			break
		}
	}

	if noX {
		var realAddr int64 = 0
		for i, c := range addr {
			var bitval int64
			if c == '0' {
				bitval = 0
			} else {
				bitval = 1
			}
			realAddr |= bitval << (35 - i)

		}
		memory[realAddr] = val

		return
	}

	for i, c := range addr {
		if c == 'X' {
			// set 0
			addr[i] = '0'
			resoveAndWrite(memory, addr, val)

			// set 1
			addr[i] = '1'
			resoveAndWrite(memory, addr, val)

			break
		}
	}
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	var zeroMask int64
	var oneMask int64
	memRe, err := regexp.Compile(`^mem\[([0-9]+)\] = ([0-9]+)$`)
	memory := make(map[int64]int64)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	for _, text := range input {
		if strings.HasPrefix(text, "mask") {
			zeroMask = 0
			oneMask = 0
			mask := strings.TrimPrefix(text, "mask = ")
			for i, c := range mask {
				if c == '0' {
					zeroMask |= 1 << (35 - i)
				} else if c == '1' {
					oneMask |= 1 << (35 - i)
				}
			}
			zeroMask ^= 0xfffffffff
		} else {
			matches := memRe.FindStringSubmatch(text)
			addr, _ := strconv.ParseInt(matches[1], 10, 64)
			val, _ := strconv.ParseInt(matches[2], 10, 64)
			val &= zeroMask
			val |= oneMask
			memory[addr] = val
		}
	}

	var sum int64 = 0
	for _, val := range memory {
		sum += val
	}

	fmt.Println("Part 1:", sum)

	// Part 2
	var mask []byte
	memory = make(map[int64]int64)
	for _, text := range input {
		if strings.HasPrefix(text, "mask") {
			mask = []byte(strings.TrimPrefix(text, "mask = "))
		} else {
			matches := memRe.FindStringSubmatch(text)
			addr, _ := strconv.ParseInt(matches[1], 10, 64)
			val, _ := strconv.ParseInt(matches[2], 10, 64)

			var maskedAddr [36]byte
			for i := 0; i < 36; i++ {
				if mask[i] == '1' {
					maskedAddr[i] = '1'
				} else if mask[i] == 'X' {
					maskedAddr[i] = 'X'
				} else {
					if addr>>(35-i)&0x01 == 0 {
						maskedAddr[i] = '0'
					} else {
						maskedAddr[i] = '1'
					}
				}

			}
			resoveAndWrite(memory, maskedAddr, val)
		}
	}

	sum = 0
	for _, val := range memory {
		sum += val
	}

	fmt.Println("Part 2:", sum)
}
