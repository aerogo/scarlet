package scarlet_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/aerogo/scarlet"
	"github.com/akyoto/color"
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
	testFile(t, "testdata/test.scarlet")
}

func TestCompiler2(t *testing.T) {
	testFile(t, "testdata/test2.scarlet")
}

func BenchmarkCompiler(b *testing.B) {
	src, _ := ioutil.ReadFile("testdata/test.scarlet")
	code := string(src)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := scarlet.Compile(code, false)

			if err != nil {
				b.Fail()
			}
		}
	})
}
