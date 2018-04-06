package scarlet

import "bytes"

// Force interface implementation
var _ Renderable = (*MediaQuery)(nil)

// MediaQuery ...
type MediaQuery struct {
	Selector string
	Rules    []*CSSRule
}

// Render renders the media query to the output stream.
func (media *MediaQuery) Render(output *bytes.Buffer, pretty bool) {
	output.WriteString(media.Selector)

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
