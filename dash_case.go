package main

import (
	"fmt"
	"os"
	"strings"

	"regexp"

	"github.com/mozillazg/go-unidecode"
)

var unwantedCharsRegex = regexp.MustCompile(`(?i)[^a-z0-9\-_]+`)
var duplicateSeparatorRegex = regexp.MustCompile(`-{2,}`)
var leadingTrailingSeparatorRegex = regexp.MustCompile(`^-|-$`)

func dashCase(value string) string {
	ascii_value := unidecode.Unidecode(value)
	withoutUnwantedChars := string(unwantedCharsRegex.ReplaceAll([]byte(ascii_value), []byte("-")))
	withoutDuplicateSeparators := string(duplicateSeparatorRegex.ReplaceAll([]byte(withoutUnwantedChars), []byte("-")))
	withoutLeadingAndTrailingSeparator := string(leadingTrailingSeparatorRegex.ReplaceAll([]byte(withoutDuplicateSeparators), []byte("")))
	return strings.ToLower(withoutLeadingAndTrailingSeparator)
}

func main() {
	inputString := strings.Join(os.Args[1:], " ")
	fmt.Println(dashCase(inputString))
}
