package scarlet

// combineDuplicates combines duplicate CSS statements into one.
// Example:
// a { color: blue; }
// p { color: blue; }
// becomes:
// a, p { color: blue; }
func combineDuplicates(rules []*CSSRule) []*CSSRule {
	result := []*CSSRule{}
	seen := map[uint64]*CSSRule{}

	for _, rule := range rules {
		hash := rule.StatementsHash()

		if existing, ok := seen[hash]; !ok {
			result = append(result, rule)
			seen[hash] = rule
		} else {
			existing.Duplicates = append(existing.Duplicates, rule)
		}
	}

	return result
}
