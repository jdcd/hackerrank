package repeatWorlds

import (
	"regexp"
	"strings"
)

type WordsCount map[string]int

func repeatWorlds(text string) WordsCount {
	totalCount := map[string]int{}
	slice := strings.Split(text, " ")
	for _, word := range slice {
		totalCount[normalize(word)]++
	}
	return totalCount
}

func normalize(word string) string {
	return strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(word, ""))
}
