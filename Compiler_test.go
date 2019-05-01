package scarlet_test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aerogo/scarlet"
	"github.com/akyoto/color"
)

func testFile(t *testing.T, filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	start := time.Now()
	css, _ := scarlet.Compile(file, true)
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
	file, err := os.Open("testdata/test.scarlet")

	if err != nil {
		panic(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := file.Seek(0, io.SeekStart)

			if err != nil {
				panic(err)
			}

			_, err = scarlet.Compile(file, false)

			if err != nil {
				panic(err)
			}
		}
	})
}
