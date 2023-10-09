package actions

import (
	"os"

	"github.com/on3dd/wot-nicknames-scrapper/internal/parser"
	"github.com/on3dd/wot-nicknames-scrapper/internal/writer"

	"github.com/urfave/cli/v2"
)

func Parse(cCtx *cli.Context) error {
	fileUrl := cCtx.String("file")

	if cCtx.NArg() > 0 {
		fileUrl = cCtx.Args().First()
	}

	parserInstance := parser.Init()

	parserInstance.Parse()

	words := <-parserInstance.Words

	file, err := os.OpenFile(fileUrl, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return err
	}

	if err := writer.Write(words, file); err != nil {
		return err
	}

	// gen.GenerateTestData(words, 10)

	return nil
}
