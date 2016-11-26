package scarlet

import (
	"strings"

	"github.com/aerogo/codetree"
)

// CSSRule ...
type CSSRule struct {
	Selector   string
	Statements []string
	Parent     *CSSRule
}

// compileChildren returns the CSS rules for a given code tree.
// It iterates over the child nodes and finds the CSS rules.
func compileChildren(node *codetree.CodeTree, parent *CSSRule) []*CSSRule {
	var rules []*CSSRule

	for _, child := range node.Children {
		if len(child.Children) > 0 {
			// Child rule
			rule := &CSSRule{
				Selector: child.Line,
				Parent:   parent,
			}

			rules = append(rules, rule)

			childRules := compileChildren(child, rule)
			for _, childRule := range childRules {
				rules = append(rules, childRule)
			}
		} else if parent != nil {
			// Definitions
			parent.Statements = append(parent.Statements, child.Line)
		}
	}

	return rules
}

// compileStatement compiles a Scarlet statement to CSS.
func compileStatement(statement string) string {
	space := strings.IndexByte(statement, ' ')

	if space == -1 {
		panic("Invalid statement: " + statement)
	}

	return statement[:space] + ":" + statement[space:] + ";"
}
