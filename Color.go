package scarlet

import (
	"regexp"
	"strings"
)

var zeroCommaRegex = regexp.MustCompile(`0(\.\d+)`)

// optimizeColors optimizes color values.
func optimizeColors(value string) string {
	// Remove 0 from comma values
	zeroCommaMatches := zeroCommaRegex.FindAllStringSubmatch(value, -1)

	for _, match := range zeroCommaMatches {
		original := match[0]
		replaced := match[1]
		value = strings.Replace(value, original, replaced, 1)
	}

	return value
}
