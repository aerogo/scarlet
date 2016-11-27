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

	var output bytes.Buffer
	rules := compileChildren(tree, nil)
	rules = combineDuplicates(rules)

	for _, rule := range rules {
		if len(rule.Statements) == 0 {
			continue
		}

		output.WriteString(rule.SelectorPath(pretty))

		if len(rule.Duplicates) > 0 {
			for _, duplicate := range rule.Duplicates {
				output.WriteString(",")

				if pretty {
					output.WriteString(" ")
				}

				output.WriteString(duplicate.SelectorPath(pretty))
			}
		}

		if pretty {
			output.WriteString(" ")
		}

		output.WriteString("{")

		if pretty {
			output.WriteString("\n")
		}

		for index, statement := range rule.Statements {
			if pretty {
				output.WriteString("\t")
			}

			statement := compileStatement(statement, pretty)

			// Remove last semicolon
			if index == len(rule.Statements)-1 && !pretty {
				statement = statement[:len(statement)-1]
			}

			output.WriteString(statement)

			if pretty {
				output.WriteString("\n")
			}
		}

		output.WriteString("}")

		if pretty {
			output.WriteString("\n\n")
		}
	}

	return strings.TrimRight(output.String(), "\n"), nil
}
