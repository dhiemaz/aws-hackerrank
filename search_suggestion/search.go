package main

import (
	"sort"
	"strings"
)

func searchSuggestion(repository []string, customerQuery string) []string {
	if len(customerQuery) < 2 {
		return nil
	}

	var (
		search       string
		matchedArray []string
	)

	for i := 0; i < len(customerQuery); i++ {
		search = search + string(customerQuery[i])
		var tempMatchedArray []string
		if len(search) > 1 {
			for _, v := range repository {
				if strings.HasPrefix(strings.ToLower(v), strings.ToLower(search)) {
					tempMatchedArray = append(tempMatchedArray, strings.ToLower(v))
				}
			}
		}

		if len(tempMatchedArray) > 0 {
			sort.Strings(tempMatchedArray)
			matchedArray = append(matchedArray, tempMatchedArray...)
		}
	}

	return matchedArray
}
