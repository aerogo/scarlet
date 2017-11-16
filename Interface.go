package scarlet

import (
	"bytes"
	"strings"

	"github.com/aerogo/codetree"
)

// Compile compiles the given scarlet code to a CSS string.
func Compile(src string, pretty bool) (string, error) {
	tree, err := codetree.New(src)
	defer tree.Close()

	if err != nil {
		return "", err
	}

	output := new(bytes.Buffer)
	state := NewState()

	// Parse it
	rules, mediaGroups, mediaQueries, animations := compileChildren(tree, nil, state)

	// Combine duplicate rules
	rules = combineDuplicates(rules)

	// CSS variables
	output.WriteString(":root{")
	for name, value := range state.Variables {
		output.WriteString("--")
		output.WriteString(name)
		output.WriteByte(':')
		output.WriteString(value)
		output.WriteByte(';')
	}
	output.WriteString("}")

	// Render to output
	for _, rule := range rules {
		rule.Render(output, pretty)
	}

	for _, animation := range animations {
		animation.Render(output, pretty)
	}

	for _, mediaGroup := range mediaGroups {
		mediaGroup.Render(output, pretty)
	}

	for _, mediaQuery := range mediaQueries {
		mediaQuery.Render(output, pretty)
	}

	return strings.TrimRight(output.String(), "\n"), nil
}
