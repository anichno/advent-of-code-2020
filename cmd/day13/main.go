package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	earliestTime, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	busIds := make([]int, 0)
	busPos := make([]int, 0)
	for i, bus := range strings.Split(scanner.Text(), ",") {
		busVal, err := strconv.Atoi(bus)
		if err == nil {
			busIds = append(busIds, busVal)
			busPos = append(busPos, i)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minWait := 1<<(64-1) - 1
	var minBus int
	for _, bus := range busIds {
		div := earliestTime / bus
		if earliestTime%bus > 0 {
			div++
		}

		nextTime := div * bus
		wait := nextTime - earliestTime
		if wait < minWait {
			minWait = wait
			minBus = bus
		}
	}

	fmt.Println("Part 1:", minBus*minWait)

	leftBusID := busIds[0]
	leftBusPos := busPos[0]
	for i := 1; i < len(busIds); i++ {

		rightBusID := busIds[i]
		rightBusPos := busPos[i]

		newLeftID := 0
		newLeftPos := 0
		for timestep := leftBusID - leftBusPos; ; timestep += leftBusID {
			if (timestep+leftBusPos)%leftBusID == 0 && (timestep+rightBusPos)%rightBusID == 0 {
				newLeftID = leftBusID * rightBusID
				newLeftPos = timestep
				break
			}
		}
		leftBusID = newLeftID
		leftBusPos = -newLeftPos
	}

	fmt.Println("Part 2:", -leftBusPos)
}
