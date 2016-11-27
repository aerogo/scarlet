package scarlet

import (
	"bytes"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/OneOfOne/xxhash"
)

// CSSRule ...
type CSSRule struct {
	Selector   string
	Statements []string
	Duplicates []*CSSRule
	Parent     *CSSRule
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
		if unicode.IsLetter(firstChar) {
			fullPath.WriteString(" ")
			fullPath.WriteString(rule.Selector)
		} else if firstChar == '<' {
			if pretty {
				fullPath.WriteString(" ")
				fullPath.WriteString(rule.Selector)
			} else {
				fullPath.WriteString(strings.Replace(rule.Selector, "< ", "<", 1))
			}
		} else {
			fullPath.WriteString(rule.Selector)
		}

		break
	}

	return fullPath.String()
}

// StatementsHash returns a hash of all the statements which is used to find duplicate CSS rules.
func (rule *CSSRule) StatementsHash() string {
	sort.Strings(rule.Statements)

	hash := xxhash.NewS64(0)

	for _, statement := range rule.Statements {
		hash.WriteString(statement)
	}

	return strconv.FormatUint(hash.Sum64(), 16)
}
