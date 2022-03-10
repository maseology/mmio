package mmio

import "regexp"

// removes tabs, duplicate spaces, returns (\n)
// from: https://gosamples.dev/remove-duplicate-spaces/
func RemoveWhiteSpaces(s string) string {
	pattern := regexp.MustCompile(`\s+`)
	return pattern.ReplaceAllString(s, " ")
}
