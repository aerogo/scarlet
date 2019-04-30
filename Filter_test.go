package scarlet

import "testing"

func TestParseSelector(t *testing.T) {

	parts := parseSelector("div div, div>div, div+div~.class#id[attribute~=value]::after:active>")
	wanted := []SelectorPart{
		{Name:"div", Type:ElementSelector},
		{Name:" ", Type:CombinatorSelector},
		{Name:"div", Type:ElementSelector},
		{Name:",", Type:SeparatorSelector},
		{Name:"div", Type:ElementSelector},
		{Name:">", Type:CombinatorSelector},
		{Name:"div", Type:ElementSelector},
		{Name:",", Type:SeparatorSelector},
		{Name:"div", Type:ElementSelector},
		{Name:"+", Type:CombinatorSelector},
		{Name:"div", Type:ElementSelector},
		{Name:"~", Type:CombinatorSelector},
		{Name:".class", Type:ClassSelector},
		{Name:"#id", Type:IDSelector},
		{Name:"[attribute~=value]", Type:AttributeSelector},
		{Name:":after", Type:PseudoSelector},
		{Name:":active", Type:PseudoSelector},
		{Name:">", Type:CombinatorSelector},
	}
	if len(parts) != len(wanted) {
		t.Error("Unexpected number of parts")
	}
	for i, part := range parts {
		want := wanted[i]
		if part.Name != want.Name || part.Type != want.Type {
			t.Errorf("Part %#v != %#v", part, want)
		}
	}
}
