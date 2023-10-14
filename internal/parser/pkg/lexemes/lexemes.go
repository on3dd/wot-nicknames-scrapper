package lexemes

import (
	"regexp"
	"sort"
	"strings"

	"github.com/on3dd/wot-nicknames-scrapper/pkg/utils"
)

func ParseLexemes(data []string) []string {
	lexemes := make([]string, 0)

	// Match redundant symbols and cyrillic letters
	rx := regexp.MustCompile(`["'\(\)\,]|[\x{0400}-\x{04FF}]`)

	for _, word := range data {
		items := stringToLexemes(word, rx)
		lexemes = append(lexemes, items...)
	}

	sort.Strings(lexemes)

	return utils.Unique(lexemes)
}

func stringToLexemes(data string, rx *regexp.Regexp) []string {
	for _, sep := range utils.Separators {
		if strings.Contains(data, sep) {
			tokens := strings.Split(data, sep)

			result := make([]string, 0)

			for _, token := range tokens {
				processed := rx.ReplaceAllString(token, "")

				if len(processed) > 1 {
					result = append(result, processed)
				}
			}

			return result
		}
	}

	return []string{data}
}
