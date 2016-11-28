package scarlet

// CSSStatement ...
type CSSStatement struct {
	Property string
	Value    string
}

type byProperty []*CSSStatement

func (c byProperty) Len() int {
	return len(c)
}

func (c byProperty) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c byProperty) Less(i, j int) bool {
	return c[i].Property < c[j].Property
}
