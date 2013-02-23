# goGradeLevel a tool to calculate the reading level of text


##Info

for more information about the Flesch/Kincaid method, see: 
http://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_test

for more information about Coleman/Liau, see: 
http://en.wikipedia.org/wiki/Coleman-Liau_Index

##Use

```go
package main

import (
	"fmt"
	"github.com/nstott/goGradeLevel"
)

func main() {

	text := "The Flesch/Flesch–Kincaid readability tests are designed to indicate comprehension " + 
			"difficulty when reading a passage of contemporary academic English. There are two tests, " + 
			"the Flesch Reading Ease, and the Flesch–Kincaid Grade Level. Although they use the same " +
			"core measures (word length and sentence length), they have different weighting factors. " +
			"The results of the two tests correlate approximately inversely: a text with a comparatively " + 
			"high score on the Reading Ease test should have a lower score on the Grade Level test."

	gradeLevel := goGradeLevel.FleschKincaidGradeLevel(text)
	ease := goGradeLevel.FleschReadingEase(text)
	clIndex := goGradeLevel.ColemanLiauIndex(text)

	fmt.Printf("Flesch Kincaid GradeLevel %f\n", gradeLevel)
	fmt.Printf("Flesch Reading Ease %f\n", ease)
	fmt.Printf("Coleman Liau Index %f\n", clIndex)

}
```