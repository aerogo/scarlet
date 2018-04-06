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
	if len(state.Variables) > 0 {
		if pretty {
			output.WriteString(":root {\n")
		} else {
			output.WriteString(":root{")
		}

		for name, value := range state.Variables {
			if pretty {
				output.WriteString("\t")
			}

			output.WriteString("--")
			output.WriteString(name)
			output.WriteByte(':')

			if pretty {
				output.WriteString(" ")
			}

			output.WriteString(value)
			output.WriteByte(';')

			if pretty {
				output.WriteString("\n")
			}
		}

		output.WriteString("}")

		if pretty {
			output.WriteString("\n\n")
		}
	}

	// Render rules
	for _, rule := range rules {
		rule.Render(output, pretty)
	}

	// Render animations
	for _, animation := range animations {
		animation.Render(output, pretty)
	}

	// Render media groups
	for _, mediaGroup := range mediaGroups {
		mediaGroup.Render(output, pretty)
	}

	// Render media queries
	for _, mediaQuery := range mediaQueries {
		mediaQuery.Render(output, pretty)
	}

	return strings.TrimRight(output.String(), "\n"), nil
}
