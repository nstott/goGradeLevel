package goGradeLevel

import (
	"fmt"
	"testing"
)

func Test_CountSentences(t *testing.T) {
	var d = []struct {
		in   string
		want int
	}{
		{"This. Is. Three.", 3},
		{"Que Paso?", 1},
		{"fancy pants", 1},
		{"I LOVE PIE! DON'T YOU LOVE PIE?", 2},
	}

	for _, v := range d {
		if CountSentences(v.in) != v.want {
			t.Error(fmt.Sprintf("CountSentences(%s) != %d, got: %d", v.in, v.want, CountSentences(v.in)))
		}
	}
}

func Test_CountSentences2(t *testing.T) {
	var d = []struct {
		in   string
		want int
	}{
		{"This. Is. Three.", 3},
		{"Que Paso?", 1},
		{"fancy pants", 1},
		{"I LOVE PIE! DON'T YOU LOVE PIE?", 2},
		{"That’s why it is also called Flesch-Kincaid Grade Level Readability Test.", 1},
		{"That’s why it is also called Flesch-Kincaid " +
			"Grade Level Readability Test. Raised in Austria, Flesch studied law and earned a Ph.D. " +
			"in English from the Columbia University.", 2},
	}

	for _, v := range d {
		c := CountSentences2(v.in)
		if c != v.want {
			t.Error(fmt.Sprintf("CountSentences(%s) != %d, got: %d", v.in, v.want, c))
		}
	}
}
