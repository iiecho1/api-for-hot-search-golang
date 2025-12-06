package utils

import (
	"regexp"
)

func ExtractMatches(text, pattern string) [][]string {
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(text, -1)
	return matches
}
