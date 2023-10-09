package writer

import (
	"encoding/csv"
	"io"
)

func Write(data []string, w io.Writer) error {
	records := prepareRecords(data)

	writer := csv.NewWriter(w)

	return writer.WriteAll(records)
}

func prepareRecords(data []string) [][]string {
	records := make([][]string, len(data))

	for i, val := range data {
		records[i] = []string{val}
	}

	return records
}
