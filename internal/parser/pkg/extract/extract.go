package extract

import (
	"strings"

	"github.com/gocolly/colly"
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
