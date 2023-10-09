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
	names := make([]string, 0)

	p.collector.OnHTML("h2 ~ ul", func(e *colly.HTMLElement) {
		parts := nicknames.ExtractNicknamesFromUl(e)
		names = append(names, parts...)
	})

	p.collector.OnHTML("h2 ~ p", func(e *colly.HTMLElement) {
		parts := nicknames.ExtractNicknamesFromParagraph(e)
		names = append(names, parts...)
	})

	p.collector.OnScraped(func(r *colly.Response) {
		lexemes := nicknames.ParseLexemes(names)

		fmt.Printf("Nicknames length: %d \n", len(names))
		fmt.Printf("Lexemes length: %d \n", len(lexemes))

		p.Words <- lexemes
	})

	p.collector.Visit("https://conterfrag.ru/niki-dlya-tanki-onlayn")
}
