package scarlet

import (
	"strings"

	"github.com/aerogo/codetree"
)

// Compiler represents a compiled Scarlet CodeTree, ready for further processing or rendering.
type Compiler struct {
	state        *State
	rules        []*CSSRule
	mediaGroups  []*MediaGroup
	mediaQueries []*MediaQuery
	animations   []*Animation
}

// FromCodeTree compiles a Scarlet CodeTree to a Compiler
func FromCodeTree(tree *codetree.CodeTree) *Compiler {
	compiler := &Compiler{}
	compiler.state = NewState()

	// Parse it
	compiler.rules, compiler.mediaGroups, compiler.mediaQueries, compiler.animations =
		compileChildren(tree, nil, compiler.state)

	compiler.rules = combineDuplicates(compiler.rules)
	for _, group := range compiler.mediaGroups {
		group.Rules = combineDuplicates(group.Rules)
	}
	for _, query := range compiler.mediaQueries {
		query.Rules = combineDuplicates(query.Rules)
	}

	return compiler
}

// Render returns a CSS string with pretty or compressed formatting, depending on the argument pretty.
func (compiler *Compiler) Render(pretty bool) (string, error) {
	output := strings.Builder{}
	// CSS variables
	if len(compiler.state.Variables) > 0 {
		if pretty {
			output.WriteString(":root {\n")
		} else {
			output.WriteString(":root{")
		}

		for _, name := range compiler.state.VariableNames {
			value := compiler.state.Variables[name]

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
	for _, rule := range compiler.rules {
		rule.Render(&output, pretty)
	}

	// Render animations
	for _, animation := range compiler.animations {
		animation.Render(&output, pretty)
	}

	// Render media groups
	for _, mediaGroup := range compiler.mediaGroups {
		mediaGroup.Render(&output, pretty)
	}

	// Render media queries
	for _, mediaQuery := range compiler.mediaQueries {
		mediaQuery.Render(&output, pretty)
	}

	return strings.TrimRight(output.String(), "\n"), nil
}

// FilterTags removes all tag representations from a Compiler not mentioned in tags.
// Useful for optimising utility stylesheets against a known template base.
func (compiler *Compiler) FilterTags(tags []string) *Compiler {
	compiler.rules = filterTags(compiler.rules, tags)
	for _, group := range compiler.mediaGroups {
		group.Rules = filterTags(group.Rules, tags)
	}
	for _, query := range compiler.mediaQueries {
		query.Rules = filterTags(query.Rules, tags)
	}
	return compiler
}

// Compile compiles the given scarlet code to a CSS string.
func Compile(src string, pretty bool) (string, error) {
	tree, err := codetree.New(src)
	if err != nil {
		return "", err
	}
	compiler := FromCodeTree(tree)
	tree.Close()
	return compiler.Render(pretty)
}
