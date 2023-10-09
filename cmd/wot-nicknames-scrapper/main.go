package main

import (
	"fmt"

	"github.com/gocolly/colly"

	"github.com/on3dd/wot-nicknames-scrapper/internal/gen"
	"github.com/on3dd/wot-nicknames-scrapper/internal/parser"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("conterfrag.ru"),
	)

	words := make([]string, 0)

	c.OnHTML("h2 ~ ul", func(e *colly.HTMLElement) {
		part := parser.ExtractNicknamesFromUl(e)
		words = append(words, part...)
	})

	c.OnHTML("h2 ~ p", func(e *colly.HTMLElement) {
		part := parser.ExtractNicknamesFromParagraph(e)
		words = append(words, part...)
	})

	c.OnScraped(func(r *colly.Response) {
		lexemes := parser.ParseLexemes(words)
		fmt.Printf("Words length: %d \n", len(words))
		fmt.Printf("Lexemes length: %d \n", len(lexemes))

		generateTestData(lexemes, 10)
	})

	c.Visit("https://conterfrag.ru/niki-dlya-tanki-onlayn")
}

func generateTestData(lexemes []string, iterationsNum int) {
	for i := 0; i < iterationsNum; i++ {
		fmt.Printf("Nickname %d: %s \n", i+1, gen.GenerateNickname(lexemes))
	}
}
