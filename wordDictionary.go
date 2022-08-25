package main

type WordDictonary struct {
	Word                    string
	InitialDisplayPositions []int
	Hint                    string
}

var easyWords = []WordDictonary{
	{
		Word:                    "Google",
		InitialDisplayPositions: []int{4},
		Hint:                    "Developers of Go Lang",
	},
	{
		Word:                    "Nepal",
		InitialDisplayPositions: []int{2},
		Hint:                    "Country of Himalayas",
	},
	{
		Word:                    "Computer",
		InitialDisplayPositions: []int{2, 6},
		Hint:                    "Device that you are working on",
	},
}

var mediumWWords = []WordDictonary{
	{
		Word:                    "Beekeeper",
		InitialDisplayPositions: []int{2, 5},
		Hint:                    "Thank them for honey",
	},
	{
		Word:                    "Strength",
		InitialDisplayPositions: []int{2, 7},
		Hint:                    "Power",
	},
	{
		Word:                    "Nightclub",
		InitialDisplayPositions: []int{1, 6},
		Hint:                    "Party after sun sets down",
	},
}

var difficultWords = []WordDictonary{
	{
		Word:                    "Squush",
		InitialDisplayPositions: []int{2},
		Hint:                    "Crushing",
	},
	{
		Word:                    "Kickshaw",
		InitialDisplayPositions: []int{1},
		Hint:                    "A fancy dish",
	},
	{
		Word:                    "Zugzwang",
		InitialDisplayPositions: []int{2},
		Hint:                    "Pull",
	},
	{
		Word:                    "Mississippi",
		InitialDisplayPositions: []int{1, 3, 8},
		Hint:                    "A state in the Southeastern region of the United States",
	},
}
