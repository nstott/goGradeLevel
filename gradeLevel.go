package goGradeLevel

import (
	"strings"
	"unicode"
)

type Result struct {
	readingEase float64
	fleschKincaidGradeLevel float64
	colemanLiauIndex float64
}


func Analyze(s string) *Result {

	numWords, numSyllables, numSentences := phraseStats(s)
	numLetters := countLetters(s)

	return &Result{
		readingEase: calcFleschReadingEase(numWords, numSyllables, numSentences),
		fleschKincaidGradeLevel: calcFleschKincaidGradeLevel(numWords, numSyllables, numSentences),
		colemanLiauIndex: calcColemanLiauIndex(numWords, numLetters, numSentences),
	}
}


func FleschReadingEase(s string) float64 {
	numWords, numSyllables, numSentences := phraseStats(s)
	return calcFleschReadingEase(numWords, numSyllables, numSentences)

}

func calcFleschReadingEase(numWords, numSyllables, numSentences float64) float64 {
	return 206.835 - (1.015 * (numWords / numSentences)) - (84.6 * (numSyllables / numWords))
}


func FleschKincaidGradeLevel(s string) float64 {
	numWords, numSyllables, numSentences := phraseStats(s)
	return calcFleschKincaidGradeLevel(numWords, numSyllables, numSentences)
}

func calcFleschKincaidGradeLevel(numWords, numSyllables, numSentences float64) float64 {
	return (0.39*(numWords/numSentences) + 11.8*(numSyllables/numWords)) - 15.59
}

func ColemanLiauIndex(s string) float64 {
	numWords, _, numSentences := phraseStats(s)
	numLetters := countLetters(s)
	return calcColemanLiauIndex(numWords, numLetters, numSentences)}

func calcColemanLiauIndex(numWords, numLetters, numSentences float64) float64 { 
	let := 100 * (numLetters / numWords)
	sen := 100 * (numSentences / numWords)

	return (0.0588 * let) - (0.296 * sen) - 15.8
}

/*
 * utility functions to count things in the text, 
 * like num Sentences, letters, syllables, etc
 */
func countLetters(s string) float64 {
	numLetters := 0.0
	for _, v := range []rune(s) {
		if unicode.IsLetter(v) {
			numLetters += 1.0
		}
	}
	return numLetters	
}

func phraseStats(s string) (numWords, numSyllables, numSentences float64) {
	words := strings.Fields(s)
	numWords = float64(len(words))
	numSyllables = 0.0
	for _, w := range words {
		numSyllables += float64(CountSyllables(w))
	}

	numSentences = float64(CountSentences(s))

	return
}
