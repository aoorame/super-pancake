// Package service предоставляет инструменты для автоматического определения формата входных данных и их конвертации между обычным текстом и кодом Морзе.
package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// AutoDecodeOrEncode принимает входную строку, автоматически определяет тип данных
// (обычный текст или код Морзе) и выполняет соответствующую конвертацию.
// Если строка содержит хотя бы одну букву, она переводится в код Морзе.
// В противном случае строка интерпретируется как код Морзе и декодируется в текст.
// Возвращает ошибку, если входная строка пуста или состоит только из пробельных символов.
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
