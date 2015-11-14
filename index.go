package dowser

import (
//	"fmt"
)

type Index struct {
	inverseIndex map[string]int
}

func (this *Index) Query(query string) []int {

	// parse as query language parser

	// Tokenize with tokenizer
	tokenizer := NewSpaceTokenizer()
	tokenForSearch := tokenizer.Tokenize(query)

	// search for each token

	return []int{1}
}
func (this *Index) GetDocument(position int) *DowserResponse {
	return nil
}
