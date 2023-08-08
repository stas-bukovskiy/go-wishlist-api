package validation

import (
	"regexp"
	"strings"
)

func IsValidEmail(s string) bool {
	matchString, err := regexp.MatchString(`^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$`, s)
	if err != nil {
		return false
	}
	return matchString
}

func IsValidImageName(s string) bool {
	matchString, err := regexp.MatchString(`^(.*/)*.+\.(png|jpg|jpeg)$`, strings.ToLower(s))
	if err != nil {
		return false
	}
	return matchString
}
