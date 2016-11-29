package scarlet

// Mixin is a collection of CSS rules.
type Mixin struct {
	Root  *CSSRule
	Rules []*CSSRule
}

// Apply ...
func (mixin *Mixin) Apply(parent *CSSRule) []*CSSRule {
	parent.Statements = append(parent.Statements, mixin.Root.Statements...)

	rules := []*CSSRule{}

	for _, rule := range mixin.Rules {
		cpy := &CSSRule{
			Selector:   rule.Selector,
			Statements: rule.Statements,
			Parent:     parent,
		}
		rules = append(rules, cpy)
	}

	return rules
}
