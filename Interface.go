package scarlet

import (
	"io"
	"strings"

	"github.com/aerogo/codetree"
)

// Compile compiles the given scarlet code to a CSS string.
func Compile(reader io.Reader, pretty bool) (string, error) {
	tree, err := codetree.New(reader)

	if err != nil {
		return "", err
	}

	defer tree.Close()
	output := strings.Builder{}
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

		for _, name := range state.VariableNames {
			value := state.Variables[name]

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
		rule.Render(&output, pretty)
	}

	// Render animations
	for _, animation := range animations {
		animation.Render(&output, pretty)
	}

	// Render media groups
	for _, mediaGroup := range mediaGroups {
		mediaGroup.Render(&output, pretty)
	}

	// Render media queries
	for _, mediaQuery := range mediaQueries {
		mediaQuery.Render(&output, pretty)
	}

	return strings.TrimRight(output.String(), "\n"), nil
}
