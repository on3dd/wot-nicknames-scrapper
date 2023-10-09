package nicknames

import (
	"sort"
	"strings"

	"github.com/gocolly/colly"
	"github.com/on3dd/wot-nicknames-scrapper/pkg/utils"
)

func ExtractNicknamesFromUl(e *colly.HTMLElement) []string {
	words := make([]string, 0)

	e.ForEach("li", func(_ int, li *colly.HTMLElement) {
		words = append(words, li.Text)
	})

	return words
}

func ExtractNicknamesFromParagraph(e *colly.HTMLElement) []string {
	return strings.Split(e.Text, "\n")
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
	for _, sep := range utils.Separators {
		if strings.Contains(word, sep) {
			return strings.Split(word, sep)
		}
	}

	return []string{word}
}
