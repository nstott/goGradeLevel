package goGradeLevel

import (
	"unicode"
	"unicode/utf8"
)

func CountSentences(s string) int {
	total := 0
	runes := []rune(s)
	l := utf8.RuneCountInString(s)
	startWord := 0
	for k, v := range runes {
		if k == l-1 { // if we're at the end of the loop, then add 1, 
			total += 1
			continue
		} 
		if isPunct(v) && unicode.IsSpace(runes[k+1]) && isNextCharCapitalized(runes[k+1:]) && (k - startWord) > 2 {
			// if we're a correct punctuation rune
			// and the next character is a space
			// and the first char after the space(s) is a capital
			// and the start of the word is at least 2 runes away
			total += 1
		}
		if !unicode.IsLetter(v) {  
			// if we're at the beginning of a word, then update the index.
			startWord = k
		}

	}
	return total
}

func isPunct(r rune) bool {
	// this needs work,
	// http://www.fileformat.info/info/unicode/category/Po/list.htm
	// which chars, apart from a comma don't mean end of sentence.  quotes perhaps, semi colon
	return r != ',' && unicode.IsPunct(r)
}

func isNextCharCapitalized(runes []rune) bool {
	for _, v := range runes {
		if unicode.IsSpace(v) {
			continue
		}
		if unicode.ToUpper(v) == v {
			return true
		}
		return false
	}
	return false
}
