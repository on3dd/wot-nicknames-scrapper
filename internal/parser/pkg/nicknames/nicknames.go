package nicknames

import (
	"slices"
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

	filtered := filterWords(words)

	for _, word := range filtered {
		tokens := stringToTokens(word)

		lexemes = append(lexemes, tokens...)
	}

	sort.Strings(lexemes)

	return utils.Unique(lexemes)
}

func filterWords(words []string) []string {
	filtered := make([]string, 0)

	for _, item := range words {
		// Skip blacklisted items
		if slices.Contains(utils.BlacklistedItems, item) {
			continue
		}

		filtered = append(filtered, item)
	}

	return filtered
}

func stringToTokens(word string) []string {
	for _, sep := range utils.Separators {
		if strings.Contains(word, sep) {
			tokens := strings.Split(word, sep)

			result := make([]string, 0)

			for _, token := range tokens {
				if len(token) > 1 {
					result = append(result, token)
				}
			}

			return result
		}
	}

	return []string{word}
}
