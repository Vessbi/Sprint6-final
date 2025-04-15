package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvMorseString(text string) (string, error) {
	var textOrmorse string
	if text == "" {
		return "", morse.ErrNoEncoding{Text: "empty string"}
	}
	// remove all special characters
	//reg := regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	//newText := reg.ReplaceAllString(text, "")
	truthCheck := func(r rune) bool {
		return r == '.' || r == '-'
	}
	if strings.ContainsFunc(text, truthCheck) {
		textOrmorse = morse.ToText(text)
	} else {
		textOrmorse = morse.ToMorse(text)
	}

	/*if newText == "" {
		// ToText converts a morse string to his textual representation, it is an alias to DefaultConverter.ToText.
		textOrmorse = morse.ToText(text)
	} else {
		// ToMorse converts a text to his morse rrpresentation, it is an alias to DefaultConverter.ToMorse.
		textOrmorse = morse.ToMorse(text)
	}*/
	return textOrmorse, nil
}
