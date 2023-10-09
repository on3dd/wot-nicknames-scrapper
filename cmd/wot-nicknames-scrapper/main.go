package main

import (
	"fmt"

	"github.com/on3dd/wot-nicknames-scrapper/internal/gen"
	"github.com/on3dd/wot-nicknames-scrapper/internal/parser"
)

func main() {
	parserInstance := parser.Init()

	parserInstance.Parse()

	select {
	case words := <-parserInstance.Words:
		gen.GenerateTestData(words, 10)
	default:
		fmt.Println("Waiting for response...")
	}
}
