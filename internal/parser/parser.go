package parser

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/on3dd/wot-nicknames-scrapper/internal/parser/pkg/extract"
	"github.com/on3dd/wot-nicknames-scrapper/internal/parser/pkg/lexemes"
)

type Parser struct {
	collector *colly.Collector
}

func Init() *Parser {
	return &Parser{
		collector: colly.NewCollector(
			colly.AllowedDomains("conterfrag.ru"),
		),
	}
}

func (p *Parser) Parse() []string {
	names := make([]string, 0)
	result := make(chan []string, 1)

	p.collector.OnHTML("h2 ~ ul", func(e *colly.HTMLElement) {
		parts := extract.ExtractNicknamesFromUl(e)
		names = append(names, parts...)
	})

	p.collector.OnHTML("h2 ~ p", func(e *colly.HTMLElement) {
		parts := extract.ExtractNicknamesFromParagraph(e)
		names = append(names, parts...)
	})

	p.collector.OnScraped(func(r *colly.Response) {
		items := lexemes.ParseLexemes(names)

		fmt.Printf("Nicknames length: %d \n", len(names))
		fmt.Printf("Lexemes length: %d \n", len(items))

		result <- items
	})

	p.collector.Visit("https://conterfrag.ru/niki-dlya-tanki-onlayn")

	return <-result
}
