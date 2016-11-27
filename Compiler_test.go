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
	styl, _ := ioutil.ReadFile("test.styl")

	start := time.Now()
	css, _ := Compile(string(styl), false)
	elapsed := time.Since(start)

	fmt.Println(css)
	fmt.Println()
	fmt.Println("Lines:", color.YellowString("%d", len(strings.Split(css, "\n"))))
	fmt.Println("Size: ", color.YellowString("%d", len(css)))
	fmt.Println("Time: ", color.GreenString("%v", elapsed))
}
