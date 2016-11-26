package scarlet

import (
	"bytes"

	"github.com/aerogo/codetree"
)

// Compile compiles the given scarlet code to a CSS string.
func Compile(src string) string {
	tree, err := codetree.New(src)

	if err != nil {
		panic(err)
	}

	var output bytes.Buffer

	for _, rule := range getRules(tree) {
		output.WriteString(rule.Selector)
		output.WriteString("\n")
	}

	return output.String()
}
