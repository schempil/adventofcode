package day22

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve() {

	absolutePath, _ := filepath.Abs("./day22/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	fmt.Println("Solution Day 22 - Part 1:", getWinningPlayersScore(inputs))
	fmt.Println("Solution Day 22 - Part 2:")
}

type player struct {
	cards []int
}

func getWinningPlayersScore(inputs []string) int {

	playerOne, playerTwo := getPlayersFromInput(inputs)
	playerOne, playerTwo = playGame(playerOne, playerTwo)

	var winningPlayer *player

	if len(playerOne.cards) > 0 {
		winningPlayer = &playerOne
	}

	if len(playerTwo.cards) > 0 {
		winningPlayer = &playerTwo
	}

	score := calculateScoreFromCards(winningPlayer.cards)

	return score
}

func calculateScoreFromCards(cards []int) int {

	score := 0

	for i := len(cards) - 1; i >= 0; i-- {
		score += cards[i] * (len(cards) - i)

	}

	return score
}

func playGame(playerOne player, playerTwo player) (playerOneNew player, playerTwoNew player) {

	if len(playerOne.cards) == 0 || len(playerTwo.cards) == 0 {
		return playerOne, playerTwo
	}

	currentCardPlayerOne := playerOne.cards[0]
	currentCardPlayerTwo := playerTwo.cards[0]

	playerOneNew.cards = playerOne.cards[1:]
	playerTwoNew.cards = playerTwo.cards[1:]

	if currentCardPlayerOne > currentCardPlayerTwo {
		playerOneNew.cards = append(playerOneNew.cards, currentCardPlayerOne)
		playerOneNew.cards = append(playerOneNew.cards, currentCardPlayerTwo)
	}

	if currentCardPlayerOne < currentCardPlayerTwo {
		playerTwoNew.cards = append(playerTwoNew.cards, currentCardPlayerTwo)
		playerTwoNew.cards = append(playerTwoNew.cards, currentCardPlayerOne)
	}

	return playGame(playerOneNew, playerTwoNew)
}

func getPlayersFromInput(inputs []string) (playerOne player, playerTwo player) {

	playerOne = player{}
	playerTwo = player{}
	var activePlayer *player

	for _, input := range inputs {

		if len(input) == 0 {
			continue
		}

		if input == "Player 1:" {
			activePlayer = &playerOne
			continue
		}
		if input == "Player 2:" {
			activePlayer = &playerTwo
			continue
		}

		card, _ := strconv.Atoi(input)
		activePlayer.cards = append(activePlayer.cards, card)
	}

	return playerOne, playerTwo
}
