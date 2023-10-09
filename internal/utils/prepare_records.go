package utils

func PrepareRecordsForWriting(data []string) [][]string {
	records := make([][]string, len(data))

	for i, val := range data {
		records[i] = []string{val}
	}

	return records
}

func PrepareRecordsForGenerating(data [][]string) []string {
	records := make([]string, len(data))

	for i, val := range data {
		records[i] = val[0]
	}

	return records
}
