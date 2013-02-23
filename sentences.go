package goGradeLevel

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func CountSentences(s string) int {
	f := strings.FieldsFunc(s, isTerminatingPunc)

	// only include sentences that are larger then 2 characters
	total := 0
	for _, v := range f {
		if len(v) > 2 {
			total += 1
		}
	}
	return total
}

func CountSentences2(s string) int {

	fmt.Printf("Counting %s\n", s)
	total := 0
	runes := []rune(s)
	l := utf8.RuneCountInString(s)
	for k, v := range runes {
		if k == l-1 {
			total += 1
			continue
		}
		if isPunct(v) {
			// punc
			// if we're at the end, then add one
			// if we're not at the end, but have a space or another punc then add one
			if isPunct(runes[k]) && unicode.IsSpace(runes[k+1]) && isNextCharCapitalized(runes[k+1:]) {
				total += 1
			}

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

func isTerminatingPunc(c rune) bool {
	return c == '.' || c == '!' || c == '?'
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
