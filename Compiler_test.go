package scarlet_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/aerogo/scarlet"
	"github.com/blitzprog/color"
)

func testFile(t *testing.T, filePath string) {
	src, _ := ioutil.ReadFile(filePath)
	code := string(src)

	start := time.Now()
	css, _ := scarlet.Compile(code, true)
	elapsed := time.Since(start)

	fmt.Println(css)

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Lines:", color.YellowString("%d", len(strings.Split(css, "\n"))))
	fmt.Println("Size: ", color.YellowString("%d", len(css)))
	fmt.Println("Time: ", color.GreenString("%v", elapsed))
}

func TestCompiler1(t *testing.T) {
	testFile(t, "test/test.scarlet")
}

func TestCompiler2(t *testing.T) {
	testFile(t, "test/test2.scarlet")
}

func BenchmarkCompiler(b *testing.B) {
	src, _ := ioutil.ReadFile("test/test.scarlet")
	code := string(src)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			scarlet.Compile(code, false)
		}
	})
}
