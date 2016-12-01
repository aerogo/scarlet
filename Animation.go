package scarlet

import "bytes"

// Animation ...
type Animation struct {
	Name      string
	Keyframes []*CSSRule
}

// Render renders the animation to the output stream.
func (anim *Animation) Render(output *bytes.Buffer, pretty bool) {
	output.WriteString("@keyframes ")
	output.WriteString(anim.Name)

	if pretty {
		output.WriteByte(' ')
	}

	output.WriteByte('{')

	if pretty {
		output.WriteByte('\n')
	}

	for _, keyframe := range anim.Keyframes {
		keyframe.Render(output, pretty)
	}

	output.WriteByte('}')

	if pretty {
		output.WriteByte('\n')
	}
}
