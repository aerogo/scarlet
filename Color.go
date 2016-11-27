package scarlet

import (
	"regexp"
	"strconv"
	"strings"

	colorful "github.com/lucasb-eyer/go-colorful"
)

var rgbRegex = regexp.MustCompile("rgb\\((.*?)\\)")

func colorComponentToFloat(value string) float64 {
	if strings.HasSuffix(value, "%") {
		normalized, _ := strconv.ParseFloat(strings.TrimSpace(value[:len(value)-1]), 64)
		return normalized
	}

	asByte, err := strconv.Atoi(value)

	if err != nil {
		panic("Invalid RGB color value: " + value)
	}

	return float64(asByte) / 255.0
}

// optimizeColors optimizes color values by converting them to HEX format.
func optimizeColors(value string) string {
	matches := rgbRegex.FindAllStringSubmatch(value, -1)
	for _, match := range matches {
		rgbFunction := match[0]
		rgbValuesString := match[1]
		rgb := strings.Split(rgbValuesString, ",")

		r := strings.TrimSpace(rgb[0])
		g := strings.TrimSpace(rgb[1])
		b := strings.TrimSpace(rgb[2])

		color := colorful.Color{
			R: colorComponentToFloat(r),
			G: colorComponentToFloat(g),
			B: colorComponentToFloat(b),
		}

		value = strings.Replace(value, rgbFunction, color.Hex(), 1)
	}

	return value
}
