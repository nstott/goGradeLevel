package goGradeLevel

import (
	"strings"
	"unicode/utf8"
)

var (
	vowels    = "aeiouy"
	dipthongs = []string{"ou", "oa", "ao", "oi", "io", "ui", "ea", "ei", "ie", "eo", "oe", "eu", "ue", "ai", "au"}
)

func isVowel(s string) bool {
	if utf8.RuneCountInString(s) < 1 {
		return false
	}
	var rune, _ = utf8.DecodeRuneInString(s)
	return strings.ContainsRune(vowels, rune)
}

func isDipthong(s string) bool {
	if utf8.RuneCountInString(s) < 2 {
		return false
	}

	firstTwo := string([]rune(s)[0:2])
	for _, d := range dipthongs {
		if d == firstTwo {
			return true
		}
	}
	return false
}

func isDouble(s string) bool {
	if utf8.RuneCountInString(s) < 2 {
		return false
	}
	var runes = []rune(s)
	if runes[0] != runes[1] {
		return false
	}

	return strings.ContainsRune(vowels, runes[0])
}

func skip(s string, i int) string {
	if utf8.RuneCountInString(s) < i {
		return ""
	}
	return string([]rune(s)[i:])
}

func endsWithSilentVowel(s string) bool {
	l := utf8.RuneCountInString(s)
	if l < 1 {
		return false
	}

	last, _ := utf8.DecodeLastRuneInString(s)
	if last != 'e' {
		return false
	}

	runes := []rune(s)
	if l > 1 && runes[l-2] == 'l' {
		return false
	}
	return true
}

func countIter(s string, count int) int {
	if s == "" {
		return count
	}
	switch {
	case isDouble(s):
		return countIter(skip(s, 2), count+1)
	case isDipthong(s):
		return countIter(skip(s, 2), count+1)
	case isVowel(s):
		return countIter(skip(s, 1), count+1)
	}
	return countIter(skip(s, 1), count)
}

func CountSyllables(s string) int {
	count := countIter(strings.ToLower(s), 0)
	if endsWithSilentVowel(s) {
		return count - 1
	}
	return count
}
