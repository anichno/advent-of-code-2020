package main

import "fmt"

const subjectNumber int = 7
const cryptoMod int = 20201227

func main() {
	cardPubKey := 9033205
	doorPubKey := 9281649

	var cardLoops int
	var doorLoops int

	for numLoops, cardVal, doorVal := 1, 1, 1; cardLoops == 0 || doorLoops == 0; numLoops++ {
		cardVal = (cardVal * subjectNumber) % cryptoMod
		doorVal = (doorVal * subjectNumber) % cryptoMod

		if cardVal == cardPubKey {
			cardLoops = numLoops
		}

		if doorVal == doorPubKey {
			doorLoops = numLoops
		}
	}

	encryptionKey := 1
	for i := 0; i < cardLoops; i++ {
		encryptionKey = (encryptionKey * doorPubKey) % cryptoMod
	}

	fmt.Println("Part 1:", encryptionKey)
}
