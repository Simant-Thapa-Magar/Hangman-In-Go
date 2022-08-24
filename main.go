package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var Scanner = bufio.NewReader(os.Stdin)

func main() {
	// starting point
	//hangmanState := 0
	selectedWord := initializeGame()
	fmt.Println("Selected word is ", selectedWord)
}

func initializeGame() WordDictonary {
	// initialize game i.e. select a word to start game
	fmt.Println("Welcome to Hangman")

	validSelection := false
	var possibleWordOptions []WordDictonary
	var selectedWord WordDictonary

	for !validSelection {
		fmt.Printf("Select difficulty mode \n 1 for easy mode \n 2 for medium mode \n 3 for hard mode\n")
		printReaderIcon()
		difficultyMode, err := Scanner.ReadString('\n')
		if err != nil {
			panic("Something went wrong !! Game terminated !!")
		}

		difficultyModeInt, err := strconv.Atoi(strings.TrimSpace(difficultyMode))

		if err == nil {

			switch difficultyModeInt {
			case 1:
				possibleWordOptions = easyWords
			case 2:
				possibleWordOptions = mediumWWords
			case 3:
				possibleWordOptions = difficultWords
			}

			if len(possibleWordOptions) > 0 {
				validSelection = true
				selectedWord = possibleWordOptions[generateANumber(len(possibleWordOptions))]
			}
		}
	}
	return selectedWord
}

func generateANumber(max int) int {
	// generates a random number between 0 to max
	// returns generated random number
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func displayHangman(hangmanState int) {
	// displays hangman based on state
}

func displayGame(targetWord string) {
	// displays word with guessed and blanks for game
}

func getUserInput() {
	// gets input from user for game
}

func validateUserInput() {
	// validates user input for game
}

func determineInputProgress(input string) {
	// determines if user input is correct or not or if user has asked for hint
}

func isWordGuessedSuccessfully() {
	// determines if user has guessed the word successfully
}

func isHangmanComplete(hangmanState int) {
	// determines if hangman has been drawn completely
}

func printReaderIcon() {
	// displays a simple greater than icon to indicate user input is expected
	fmt.Print("> ")
}
