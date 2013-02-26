package goGradeLevel

import (
	"log"
	"testing"
	"code.google.com/p/go-html-transform/h5"
)

func Test_getRemoteText(t *testing.T) {
	// fmt.Println(getRemoteText("http://apod.nasa.gov/apod/ap130220.html"))
	// log.Print(getRemoteText("http://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_test"))
	// fmt.Println(getRemoteText("http://www.wired.com/wiredscience/2013/02/how-mathematical-research-is-making-the-life-of-pi-tiger-even-better/"))
}

func Test_analyzeUrl(t *testing.T) {
	var d = []struct{
		in string
		want string
	} {
		{"http://apod.nasa.gov/apod/ap130220.html", ""},
	}	

	for _, v := range d {
		_, err := AnalyzeUrl(v.in)
		if err != nil {
			t.Error("%s threw a %s", v.in, err.Error())
		}
	}
}

func Test_endSentenceWithNode(t *testing.T) {
	var d = []struct{
		in string
		want bool
	} {
		{"<li><a>Privacy policy</a></li>", true},
		{"<a>Privacy policy</a>", false},
		{"<p>Text <i><a>Anchor</a></i> </p>", true},
		{"<i><a>Anchor</a></i>", false},
	}

	for _, v := range d {
		tree := parseString(v.in)
		if endSentenceWithNode(tree) != v.want {
			t.Errorf("%s != %t", v.in, v.want)
		}
	}
}

func Test_collectText(t *testing.T) {
	var d = []struct{
		in string
		want []string
	} {
		{"<li><a>A</a></li>", []string{"A", eos}},
		{"<a>A</a>", []string{"A"}},
		{"<p>P<i><a>A</a></i></p>", []string{"P", "A", eos}},
		{"<i><a>A</a></i>", []string{"A"}},
	}

	for _, v := range d {

		text := collectText(parseString(v.in))
		if !compareStringArray(text, v.want) {
			t.Errorf("%s != %v, got: %v", v.in, v.want, text)
		}
	}
}


func parseString(s string) *h5.Node {
	p := h5.NewParserFromString(s)

	err := p.Parse()
	if err != nil {
		log.Fatal(err.Error())
	}
	return p.Tree()
}

func compareStringArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}