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

	// Deep copy every rule
	for _, rule := range mixin.Rules {
		var previous *CSSRule

		for {
			cpy := rule.Copy()

			if previous != nil {
				previous.Parent = cpy
			} else {
				rules = append(rules, cpy)
			}

			if cpy.Parent == mixin.Root {
				cpy.Parent = parent
				break
			}

			previous = cpy
			rule = cpy.Parent
		}
	}

	return rules
}
