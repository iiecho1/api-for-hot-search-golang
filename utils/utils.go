package utils

import (
	"fmt"
	"regexp"
)

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func ExtractMatches(text, pattern string) [][]string {
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(text, -1)
	return matches
}
