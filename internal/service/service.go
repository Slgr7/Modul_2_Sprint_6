package service

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Service(input string) (string, error) {

	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("пустая строка")
	}

	isMorse := true

	runes := []rune(input)
	for _, r := range runes {
		if r == '.' || r == '-' || unicode.IsSpace(r) {
			continue
		}

		isMorse = false
		break
	}

	if isMorse == true {
		text := morse.ToText(input)
		return text, nil
	}
	textmorse := morse.ToMorse(input)
	return textmorse, nil

}
