package writer

import (
	"encoding/csv"
	"io"

	"github.com/on3dd/wot-nicknames-scrapper/pkg/utils"
)

func Write(data []string, w io.Writer) error {
	records := utils.PrepareRecordsForWriting(data)

	writer := csv.NewWriter(w)

	return writer.WriteAll(records)
}
