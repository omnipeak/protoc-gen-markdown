package utils

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
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

func PadRightSlice(inputs []string, padWith string, padTo []int) ([]string, error) {
	if len(inputs) != len(padTo) {
		return nil, errors.Errorf(
			"inputs slice length and pad to slice length do not match: %d != %d",
			len(inputs),
			len(padTo),
		)
	}

	output := make([]string, len(inputs))

	for i, v := range inputs {
		output[i] = PadRight(v, padWith, padTo[i])
	}

	return output, nil
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
