package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-collections/collections/set"
	"github.com/mitchellh/hashstructure"
)

type GameState struct {
	Player1 []int
	Player2 []int
}

type player int

const (
	player1ID = iota
	player2ID
)

func play(player1Cards, player2Cards []int) (player, []int) {
	prevStates := set.New()
	for len(player1Cards) > 0 && len(player2Cards) > 0 {
		state, err := hashstructure.Hash(GameState{player1Cards, player2Cards}, hashstructure.FormatV2, nil)
		if err != nil {
			log.Fatal(err)
		}

		if prevStates.Has(state) {
			return player1ID, player1Cards
		}
		prevStates.Insert(state)

		player1 := player1Cards[0]
		player2 := player2Cards[0]

		if len(player1Cards)-1 >= player1 && len(player2Cards)-1 >= player2 {
			// recurse
			player1NewDeck := make([]int, player1)
			copy(player1NewDeck, player1Cards[1:player1+1])
			player2NewDeck := make([]int, player2)
			copy(player2NewDeck, player2Cards[1:player2+1])

			winnerID, _ := play(player1NewDeck, player2NewDeck)
			if winnerID == player1ID {
				player1Cards = append(player1Cards[1:], player1, player2)
				player2Cards = player2Cards[1:]
			} else {
				player2Cards = append(player2Cards[1:], player2, player1)
				player1Cards = player1Cards[1:]
			}
		} else {
			if player1 > player2 {
				player1Cards = append(player1Cards[1:], player1, player2)
				player2Cards = player2Cards[1:]
			} else {
				player2Cards = append(player2Cards[1:], player2, player1)
				player1Cards = player1Cards[1:]
			}

		}
	}

	winnerDeck := append(player1Cards, player2Cards...)
	var winnerID player
	if len(player1Cards) > 0 {
		winnerID = player1ID
	} else {
		winnerID = player2ID
	}

	return winnerID, winnerDeck
}

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var player1CardsOrig []int
	var player2CardsOrig []int
	scanner := bufio.NewScanner(file)

	// Player 1
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		val, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}

		player1CardsOrig = append(player1CardsOrig, val)
	}

	// Player 2
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		val, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}

		player2CardsOrig = append(player2CardsOrig, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	player1Cards := make([]int, len(player1CardsOrig))
	copy(player1Cards, player1CardsOrig)
	player2Cards := make([]int, len(player2CardsOrig))
	copy(player2Cards, player2CardsOrig)

	for len(player1Cards) > 0 && len(player2Cards) > 0 {
		player1 := player1Cards[0]
		player2 := player2Cards[0]

		if player1 > player2 {
			player1Cards = append(player1Cards[1:], player1, player2)
			player2Cards = player2Cards[1:]
		} else {
			player2Cards = append(player2Cards[1:], player2, player1)
			player1Cards = player1Cards[1:]
		}
	}

	winnerDeck := append(player1Cards, player2Cards...)
	score := 0
	for i, j := len(winnerDeck), 0; i > 0; i-- {
		score += winnerDeck[j] * i
		j++
	}

	fmt.Println("Part 1:", score)

	// Part 2
	player1Cards = make([]int, len(player1CardsOrig))
	copy(player1Cards, player1CardsOrig)
	player2Cards = make([]int, len(player2CardsOrig))
	copy(player2Cards, player2CardsOrig)

	_, winnerDeck = play(player1Cards, player2Cards)

	score = 0
	for i, j := len(winnerDeck), 0; i > 0; i-- {
		score += winnerDeck[j] * i
		j++
	}

	fmt.Println("Part 2:", score)
}
