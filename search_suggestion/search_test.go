package main

import (
	"fmt"
	"testing"
)

func TestSearchSuggestion(t *testing.T) {
	t.Run("testing suggestion", func(t *testing.T) {

		var (
			repository    []string = []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
			customerQuery string   = "mouse"
		)

		result := searchSuggestion(repository, customerQuery)
		fmt.Println(result)
		t.Log(result)
	})
}
