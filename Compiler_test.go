package scarlet

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/fatih/color"
)

func TestCompiler(t *testing.T) {
	code, _ := ioutil.ReadFile("test/test.scarlet")

	start := time.Now()
	css, _ := Compile(string(code), true)
	elapsed := time.Since(start)

	fmt.Println(css)

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Lines:", color.YellowString("%d", len(strings.Split(css, "\n"))))
	fmt.Println("Size: ", color.YellowString("%d", len(css)))
	fmt.Println("Time: ", color.GreenString("%v", elapsed))
}

func TestVariableInsertion(t *testing.T) {
	state := NewState()
	state.Variables["bg-color"] = "red"

	src := "linear-gradient(to bottom, 0% bg-color, 100% bg-color) 'bg-color'"

	fmt.Println(src)
	fmt.Println(insertVariableValues(src, state))
}
