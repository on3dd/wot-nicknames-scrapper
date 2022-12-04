package parser

import (
	"sort"
	"strings"

	"github.com/gocolly/colly"
	"github.com/on3dd/wot-nicknames-scrapper/utils"
)

// Typical delimiters from most common to least common.
// For example, string "MaJleHbKuu_MaJlbuK" contains 2 words separated by "_".
var SEPARATORS = []string{" ", "_", "-", "."}

func GetNicknames(e *colly.HTMLElement) []string {
	nicknames := make([]string, 0)

	e.ForEach("li", func(_ int, li *colly.HTMLElement) {
		nicknames = append(nicknames, li.Text)
	})

	return nicknames
}

func ParseLexemes(words []string) []string {
	lexemes := make([]string, 0)

	for _, word := range words {
		tokens := stringToTokens(word)
		lexemes = append(lexemes, tokens...)
	}

	sort.Strings(lexemes)

	return utils.Unique(lexemes)
}

func stringToTokens(word string) []string {
	for _, sep := range SEPARATORS {
		if strings.Contains(word, sep) {
			return strings.Split(word, sep)
		}
	}

	return []string{word}
}
