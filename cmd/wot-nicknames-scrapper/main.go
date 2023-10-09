package main

import (
	"os"

	"github.com/on3dd/wot-nicknames-scrapper/internal/gen"
	"github.com/on3dd/wot-nicknames-scrapper/internal/parser"
	"github.com/on3dd/wot-nicknames-scrapper/internal/writer"
)

func main() {
	parserInstance := parser.Init()

	parserInstance.Parse()

	words := <-parserInstance.Words

	file, _ := os.OpenFile("./assets/data.csv", os.O_RDWR|os.O_CREATE, 0755)

	writer.Write(words, file)

	gen.GenerateTestData(words, 10)
}
