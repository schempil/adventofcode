package day22

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func Solve() {

	absolutePath, _ := filepath.Abs("./day22/input.txt")
	content, _ := ioutil.ReadFile(absolutePath)
	text := string(content)
	inputs := strings.Split(text, "\n")

	//fmt.Println("Solution Day 22 - Part 1:", getWinningPlayersScore(inputs, playNormalGame))
	fmt.Println("Solution Day 22 - Part 2:", getWinningPlayersScore(inputs, playRecursiveCombat))
}

type player struct {
	cards []int
}

type round struct {
	playerOne player
	playerTwo player
}

type playFunction func(playerOne player, playerTwo player, history []round) (playerOneNew player, playerTwoNew player, earlyFinish bool)

func getWinningPlayersScore(inputs []string, play playFunction) int {

	playerOne, playerTwo := getPlayersFromInput(inputs)

	playerOne, playerTwo, _ = play(playerOne, playerTwo, []round{})

	var winningPlayer *player

	if len(playerTwo.cards) > 0 {
		winningPlayer = &playerTwo
	}

	if len(playerOne.cards) > 0 {
		winningPlayer = &playerOne
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

func didAlreadyOccurInHistory(playerOne player, playerTwo player, history []round) bool {

	for _, entry := range history {
		if reflect.DeepEqual(entry.playerOne.cards, playerOne.cards) &&
			reflect.DeepEqual(entry.playerTwo.cards, playerTwo.cards) {
			return true
		}
	}

	return false
}

func playRecursiveCombat(playerOne player, playerTwo player, history []round) (playerOneNew player, playerTwoNew player, earlyFinish bool) {

	if didAlreadyOccurInHistory(playerOne, playerTwo, history) {
		return playerOne, playerTwo, true
	}

	history = append(history, round{
		playerOne: playerOne,
		playerTwo: playerTwo,
	})

	if len(playerOne.cards) == 0 || len(playerTwo.cards) == 0 {
		return playerOne, playerTwo, false
	}

	currentCardPlayerOne := playerOne.cards[0]
	currentCardPlayerTwo := playerTwo.cards[0]

	playerOneNew.cards = playerOne.cards[1:]
	playerTwoNew.cards = playerTwo.cards[1:]

	if currentCardPlayerOne <= len(playerOneNew.cards) && currentCardPlayerTwo <= len(playerTwoNew.cards) {

		subPlayerOne, subPlayerTwo, subEarlyFinish := playSubGame(
			append([]int{}, playerOneNew.cards[0:currentCardPlayerOne]...),
			append([]int{}, playerTwoNew.cards[0:currentCardPlayerTwo]...))

		if subEarlyFinish {
			playerOneNew.cards = append(playerOneNew.cards, currentCardPlayerOne)
			playerOneNew.cards = append(playerOneNew.cards, currentCardPlayerTwo)
			return playRecursiveCombat(playerOneNew, playerTwoNew, history)
		}

		if len(subPlayerOne.cards) > 0 {
			playerOneNew.cards = append(playerOneNew.cards, currentCardPlayerOne)
			playerOneNew.cards = append(playerOneNew.cards, currentCardPlayerTwo)
		}

		if len(subPlayerTwo.cards) > 0 {
			playerTwoNew.cards = append(playerTwoNew.cards, currentCardPlayerTwo)
			playerTwoNew.cards = append(playerTwoNew.cards, currentCardPlayerOne)
		}

		return playRecursiveCombat(playerOneNew, playerTwoNew, history)
	}

	if currentCardPlayerOne > currentCardPlayerTwo {
		playerOneNew.cards = append(playerOneNew.cards, currentCardPlayerOne)
		playerOneNew.cards = append(playerOneNew.cards, currentCardPlayerTwo)
	}

	if currentCardPlayerOne < currentCardPlayerTwo {
		playerTwoNew.cards = append(playerTwoNew.cards, currentCardPlayerTwo)
		playerTwoNew.cards = append(playerTwoNew.cards, currentCardPlayerOne)
	}

	return playRecursiveCombat(playerOneNew, playerTwoNew, history)
}

func playSubGame(playerOneCards []int, playerTwoCards []int) (subPlayerOne player, subPlayerTwo player, earlyFinish bool) {
	subPlayerOne = player{cards: playerOneCards}
	subPlayerTwo = player{cards: playerTwoCards}

	return playRecursiveCombat(subPlayerOne, subPlayerTwo, []round{})
}

func playNormalGame(playerOne player, playerTwo player, history []round) (playerOneNew player, playerTwoNew player, earlyFinish bool) {

	if len(playerOne.cards) == 0 || len(playerTwo.cards) == 0 {
		return playerOne, playerTwo, false
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

	return playNormalGame(playerOneNew, playerTwoNew, history)
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
