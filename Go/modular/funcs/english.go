package english

import (
	"strings"
)


func reverse(word string) (result string){
// Assign letters to new variable from the last to the first
	for _,v := range word {
		result = string(v) + result
	}
	return
}

func IsPalin(word string) bool{
	// take the word and reverse
	rev_string := reverse(word)

	// compare
	return strings.EqualFold(word, rev_string)
}