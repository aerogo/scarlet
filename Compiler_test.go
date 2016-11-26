package scarlet

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestCompiler(t *testing.T) {
	styl, _ := ioutil.ReadFile("test.styl")

	start := time.Now()
	css, _ := Compile(string(styl))
	elapsed := time.Since(start)

	fmt.Println(css)
	fmt.Println("\nCompile time:", elapsed)
}
