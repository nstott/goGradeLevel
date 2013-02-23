package goGradeLevel

import (
	"fmt"
	"strings"
	"unicode"
)

func FleschReadingEase(s string) float64 {
	numWords, numSyllables, numSentences := phraseStats(s)
	score := 206.835 - (1.015 * (numWords / numSentences)) - (84.6 * (numSyllables / numWords))

	return score
}

func FleschKincaidGradeLevel(s string) float64 {
	numWords, numSyllables, numSentences := phraseStats(s)
	return (0.39*(numWords/numSentences) + 11.8*(numSyllables/numWords)) - 15.59
}

func ColemanLiauIndex(s string) float64 {
	numWords, _, numSentences := phraseStats(s)
	numLetters := 0.0
	for _, v := range []rune(s) {
		if unicode.IsLetter(v) {
			numLetters += 1.0
		}
	}

	let := 100 * (numLetters / numWords)
	sen := 100 * (numSentences / numWords)

	return (0.0588 * let) - (0.296 * sen) - 15.8
}

func phraseStats(s string) (numWords, numSyllables, numSentences float64) {
	words := strings.Fields(s)
	numWords = float64(len(words))
	numSyllables = 0.0
	for _, w := range words {
		numSyllables += float64(CountSyllables(w))
	}

	numSentences = float64(CountSentences(s))

	fmt.Printf("numWords = %f, numSyllables = %f, numSentences = %f\n", numWords, numSyllables, numSentences)

	return
}
