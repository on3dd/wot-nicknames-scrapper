package reader

import (
	"encoding/csv"
	"os"
)

func Read(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	return reader.ReadAll()
}
