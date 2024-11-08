package lgo

import (
	"fmt"
	"testing"
)

func TestKMPSearch(t *testing.T) {
	text := "ababcababcabc"
	pattern := "abc"
	result := KMPSearch(text, pattern)
	if len(result) == 0 {
		fmt.Println("Pattern not found in text.")
	} else {
		fmt.Printf("Pattern found at indices: %v\n", result)
	}
}
