package main

import (
	"strings"
	"sync"
)

type LetterMap map[rune]int

type AsyncLetterMap struct {
	sync.Mutex
	internal LetterMap
}

var wg sync.WaitGroup

// Recieve a string and returns the frequency of each UNICODE character
func GetLettersFrequency(text string) LetterMap {
	wg.Add(1)
	
	var letters AsyncLetterMap
	letters.internal = make(LetterMap)
	text = strings.Replace(text, " ", "", -1) // trims whitespace

	// Go over the letters and add them to the map
	for _, char := range text {
		if _, isInMap := letters.internal[rune(char)]; !isInMap {
			wg.Add(1)
			go func() {
				char = rune(char)
				freq := CountLetterFrequency(text, char) 
				AddLetterToMap(char, freq, &letters, &wg)
			}()
		}
	}

	defer wg.Done();
	return letters.internal
}

// Count the frequency of a certain letter in a string
func CountLetterFrequency(text string, letter rune) int {
	frequency := 0
	
	for _, currLetter := range text {
		if currLetter == letter{
			frequency++
		}
	}
	
	return frequency
}

// Adds a letter to a map in a safe way
func AddLetterToMap(letter rune, amount int, letterMap *AsyncLetterMap, wg *sync.WaitGroup) {
	defer wg.Done();
	letterMap.Lock()
	letterMap.internal[letter] = amount
	letterMap.Unlock()
}