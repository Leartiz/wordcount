package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(0)
		return
	}

	var sentence = os.Args[1]
	sentence = strings.Trim(sentence, " \n\r")

	var wordCount int = 0
	var runes = []rune(sentence)

	var inWord bool = false
	for _, one := range runes {
		if inWord {
			if !unicode.IsSpace(one) {
				continue
			}
			inWord = false
			continue
		}

		if !unicode.IsSpace(one) {
			inWord = true
			wordCount++
		}
	}

	fmt.Println(wordCount)
}
