package scarlet

import (
	"bytes"
	"strings"

	"github.com/aerogo/codetree"
)

// CSSRule ...
type CSSRule struct {
	Selector   string
	Statements []string
	Parent     *CSSRule
}

// SelectorPath returns the selector string for the parent node (recursive, returns absolute path).
func (rule *CSSRule) SelectorPath() string {
	if rule.Parent == nil {
		return rule.Selector
	}

	// Parent path
	var fullPath bytes.Buffer
	fullPath.WriteString(rule.Parent.SelectorPath())

	// Whitespace if needed
	if !strings.HasPrefix(rule.Selector, ":") && !strings.HasPrefix(rule.Selector, "[") {
		fullPath.WriteString(" ")
	}

	// The node's selector
	fullPath.WriteString(rule.Selector)

	return fullPath.String()
}

// compileChildren returns the CSS rules for a given code tree.
// It iterates over the child nodes and finds the CSS rules.
func compileChildren(node *codetree.CodeTree, parent *CSSRule) []*CSSRule {
	// Comments
	if strings.HasPrefix(node.Line, "//") {
		return nil
	}

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
			// Comments
			if strings.HasPrefix(child.Line, "//") {
				continue
			}

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
