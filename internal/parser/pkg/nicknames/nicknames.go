package nicknames

import (
	"regexp"
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

	// Match redundant symbols and cyrillic letters
	rx := regexp.MustCompile(`["'\(\)\,]|[\x{0400}-\x{04FF}]`)

	for _, word := range words {
		tokens := stringToTokens(word, rx)
		lexemes = append(lexemes, tokens...)
	}

	sort.Strings(lexemes)

	return utils.Unique(lexemes)
}

func stringToTokens(word string, rx *regexp.Regexp) []string {
	for _, sep := range utils.Separators {
		if strings.Contains(word, sep) {
			tokens := strings.Split(word, sep)

			result := make([]string, 0)

			for _, token := range tokens {
				processed := rx.ReplaceAllString(token, "")

				if len(processed) > 1 {
					result = append(result, processed)
				}
			}

			return result
		}
	}

	return []string{word}
}
