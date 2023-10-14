package actions

import (
	"os"

	"github.com/on3dd/wot-nicknames-scrapper/internal/parser"
	"github.com/on3dd/wot-nicknames-scrapper/internal/writer"

	"github.com/urfave/cli/v2"
)

func Parse(cCtx *cli.Context) error {
	parserInstance := parser.Init()
	words := parserInstance.Parse()

	fileUrl := cCtx.String("file")
	file, err := os.Create(fileUrl)

	if err != nil {
		return err
	}

	if err := writer.Write(words, file); err != nil {
		return err
	}

	return nil
}
