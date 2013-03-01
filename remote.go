package goGradeLevel

import (
	"net/http"
	"strings"
	"code.google.com/p/go-html-transform/h5"
)

/* remote error type */
type RemoteError struct {
	msg string
}

func (e *RemoteError) Error() string {
	return e.msg
}

func NewRemoteError(msg string) *RemoteError {
	return &RemoteError{msg}
}

type RemoteResult struct {
	url string
	result *Result
}

var validElementTypes = []string{
	"title", "div", "span", "h1", "h2", "h3",
	"h4", "h5", "h6", "li", "a", "p", "em", "i", 
}

var endSentenceElementTypes = []string {
	"li", "div", "p", "h1", "h2", "h3", "h4", "h5", "h6",
}

var eos = ". "

func AnalyzeUrl(url string) *Result {
	text, err := getRemoteText(url)
	if (err != nil) {
		return &Result{}
	}
	res := Analyze(text)
	return res
}

func AnalyzeUrls(in []string) map[string]*Result {
	ch := make(chan *RemoteResult, len(in))
	for _, v := range in {
		go func (url string) {
			s := AnalyzeUrl(url)
			ch <- &RemoteResult{url: url, result: s}
			}(v)
	}
	ret := make(map[string]*Result)
	for i := 0; i < len(in); i ++  {
		r := <- ch		
		ret[r.url] = r.result
	}
	return ret
}

func getRemoteText(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	p := h5.NewParser(resp.Body)

	err = p.Parse()
	if err != nil {
		return "", NewRemoteError("Parse Error")
	}

	found := collectText(p.Tree())
	return strings.Replace(strings.Join(found, ""), "\n", " ", -1), nil
}

func walk(n *h5.Node, pre func(*h5.Node), post func(*h5.Node)) {
	pre(n)
	if len(n.Children) > 0 {
		for _, v := range n.Children {
			walk(v, pre, post)
		}
	}
	post(n)
}

func collectText (tree *h5.Node) []string {
	found := make([]string, 0)
    walk(tree, func(n *h5.Node) {
    	// do nothing pre
    }, func(n *h5.Node) {
    	if n.Type == h5.TextNode && validNodeType(n.Parent) && strings.TrimSpace(n.Data()) != "" {
    		if endSentenceWithNode(n.Parent) {
    			found = append(found, n.Data())
    		} else {
    			found = append(found, n.Data())    			
    		}
    	}
    	if endSentenceWithNode(n) {
    		if len(found) > 0 && found[len(found)-1] != eos {
    			found = append(found, eos)	
    		}	
    	}
	})
	return found
}

func validNodeType(n *h5.Node) bool {
	switch n.Type {
	case h5.TextNode: 
		return true
	case h5.ElementNode:
		for _, v := range validElementTypes {
			if v == n.Data() {
				return true
			}
		} 
	}
	return false
}

func endSentenceWithNode(n *h5.Node) bool {
	for _, v := range endSentenceElementTypes {
		if v == n.Data() {//  || (n.Parent != nil && v == n.Parent.Data()) {
			return true
		}
	}
	return false
}