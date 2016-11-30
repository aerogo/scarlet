package scarlet

import (
	"bytes"
	"strings"

	"github.com/aerogo/codetree"
)

// Compile compiles the given scarlet code to a CSS string.
func Compile(src string, pretty bool) (string, error) {
	tree, err := codetree.New(src)

	if err != nil {
		return "", err
	}

	output := new(bytes.Buffer)
	state := NewState()

	// Parse it
	rules, mediaGroups := compileChildren(tree, nil, state)

	// Combine duplicate rules
	rules = combineDuplicates(rules)

	// Render to output
	for _, rule := range rules {
		rule.Render(output, pretty)
	}

	for _, mediaGroup := range mediaGroups {
		mediaGroup.Render(output, pretty)
	}

	return strings.TrimRight(output.String(), "\n"), nil
}
