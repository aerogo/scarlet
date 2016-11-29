package scarlet

import (
	"bytes"
	"strings"

	"unicode"

	"github.com/aerogo/codetree"
)

// compileChildren returns the CSS rules for a given code tree.
// It iterates over the child nodes and finds the CSS rules.
func compileChildren(node *codetree.CodeTree, parent *CSSRule, state *State) []*CSSRule {
	// Comments
	if strings.HasPrefix(node.Line, "//") {
		return nil
	}

	rules := []*CSSRule{}
	selectorsOnPreviousLines := []string{}

	for _, child := range node.Children {
		if len(child.Children) > 0 {
			// Mixin
			if strings.HasPrefix(child.Line, "mixin ") {
				name := child.Line[len("mixin "):]

				mixin := &Mixin{
					Root:  &CSSRule{},
					Rules: []*CSSRule{},
				}

				childRules := compileChildren(child, mixin.Root, state)
				for _, childRule := range childRules {
					mixin.Rules = append(mixin.Rules, childRule)
				}

				state.Mixins[name] = mixin
				continue
			}

			// This isn't 100% correct but works in 99.9% of cases.
			// TODO: Make this work for funky stuff like a[href$="a,b"]
			selectors := strings.Split(child.Line, ",")

			// Append selectors from previous lines
			selectors = append(selectors, selectorsOnPreviousLines...)
			selectorsOnPreviousLines = selectorsOnPreviousLines[:0]

			for _, selector := range selectors {
				selector = strings.TrimSpace(selector)

				// Child rule
				rule := &CSSRule{
					Selector: selector,
					Parent:   parent,
				}

				rules = append(rules, rule)

				childRules := compileChildren(child, rule, state)
				rules = append(rules, childRules...)
			}
		} else {
			// Comments
			if strings.HasPrefix(child.Line, "//") {
				continue
			}

			// Selector on previous line
			if strings.HasSuffix(child.Line, ",") {
				selectorsOnPreviousLines = append(selectorsOnPreviousLines, child.Line[:len(child.Line)-1])
				continue
			}

			equal := strings.IndexByte(child.Line, '=')

			if equal != -1 {
				// Variables
				name := strings.TrimSpace(child.Line[:equal])
				value := strings.TrimSpace(child.Line[equal+1:])
				value = insertVariableValues(value, state)
				state.Variables[name] = value
			} else if parent != nil && strings.IndexByte(child.Line, ' ') != -1 {
				// Statements
				statement := compileStatement(child.Line, state)
				parent.Statements = append(parent.Statements, statement)
			} else {
				mixin, exists := state.Mixins[child.Line]

				if exists && parent != nil {
					mixinRules := mixin.Apply(parent)
					rules = append(rules, mixinRules...)
				} else {
					panic("Invalid statement: " + child.Line)
				}
			}
		}
	}

	return rules
}

// compileStatement compiles a Scarlet statement to CSS.
func compileStatement(statement string, state *State) *CSSStatement {
	space := strings.IndexByte(statement, ' ')

	if space == -1 {
		panic("Invalid statement: " + statement)
	}

	value := strings.TrimSpace(statement[space:])

	// Optimize color values
	value = insertVariableValues(value, state)
	value = optimizeColors(value)

	return &CSSStatement{
		Property: statement[:space],
		Value:    value,
	}
}

// insertVariableValues
func insertVariableValues(expression string, state *State) string {
	// EOF
	runes := append([]rune(expression), ' ')
	buffer := bytes.Buffer{}
	ignore := ignoreReader{}
	cursor := 0

	for index, char := range runes {
		if ignore.canIgnore(char) {
			buffer.WriteRune(char)
			cursor = index + 1
			continue
		}

		if char != '-' && (unicode.IsSpace(char) || unicode.IsPunct(char)) {
			if index != cursor {
				token := string(runes[cursor:index])
				value, exists := state.Variables[token]

				if exists {
					buffer.WriteString(value)
				} else {
					buffer.WriteString(token)
				}
			}

			if index == len(runes)-1 {
				break
			}

			buffer.WriteRune(char)
			cursor = index + 1
		}
	}

	return buffer.String()
}
