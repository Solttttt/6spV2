package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseTranslater(inStr string) (string, error) {
	cleanStr := strings.TrimSpace(inStr)

	if cleanStr == "" {
		return "", errors.New("empty document")
	}

	hasNonMorseChars := false
	for _, r := range cleanStr {
		if r != '.' && r != '-' && r != ' ' {
			hasNonMorseChars = true
			break
		}
	}

	if hasNonMorseChars {
		result := morse.ToMorse(strings.ToUpper(cleanStr))
		return result, nil
	} else {
		result := morse.ToText(cleanStr)
		return result, nil
	}
}
