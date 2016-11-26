package scarlet

import "github.com/aerogo/codetree"

// CSSRule ...
type CSSRule struct {
	Selector    string
	Definitions []string
}

// getRules returns the CSS rules for a given code tree.
func getRules(node *codetree.CodeTree) []*CSSRule {
	var rules []*CSSRule

	for _, child := range node.Children {
		if len(child.Children) > 0 {
			rule := &CSSRule{
				Selector: child.Line,
			}

			rules = append(rules, rule)

			childRules := getRules(child)
			for _, childRule := range childRules {
				rules = append(rules, childRule)
			}
		} else {
			// ...
		}
	}

	return rules
}
