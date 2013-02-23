package goGradeLevel

import (
	"fmt"
	"testing"
)

func Test_isVowel(t *testing.T) {
	var d = []struct {
		in   string
		want bool
	}{
		{"Nick", false},
		{"aNick", true},
		{"", false},
	}

	for _, v := range d {
		if isVowel(v.in) != v.want {
			t.Error(fmt.Sprintf("isVowel(%s) != %v", v.in, v.want))
		}
	}
}

func Test_isDipthong(t *testing.T) {
	var d = []struct {
		in   string
		want bool
	}{
		{"Nick", false},
		{"ouNick", true},
		{"n", false},
	}

	for _, v := range d {
		if isDipthong(v.in) != v.want {
			t.Error(fmt.Sprintf("isDipthong(%s) != %t", v.in, v.want))
		}
	}
}

func Test_isDouble(t *testing.T) {
	var d = []struct {
		in   string
		want bool
	}{
		{"ooLite", true},
		{"pickle", false},
		{"n", false},
	}

	for _, v := range d {
		if isDouble(v.in) != v.want {
			t.Error(fmt.Sprintf("isDouble(%s) != %t", v.in, v.want))
		}
	}
}

func Test_skip(t *testing.T) {
	var d = []struct {
		in   string
		s    int
		want string
	}{
		{"Nick", 1, "ick"},
		{"Pickle", 2, "ckle"},
		{"N", 2, ""},
	}

	for _, v := range d {
		if skip(v.in, v.s) != v.want {
			t.Error(fmt.Sprintf("skip(%s, %d) != %s", v.in, v.s, v.want))
		}
	}
}

func Test_endsWithSilentVowel(t *testing.T) {
	var d = []struct {
		in   string
		want bool
	}{
		{"mumble", false},
		{"tickle", false},
		{"rhyme", true},
	}

	for _, v := range d {
		if endsWithSilentVowel(v.in) != v.want {
			t.Error(fmt.Sprintf("endsWithSilentVowel(%s) != %t", v.in, v.want))
		}
	}
}

func Test_CountSyllables(t *testing.T) {
	var d = []struct {
		in   string
		want int
	}{
		{"nick", 1},
		{"NICk", 1},
		{"oolite", 2},
		{"oOlIte", 2},
	}

	for _, v := range d {
		if CountSyllables(v.in) != v.want {
			t.Error(fmt.Sprintf("CountSyllables(%s) != %d, got: %d", v.in, v.want, CountSyllables(v.in)))
		}
	}
}
