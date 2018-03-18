package scarlet

import (
	"fmt"
	"testing"
)

func TestVariableInsertion(t *testing.T) {
	state := NewState()
	state.Constants["bg-color"] = "red"

	src := "linear-gradient(to bottom, 0% bg-color, 100% bg-color) 'bg-color'"

	fmt.Println(src)
	fmt.Println(insertVariableValues(src, state))
}
