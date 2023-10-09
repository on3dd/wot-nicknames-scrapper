package actions

import (
	"github.com/on3dd/wot-nicknames-scrapper/internal/gen"
	"github.com/on3dd/wot-nicknames-scrapper/internal/reader"
	"github.com/on3dd/wot-nicknames-scrapper/internal/utils"
	"github.com/urfave/cli/v2"
)

func Generate(cCtx *cli.Context) error {
	data, err := reader.Read(cCtx.String("file"))

	if err != nil {
		return err
	}

	lexemes := utils.PrepareRecordsForGenerating(data)

	gen.GenerateTestData(lexemes, cCtx.Int("count"))

	return nil
}
