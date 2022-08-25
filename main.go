package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var Scanner = bufio.NewReader(os.Stdin)

func main() {
	// starting point
	hangmanState := 0
	selectedWord := initializeGame()
	var wordGuessTrack = make(map[rune]bool)
	gameEnd := false
	showHint := false
	success := false
	var oldMessage string
	for {
		displayGame(hangmanState, selectedWord, wordGuessTrack, showHint, oldMessage)
		if gameEnd {
			break
		}
		userInput := getUserInput()
		if userInput == "hint" {
			showHint = true
		} else {
			message, isCorrectGuess := determineInputProgress(userInput, selectedWord, wordGuessTrack)
			oldMessage = message
			if !isCorrectGuess {
				hangmanState++
			}
		}
		if isHangmanComplete(hangmanState) {
			gameEnd = true
		} else if isWordGuessedSuccessfully(selectedWord, wordGuessTrack) {
			gameEnd = true
			success = true
		}
	}

	if success {
		fmt.Println("Congratulations !!!! You won")
	} else {
		fmt.Println("Sorry you're dead !!!!")
	}
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
	content, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))

	if err != nil {
		panic("Something went wrong while drawing hangman !!!!")
	}

	fmt.Println(string(content))
}

func displayWords(word WordDictonary, guessTrack map[rune]bool) {
	// displays word with guessed and blanks for game
	targetWord := word.Word
	for index, letter := range targetWord {
		if isInitiallyDisplayed(word.InitialDisplayPositions, index) || guessTrack[unicode.ToLower(rune(letter))] {
			fmt.Print(string(letter))
		} else {
			fmt.Print("_")
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

func getUserInput() string {
	// gets input from user for game
	var trimmedInput string
	isValidInput := false
	for !isValidInput {
		fmt.Println("Start Guessing (a letter) ")
		printReaderIcon()
		input, err := Scanner.ReadString('\n')

		if err != nil {
			panic("Error while reading input")
		}

		trimmedInput = strings.TrimSpace(input)
		trimmedInput = strings.ToLower(trimmedInput)
		isValidInput = validateUserInput(trimmedInput)
	}

	return trimmedInput
}

func validateUserInput(input string) bool {
	// validates user input for game
	if input == "hint" || len(input) == 1 {
		return true
	}
	return false
}

func determineInputProgress(input string, selectedWord WordDictonary, wordGuessTrack map[rune]bool) (string, bool) {
	// determines if user input is correct or not
	var message string
	var isCorrectGuess bool
	if wordGuessTrack[rune(input[0])] {
		message = fmt.Sprintf("You have already guessed %s\n", string(input[0]))
		isCorrectGuess = true
	} else {
		for _, letter := range selectedWord.Word {
			if strings.ToLower(string(letter)) == input {
				isCorrectGuess = true
				wordGuessTrack[rune(strings.ToLower(input)[0])] = true
			}
		}
	}
	return message, isCorrectGuess
}

func isWordGuessedSuccessfully(selectedWord WordDictonary, wordGuessTrack map[rune]bool) bool {
	// determines if user has guessed the word successfully
	for index, letter := range selectedWord.Word {
		lowerRune := rune(strings.ToLower(string(letter))[0])
		if !wordGuessTrack[lowerRune] && !isInitiallyDisplayed(selectedWord.InitialDisplayPositions, index) {
			return false
		}
	}
	return true
}

func isHangmanComplete(hangmanState int) bool {
	// determines if hangman has been drawn completely
	return hangmanState >= 9
}

func printReaderIcon() {
	// displays a simple greater than icon to indicate user input is expected
	fmt.Print("> ")
}

func isInitiallyDisplayed(arr []int, elem int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

func displayGame(hangmanState int, selectedWord WordDictonary, wordGuessTrack map[rune]bool, displayHint bool, oldMessage string) {
	clearScreen()
	displayHangman(hangmanState)
	displayWords(selectedWord, wordGuessTrack)
	if displayHint {
		fmt.Println("HINT : ", selectedWord.Hint)
	} else {
		fmt.Println("Type hint to see word hint")
	}

	if len(oldMessage) > 0 {
		fmt.Println(oldMessage)
	}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
