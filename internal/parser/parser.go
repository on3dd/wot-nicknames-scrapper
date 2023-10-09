package parser

import (
	"fmt"

	"github.com/gocolly/colly"

	"github.com/on3dd/wot-nicknames-scrapper/internal/parser/pkg/nicknames"
)

type Parser struct {
	Words     chan []string
	collector *colly.Collector
}

func Init() *Parser {
	return &Parser{
		Words: make(chan []string, 1),
		collector: colly.NewCollector(
			colly.AllowedDomains("conterfrag.ru"),
		),
	}
}

func (p *Parser) Parse() {
	words := make([]string, 0)

	p.collector.OnHTML("h2 ~ ul", func(e *colly.HTMLElement) {
		part := nicknames.ExtractNicknamesFromUl(e)
		words = append(words, part...)
	})

	p.collector.OnHTML("h2 ~ p", func(e *colly.HTMLElement) {
		part := nicknames.ExtractNicknamesFromParagraph(e)
		words = append(words, part...)
	})

	p.collector.OnScraped(func(r *colly.Response) {
		lexemes := nicknames.ParseLexemes(words)

		fmt.Printf("Words length: %d \n", len(words))
		fmt.Printf("Lexemes length: %d \n", len(lexemes))

		p.Words <- words
	})

	p.collector.Visit("https://conterfrag.ru/niki-dlya-tanki-onlayn")
}
