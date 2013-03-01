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

func Test_AnalyzeUrl(t *testing.T) {
	var d = []struct{
		in string
		want *Result
	} {
		{"http://apod.nasa.gov/apod/ap130220.html", &Result{ReadingEase:41.01510209042297,FleschKincaidGradeLevel:11.074987846378221,ColemanLiauIndex:13.415867768595039}},
	}	

	for _, v := range d {
		b := AnalyzeUrl(v.in)
		if !compareResult(b, v.want) {
			t.Errorf("url: %s wanted %v got %v", v.in, v.want, b)
		}
	}
}

func Test_AnalyzeUrls(t *testing.T) {
	var d = []struct{
		in []string
		want map[string]*Result
	} {
		{
			[]string{
				"http://apod.nasa.gov/apod/ap130220.html",
				"http://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_test",
				}, 
			map[string]*Result{
				"http://apod.nasa.gov/apod/ap130220.html": 
					&Result{ReadingEase:41.01510209042297,FleschKincaidGradeLevel:11.074987846378221,ColemanLiauIndex:13.415867768595039},
				"http://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_test": 
					&Result{ReadingEase:39.94790088638197,FleschKincaidGradeLevel:10.4305418057634,ColemanLiauIndex:13.429741248097411},
			},
		},
	}	

	for _, v := range d {
		b := AnalyzeUrls(v.in)
		for kk, vv := range v.want {
			if !compareResult(b[kk], vv) {
				t.Errorf("url: %s, wanted %v,\n got %v", kk, vv, b[kk])				
			}
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

/*
 * utility functions
 */
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

func compareResult(a, b *Result) bool {
	if a == nil || b == nil {
		return false
	}
	if compareFloats(a.ReadingEase, b.ReadingEase) && 
		compareFloats(a.FleschKincaidGradeLevel, b.FleschKincaidGradeLevel) && 
		compareFloats(a.ColemanLiauIndex, b.ColemanLiauIndex) {
			return true
		}
	return false
}