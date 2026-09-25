package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDecodeOrEncode(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("пустые входные данные")
	}
	hasLetters := false
	for _, r := range trimmed {
		if unicode.IsLetter(r) {
			hasLetters = true
			break
		}
	}
	if hasLetters {
		return morse.ToMorse(trimmed), nil
	}
	return morse.ToText(trimmed), nil
}
