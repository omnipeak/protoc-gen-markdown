package utils

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func StringGTLengthHelper(v *int, s string) {
	if utf8.RuneCountInString(s) > *v {
		*v = utf8.RuneCountInString(s)
	}
}

func PadRight(s string, padWith string, wantedLength int) string {
	output := s

	if padWith == "" {
		padWith = " "
	}

	if utf8.RuneCountInString(s) > wantedLength {
		output = s[:wantedLength]
	}

	if utf8.RuneCountInString(output) < wantedLength {
		output += strings.Repeat(padWith, wantedLength-utf8.RuneCountInString(output))
	}

	return output
}

func BoolToTickOrCross(b bool) string {
	if b {
		return "✅"
	}

	return "❌"
}

var commentStripRegex = regexp.MustCompile(`(?m)^\s*//\s*`)
var newlineRegex = regexp.MustCompile(`(?m)\s*\n\s*`)

func FlattenComment(s string) string {
	return strings.Trim(
		newlineRegex.ReplaceAllString(
			commentStripRegex.ReplaceAllString(s, " "),
			" ",
		),
		" \t\r\n",
	)
}

func PluralSuffix(check int, pluralSuffix string, singularSuffix string) string {
	if check == 1 {
		return singularSuffix
	}

	return pluralSuffix
}
