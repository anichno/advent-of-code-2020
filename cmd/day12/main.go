package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	commands := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		commands = append(commands, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	positionX, positionY := 0, 0
	direction := 0
	for _, text := range commands {
		action := text[0]
		val, err := strconv.Atoi(text[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case 'N':
			positionY += val
		case 'S':
			positionY -= val
		case 'E':
			positionX += val
		case 'W':
			positionX -= val
		case 'L':
			direction = (direction - val) % 360
			if direction < 0 {
				direction += 360
			}
		case 'R':
			direction = (direction + val) % 360
		case 'F':
			switch direction {
			case 0:
				positionX += val
			case 90:
				positionY -= val
			case 180:
				positionX -= val
			case 270:
				positionY += val
			default:
				log.Fatal(direction)
			}
		default:
			log.Fatal(action)
		}
	}

	distance := math.Abs(float64(positionX)) + math.Abs(float64(positionY))
	fmt.Println("Part 1:", distance)

	// Part 2
	positionX, positionY = 0, 0
	waypointX, waypointY := 10, 1
	for _, text := range commands {
		action := text[0]
		val, _ := strconv.Atoi(text[1:])

		switch action {
		case 'N':
			waypointY += val
		case 'S':
			waypointY -= val
		case 'E':
			waypointX += val
		case 'W':
			waypointX -= val
		case 'L':
			for i := 0; i < val/90; i++ {
				waypointX, waypointY = -waypointY, waypointX
			}
		case 'R':
			for i := 0; i < val/90; i++ {
				waypointX, waypointY = waypointY, -waypointX
			}
		case 'F':
			for i := 0; i < val; i++ {
				positionX += waypointX
				positionY += waypointY
			}
		}
	}

	distance = math.Abs(float64(positionX)) + math.Abs(float64(positionY))
	fmt.Println("Part 2:", distance)
}
