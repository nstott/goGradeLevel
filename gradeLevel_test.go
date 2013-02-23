package goGradeLevel

import (
	"fmt"
	"testing"
)

var floatErr = 0.0001

func Test_FleschReadingEase(t *testing.T) {
	var d = []struct {
		in   string
		want float64
	}{
		{"I LOVE PIE", 90.990000},
		{"Flesch Grade Level Readability Formula improves upon the Flesch Reading Ease Readability Formula. " +
			"Rudolph Flesch, an author, writing consultant, and the supporter of Plain English Movement, " +
			"is the co-author of this formula along with John P. Kincaid. That’s why it is also called Flesch-Kincaid " +
			"Grade Level Readability Test. Raised in Austria, Flesch studied law and earned a Ph.D. " +
			"in English from the Columbia University. Flesch, through his writings and speeches, advocated a return to phonics. " +
			"In his article, A New Readability Yardstick, published in the Journal of Applied Psychology in 1948, Flesch proposed " +
			"the Reading Ease Readability Formula. ", 43.329821},
	}

	for _, v := range d {
		r := FleschReadingEase(v.in)
		if !compareFloats(r, v.want) {
			t.Error(fmt.Sprintf("FleschReadingEase(%s) != %f, got %f", v.in, v.want, r))
		}
	}
}

func Test_FleschKincaidGradeLevel(t *testing.T) {
	var d = []struct {
		in   string
		want float64
	}{
		{"I LOVE PIE", 1.313333},
		{"Flesch Grade Level Readability Formula improves upon the Flesch Reading Ease Readability Formula. " +
			"Rudolph Flesch, an author, writing consultant, and the supporter of Plain English Movement, " +
			"is the co-author of this formula along with John P. Kincaid. That’s why it is also called Flesch-Kincaid " +
			"Grade Level Readability Test. Raised in Austria, Flesch studied law and earned a Ph.D. " +
			"in English from the Columbia University. Flesch, through his writings and speeches, advocated a return to phonics. " +
			"In his article, A New Readability Yardstick, published in the Journal of Applied Psychology in 1948, Flesch proposed " +
			"the Reading Ease Readability Formula.", 10.258929},
	}

	for _, v := range d {
		r := FleschKincaidGradeLevel(v.in)
		if !compareFloats(r, v.want) {
			t.Error(fmt.Sprintf("FleschKincaidGradeLevel(%s) != %f, got %f", v.in, v.want, r))
		}
	}
}

func compareFloats(a, b float64) bool {
	return a > (1-floatErr)*b && a < (1+floatErr)*b
}
