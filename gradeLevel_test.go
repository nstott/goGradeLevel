package goGradeLevel

import (
	"fmt"
	"testing"
)

var floatErr = 0.0001

type textFloat struct {
	in string
	want float64
}

var text = []string{
	"Existing computer programs that measure readability are based largely upon subroutines which estimate " + 
		"number of syllables, usually by counting vowels. The shortcoming in estimating syllables is that " +
		"it necessitates keypunching the prose into the computer. There is no need to estimate syllables " +
		"since word length in letters is a better predictor of readability than word length in syllables. " +
		"Therefore, a new readability formula was computed that has for its predictors letters per 100 words " +
		"and sentences per 100 words. Both predictors can be counted by an optical scanning device, and thus " +
		"the formula makes it economically feasible for an organization such as the U.S. Office of Education to " +
		"calibrate the readability of all textbooks for the public school system.",
	"Flesch Grade Level Readability Formula improves upon the Flesch Reading Ease Readability Formula. " +
		"Rudolph Flesch, an author, writing consultant, and the supporter of Plain English Movement, " +
		"is the co-author of this formula along with John P. Kincaid. That’s why it is also called Flesch-Kincaid " +
		"Grade Level Readability Test. Raised in Austria, Flesch studied law and earned a Ph.D. " +
		"in English from the Columbia University. Flesch, through his writings and speeches, advocated a return to phonics. " +
		"In his article, A New Readability Yardstick, published in the Journal of Applied Psychology in 1948, Flesch proposed " +
		"the Reading Ease Readability Formula. ",
}


func Test_FleschReadingEase(t *testing.T) {
	var d = []textFloat{
		{text[0], 27.696487},
		{text[1], 39.185238},
	}

	for _, v := range d {
		r := FleschReadingEase(v.in)
		if !compareFloats(r, v.want) {
			t.Error(fmt.Sprintf("FleschReadingEase(%s) != %f, got %f", v.in, v.want, r))
		}
	}
}

func Test_FleschKincaidGradeLevel(t *testing.T) {
	var d = []textFloat{
		{text[0], 15.308807},
		{text[1], 11.851429},
	}

	for _, v := range d {
		r := FleschKincaidGradeLevel(v.in)
		if !compareFloats(r, v.want) {
			t.Error(fmt.Sprintf("FleschKincaidGradeLevel(%s) != %f, got %f", v.in, v.want, r))
		}
	}
}

func Test_ColemanLiauIndex(t *testing.T) {
	var d = []textFloat{
		{text[0], 14.233950},
		{text[1], 14.067755},
	}

	for _, v := range d {
		r := ColemanLiauIndex(v.in)
		if !compareFloats(r, v.want) {
			t.Error(fmt.Sprintf("ColemanLiauIndex(%s) != %f, got %f", v.in, v.want, r))
		}
	}
}


func compareFloats(a, b float64) bool {
	return a > (1-floatErr)*b && a < (1+floatErr)*b
}
