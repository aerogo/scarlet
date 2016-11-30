package scarlet

import "bytes"

// MediaGroup ...
type MediaGroup struct {
	Operator string
	Size     string
	Property string
	Rules    []*CSSRule
}

// Render renders the media group into the output stream.
func (media *MediaGroup) Render(output *bytes.Buffer, pretty bool) {
	output.WriteString("@media all and (")

	switch media.Operator {
	case "<":
		output.WriteString("max")
	case ">":
		output.WriteString("min")
	default:
		panic("Invalid screen size operator in media query")
	}

	output.WriteByte('-')
	output.WriteString(media.Property)
	output.WriteByte(':')

	if pretty {
		output.WriteByte(' ')
	}

	output.WriteString(media.Size)
	output.WriteByte(')')

	if pretty {
		output.WriteByte(' ')
	}

	output.WriteByte('{')

	if pretty {
		output.WriteByte('\n')
	}

	for _, rule := range media.Rules {
		rule.Render(output, pretty)
	}

	output.WriteByte('}')

	if pretty {
		output.WriteByte('\n')
	}
}
