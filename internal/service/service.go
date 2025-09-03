package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// функция которя берет строчку из файла и проверяет ее на наличие кирилицы
// если есть кирилица то текст в морзе, если нет то наоборот

/*
func isTextInput(input string) bool {
	hasLettersOrNumbers := false
	hasNonMorseChars := false

	for _, r := range input {
		switch {
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			hasLettersOrNumbers = true
		case r != '.' && r != '-' && r != ' ' && !unicode.IsSpace(r):
			hasNonMorseChars = true
		}
	}

	return hasLettersOrNumbers || hasNonMorseChars
}
*/

func IsMorseInput(input string) bool {
	input = strings.TrimSpace(input)
	var num int
	for _, char := range input {
		if (string(char) == ".") || (string(char) == "-") {
			num += 1
		}
	}
	if len(input) == num {
		return true
	}
	return false
}

func ReverseMorse(input string) string {
	if IsMorseInput(input) {
		return morse.ToText(input)
	}
	return morse.ToMorse(input)
}

//сделать интерфейс с пакетом морзе
//сохранить файл
