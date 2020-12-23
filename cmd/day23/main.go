package main

import (
	"fmt"
	"strconv"
	"strings"
)

type cup struct {
	value int
	next  *cup
	prev  *cup
}

func move(head *cup) *cup {
	// remove 3 cups
	removedCupsHead := head.next
	removedCupsTail := removedCupsHead.next.next

	// fixup cups
	head.next = removedCupsTail.next
	removedCupsTail.next.prev = head

	// destination cup
	var destinationCup *cup
	for i, curCup := head.value-1, head.next; ; {
		if curCup == head {
			i--
		}
		if i == 0 {
			i = 9
		}

		if curCup.value == i {
			destinationCup = curCup
			break
		}
		curCup = curCup.next
	}

	// place cups
	removedCupsTail.next = destinationCup.next
	destinationCup.next.prev = removedCupsTail
	destinationCup.next = removedCupsHead
	removedCupsHead.prev = destinationCup

	return head.next
}

func move2(head *cup, cupMap map[int]*cup, maxVal int) *cup {
	// remove 3 cups
	removedCupsHead := head.next
	removedCupsTail := removedCupsHead.next.next

	// fixup cups
	head.next = removedCupsTail.next
	removedCupsTail.next.prev = head

	// destination cup
	dCupVal := head.value - 1

	for {
		if dCupVal == 0 {
			dCupVal = maxVal
		}
		if dCupVal == removedCupsHead.value || dCupVal == removedCupsHead.next.value || dCupVal == removedCupsHead.next.next.value {
			dCupVal--
		} else {
			break
		}
	}

	destinationCup := cupMap[dCupVal]

	// place cups
	removedCupsTail.next = destinationCup.next
	destinationCup.next.prev = removedCupsTail
	destinationCup.next = removedCupsHead
	removedCupsHead.prev = destinationCup

	return head.next
}

func main() {
	input := []int{8, 7, 1, 3, 6, 9, 4, 5, 2}

	head := &cup{input[0], nil, nil}
	head.next = head
	head.prev = head
	curCup := head
	for _, val := range input[1:] {
		newCup := &cup{val, head, curCup}
		curCup.next = newCup
		head.prev = newCup
		curCup = newCup
	}

	for i := 0; i < 100; i++ {
		head = move(head)
	}

	var oneCup *cup
	curCup = head
	for {
		if curCup.value == 1 {
			oneCup = curCup
			break
		}
		curCup = curCup.next
	}

	var part1Answer strings.Builder
	curCup = oneCup.next
	for curCup != oneCup {
		part1Answer.WriteString(strconv.Itoa(curCup.value))
		curCup = curCup.next
	}

	fmt.Println("Part 1:", part1Answer.String())

	// Part 2
	maxVal := 0
	cupMap := make(map[int]*cup)
	head = &cup{input[0], nil, nil}
	cupMap[input[0]] = head
	head.next = head
	head.prev = head
	curCup = head
	for _, val := range input[1:] {
		if val > maxVal {
			maxVal = val
		}
		newCup := &cup{val, head, curCup}
		cupMap[val] = newCup
		curCup.next = newCup
		head.prev = newCup
		curCup = newCup
	}

	for i := 10; i <= 1000000; i++ {
		if i > maxVal {
			maxVal = i
		}
		newCup := &cup{i, head, curCup}
		cupMap[i] = newCup
		curCup.next = newCup
		head.prev = newCup
		curCup = newCup
	}

	for i := 0; i < 10000000; i++ {
		head = move2(head, cupMap, maxVal)
	}

	oneCup = cupMap[1]
	part2Answer := oneCup.next.value * oneCup.next.next.value

	fmt.Println("Part 2:", part2Answer)
}
