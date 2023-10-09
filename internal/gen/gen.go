package gen

import (
	"fmt"
	"strings"

	"github.com/on3dd/wot-nicknames-scrapper/pkg/utils"
)

func GenerateNickname(lexemes []string) string {
	length := utils.GetRandomIntInRange(2, 4)

	words := make([]string, 0)

	for i := 0; i < length; i++ {
		words = append(words, utils.GetRandomElementOfSlice(lexemes))
	}

	separator := utils.GetRandomElementOfSlice(utils.Separators)

	return strings.Join(words, separator)
}

func GenerateTestData(lexemes []string, iterationsNum int) {
	for i := 0; i < iterationsNum; i++ {
		fmt.Printf("Nickname %d: %s \n", i+1, GenerateNickname(lexemes))
	}
}
