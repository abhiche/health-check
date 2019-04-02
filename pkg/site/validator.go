package site

import (
	"regexp"
)

// IsValidRequestURL check if the string valid url
func IsValidRequestURL(str string) bool {
	matched, _ := regexp.MatchString(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`, str)
	return matched
}
