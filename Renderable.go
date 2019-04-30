package scarlet

import "strings"

// Renderable represents anything that can be rendered into final output.
type Renderable interface {
	Render(*strings.Builder, bool)
}
