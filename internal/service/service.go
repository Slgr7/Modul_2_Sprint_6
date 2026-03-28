package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Service(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", fmt.Errorf("пустая строка")
	}

	isMorse := true
	for _, r := range trimmed {
		if r == '.' || r == '-' || r == ' ' || r == '/' {
			continue
		}
		isMorse = false
		break
	}

	if isMorse {

		return morse.ToText(trimmed), nil
	}

	return morse.ToMorse(trimmed), nil
}
