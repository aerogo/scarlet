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

func testFile(t *testing.T, filePath string, result string) {
	src, _ := ioutil.ReadFile(filePath)
	code := string(src)

	start := time.Now()
	css, _ := scarlet.Compile(code, false)
	elapsed := time.Since(start)

	fmt.Println(css)

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Lines:", color.YellowString("%d", len(strings.Split(css, "\n"))))
	fmt.Println("Size: ", color.YellowString("%d", len(css)))
	fmt.Println("Time: ", color.GreenString("%v", elapsed))

	if css != result {
		t.Error("Unexpected output")
	}
}

func TestCompiler1(t *testing.T) {
	testFile(t, "testdata/test.scarlet", `a{color:black}a:hover,p:hover{color:blue}a:hover div,p:hover div{color:red}@keyframes appear{0%{opacity:0}100%{opacity:1}}@media all and (min-width:900px){p{animation-name:appear;display:flex}}`)
}

func TestCompiler2(t *testing.T) {
	testFile(t, "testdata/test2.scarlet", `:root{--text-color:blue;--text-hover-color:var(--text-color);--gradient:linear-gradient(to bottom,0% var(--text-color),100% var(--text-color));}body{background-color:#202020;color:var(--text-color);display:flex;flex-direction:row}p{background-color:#202020;color:blue;display:flex;flex-direction:row}a,#content,#content:hover{color:red}a:hover{color:var(--text-hover-color)}a:hover div{width:100%}a:hover div img{height:100%}a:active{color:blue}#content>div{color:orange}#content img{border:none}#content[aria-class="button"]{color:green}div:hover,p:hover{color:white}div span,div address,p span,p address,h2,h1{display:none}@media all and (min-height: 320px){body{background-color:red}}`)
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
