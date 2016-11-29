package scarlet

import (
	"bytes"
	"sort"
	"strconv"
	"strings"

	"github.com/OneOfOne/xxhash"
)

// CSSRule ...
type CSSRule struct {
	Selector   string
	Statements []*CSSStatement
	Duplicates []*CSSRule
	Parent     *CSSRule
}

// Root ...
func (rule *CSSRule) Root() *CSSRule {
	parent := rule

	for {
		nextParent := parent.Parent

		if nextParent == nil {
			return parent
		}

		parent = nextParent
	}
}

// Copy ...
func (rule *CSSRule) Copy() *CSSRule {
	return &CSSRule{
		Selector:   rule.Selector,
		Statements: rule.Statements,
		Parent:     rule.Parent,
	}
}

// SelectorPath returns the selector string for the rule (recursive, returns absolute path).
func (rule *CSSRule) SelectorPath(pretty bool) string {
	if rule.Parent == nil {
		return rule.Selector
	}

	// Parent path
	var fullPath bytes.Buffer
	fullPath.WriteString(rule.Parent.SelectorPath(pretty))

	// Whitespace if needed
	for _, firstChar := range rule.Selector {
		// Sub-elements always have a whitespace
		switch firstChar {
		case '|':
			fullPath.WriteString(" ")
			fullPath.WriteString(rule.Selector[1:])

		case '&':
			fullPath.WriteString(rule.Selector[1:])

		case ':':
			fullPath.WriteString(rule.Selector)

		case '>':
			if pretty {
				fullPath.WriteString(" ")
				fullPath.WriteString(rule.Selector)
			} else {
				fullPath.WriteString("<")
				fullPath.WriteString(strings.TrimSpace(rule.Selector[1:]))
			}

		default:
			fullPath.WriteString(" ")
			fullPath.WriteString(rule.Selector)
		}

		break
	}

	return fullPath.String()
}

// StatementsHash returns a hash of all the statements which is used to find duplicate CSS rules.
func (rule *CSSRule) StatementsHash() string {
	sort.Sort(byProperty(rule.Statements))

	hash := xxhash.NewS64(0)

	for _, statement := range rule.Statements {
		hash.WriteString(statement.Property)
		hash.WriteString(statement.Value)
	}

	return strconv.FormatUint(hash.Sum64(), 16)
}
