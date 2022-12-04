package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("conterfrag.ru"),
	)

	count := 0
	total := make([]string, 0)

	c.OnHTML("h2 ~ ul", func(e *colly.HTMLElement) {
		count += 1
		part := parseNicknames(e)
		total = append(total, part...)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Total length: %d \n", len(total))
		fmt.Printf("Sections count: %d \n", count)
	})

	c.Visit("https://conterfrag.ru/niki-dlya-world-of-tanks")
}

func parseNicknames(e *colly.HTMLElement) []string {
	nicknames := make([]string, 0)

	e.ForEach("li", func(_ int, li *colly.HTMLElement) {
		nicknames = append(nicknames, li.Text)
	})

	return nicknames
}
