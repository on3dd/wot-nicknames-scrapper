package main

import (
	"fmt"

	"github.com/gocolly/colly"

	"github.com/on3dd/wot-nicknames-scrapper/parser"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("conterfrag.ru"),
	)

	words := make([]string, 0)

	c.OnHTML("h2 ~ ul", func(e *colly.HTMLElement) {
		part := parser.GetNicknames(e)
		words = append(words, part...)
	})

	c.OnScraped(func(r *colly.Response) {
		lexemes := parser.ParseLexemes(words)
		fmt.Printf("Words length: %d \n", len(words))
		fmt.Printf("Lexemes length: %d \n", len(lexemes))
	})

	c.Visit("https://conterfrag.ru/niki-dlya-world-of-tanks")
}
