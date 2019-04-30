package scarlet_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aerogo/codetree"
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

func TestCompilerFilterTags(t *testing.T) {
	tags := []string{"a", "b", "body", "br", "button", "div", "fieldset", "footer", "form", "h1", "h2", "head",
		"header", "html", "iframe", "input", "legend", "li", "link", "meta", "noscript", "ol", "p", "pre", "script",
		"span", "string", "style", "table", "textarea", "title", "ul"}
	reader, _ := os.Open("testdata/normalize.scarlet")
	tree, err := codetree.FromReader(reader)
	defer reader.Close()
	if err != nil {
		t.Fatal(err)
		return
	}
	css, err := scarlet.FromCodeTree(tree).FilterTags(tags).Render(false)
	if err != nil {
		t.Fatalf("Error compiling:%s", err)
		return
	}
	expected := `footer,header{display:block}html{-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%;font-size:100%;overflow-y:scroll}body{font-size:13px;line-height:1.231;margin:0}body,button,input,textarea{color:#222;font-family:sans-serif}a{color:#00e}b{font-weight:bold}pre{_font-family:'courier new',monospace;font-family:monospace,serif;font-size:1em}pre{white-space:pre;white-space:pre-wrap;word-wrap:break-word}ul,ol{margin:1em 0;padding:0 0 0 40px}nav ul,nav ol{list-style:none;list-style-image:none;margin:0;padding:0}form{margin:0}fieldset{border:0;margin:0;padding:0}legend{*margin-left:-7px;border:0}button,input,textarea{*vertical-align:middle;font-size:100%;margin:0;vertical-align:baseline}button,input{*overflow:visible;line-height:normal}table button,table input{*overflow:auto}button,input[type="button"],input[type="reset"],input[type="submit"]{-webkit-appearance:button;cursor:pointer}textarea{overflow:auto;resize:vertical;vertical-align:top}table{border-collapse:collapse;border-spacing:0}@media print{a,a:visited{text-decoration:underline}a[href]:after{content:" (" attr(href) ")"}a[href^="javascript:"]:after,a[href^="#"]:after{content:""}pre{border:1px solid #999;page-break-inside:avoid}p,h2{orphans:3;widows:3}h2{page-break-after:avoid}}`
	if css != expected {
		t.Errorf("Unexpected output: %s", css)
	}
}
