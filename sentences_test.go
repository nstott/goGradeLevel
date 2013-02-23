package goGradeLevel

import (
	"fmt"
	"testing"
)

type textInt struct {
	in string
	want int
}

func Test_CountSentences(t *testing.T) {
	var d = []textInt{
		{"This. Is. Three.", 3},
		{"Que Paso?", 1},
		{"fancy pants", 1},
		{"I LOVE PIE! DON'T YOU LOVE PIE?", 2},
		{"That’s why it is also called Flesch-Kincaid Grade Level Readability Test.", 1},
		{"That’s why it is also called Flesch-Kincaid " +
			"Grade Level Readability Test. Raised in Austria, Flesch studied law and earned a Ph.D. " +
			"in English from the Columbia University.", 2},
		{"Existing computer programs that measure readability are based largely upon subroutines which estimate " + 
			"number of syllables, usually by counting vowels. The shortcoming in estimating syllables is that " +
			"it necessitates keypunching the prose into the computer. There is no need to estimate syllables " +
			"since word length in letters is a better predictor of readability than word length in syllables. " +
			"Therefore, a new readability formula was computed that has for its predictors letters per 100 words " +
			"and sentences per 100 words. Both predictors can be counted by an optical scanning device, and thus " +
			"the formula makes it economically feasible for an organization such as the U.S. Office of Education to " +
			"calibrate the readability of all textbooks for the public school system.", 5},
		{"Flesch Grade Level Readability Formula improves upon the Flesch Reading Ease Readability Formula. " +
			"Rudolph Flesch, an author, writing consultant, and the supporter of Plain English Movement, " +
			"is the co-author of this formula along with John P. Kincaid. That’s why it is also called Flesch-Kincaid " +
			"Grade Level Readability Test. Raised in Austria, Flesch studied law and earned a Ph.D. " +
			"in English from the Columbia University. Flesch, through his writings and speeches, advocated a return to phonics. " +
			"In his article, A New Readability Yardstick, published in the Journal of Applied Psychology in 1948, Flesch proposed " +
			"the Reading Ease Readability Formula. ", 6},
	}

	for _, v := range d {
		c := CountSentences(v.in)
		if c != v.want {
			t.Error(fmt.Sprintf("CountSentences(%s) != %d, got: %d", v.in, v.want, c))
		}
	}
}
