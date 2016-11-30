package scarlet

import "bytes"

// Renderable represents anything that can be rendered into final output.
type Renderable interface {
	Render(*bytes.Buffer, bool)
}
