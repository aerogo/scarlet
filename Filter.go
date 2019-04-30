package scarlet

const (
	ElementSelector = iota
	ClassSelector
	IDSelector
	AttributeSelector
	PseudoSelector
	CombinatorSelector
	SeparatorSelector
)

type SelectorPart struct {
	Name string
	Type int
}

func filterTags(rules []*CSSRule, tags []string) []*CSSRule {
	if len(tags) == 0 {
		return rules
	}
	var out []*CSSRule
	for _, rule := range rules {
		rule.Duplicates = filterTags(rule.Duplicates, tags)
		include := true
	loop:
		for _, part := range parseSelector(rule.Selector) {
			if part.Type != ElementSelector {
				continue
			}
			include = false
			for _, t := range tags {
				if t == part.Name {
					include = true
					break loop
				}
			}
		}
		if include {
			out = append(out, rule)
			continue
		}
		if len(rule.Duplicates) > 0 {
			r := rule.Duplicates[0]
			r.Duplicates = append(r.Duplicates, rule.Duplicates[1:]...)
			out = append(out, r)
		}
	}
	return out
}

func parseSelector(selector string) []SelectorPart {

	var (
		name      []byte
		parts     []SelectorPart
		t         int
		lastT     int
		r         byte
		lookahead byte
		attr      bool
	)

	next := func() {
		parts = append(parts, SelectorPart{
			Name: string(name),
			Type: lastT,
		})
		lastT = t
		name = append(name[:0], r)
	}

	for i := 0; i < len(selector); i++ {
		r = selector[i]

		if attr {
			if r == ']' {
				attr = false
			}
			name = append(name, r)
			continue
		}

		if lookahead != 0 {
			if r == lookahead {
				continue
			}
			lookahead = 0
		}

		switch r {
		case ',':
			t = SeparatorSelector
			lookahead = ' '
		case ' ', '>', '+', '~':
			t = CombinatorSelector
		case '.':
			t = ClassSelector
		case '#':
			t = IDSelector
		case '[':
			t = AttributeSelector
			attr = true
		case ':':
			t = PseudoSelector
			lookahead = ':'
		default:
			if t == CombinatorSelector || t == SeparatorSelector {
				t = ElementSelector
				next()
				continue
			}
			name = append(name, r)
			continue
		}
		next()
	}
	next()
	return parts
}
