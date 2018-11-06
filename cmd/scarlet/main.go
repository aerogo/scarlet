package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var (
	check      bool
	classRegex = regexp.MustCompile(`\.(-?[_a-zA-Z]+[_a-zA-Z0-9-]*)`)
)

// Shell flags
func init() {
	flag.BoolVar(&check, "check", false, "Checks all scarlet classes and tests whether they are referenced in .ts or .go files")
	flag.Parse()
}

func main() {
	if !check {
		flag.Usage()
		return
	}

	checkDirectory(".")
}

func checkDirectory(dir string) {
	wg := sync.WaitGroup{}
	sourceFiles := sync.Map{}
	cssClasses := sync.Map{}

	filepath.Walk(dir, func(file string, f os.FileInfo, err error) error {
		if f.IsDir() || strings.HasPrefix(file, ".") {
			return nil
		}

		if strings.HasSuffix(file, ".scarlet") {
			wg.Add(1)

			go func() {
				defer wg.Done()
				contents, err := ioutil.ReadFile(file)

				if err != nil {
					color.Red(err.Error())
					return
				}

				matches := classRegex.FindAllSubmatch(contents, -1)

				if len(matches) > 0 {
					for _, match := range matches {
						className := string(match[1])
						cssClasses.Store(className, file)
					}
				}
			}()
		}

		if strings.HasSuffix(file, ".go") || strings.HasSuffix(file, ".ts") {
			wg.Add(1)

			go func() {
				defer wg.Done()
				contents, err := ioutil.ReadFile(file)

				if err != nil {
					color.Red(err.Error())
					return
				}

				sourceFiles.Store(file, string(contents))
			}()
		}

		return nil
	})

	// Wait for goroutines to finish
	wg.Wait()

	usedClasses := []string{}
	unusedClasses := []string{}

	cssClasses.Range(func(key interface{}, value interface{}) bool {
		className := key.(string)

		if classIsUsed(className, &sourceFiles) {
			usedClasses = append(usedClasses, className)
		} else {
			unusedClasses = append(unusedClasses, className)
		}

		return true
	})

	sort.Strings(usedClasses)
	sort.Strings(unusedClasses)

	fmt.Println("Referenced classes:")

	for _, class := range usedClasses {
		file, _ := cssClasses.Load(class)
		fmt.Println(color.GreenString(class), file.(string))
	}

	fmt.Println("\nNot referenced classes (likely outdated):")

	for _, class := range unusedClasses {
		file, _ := cssClasses.Load(class)
		fmt.Println(color.RedString(class), file.(string))
	}
}

func classIsUsed(className string, sourceFiles *sync.Map) bool {
	used := false

	sourceFiles.Range(func(key interface{}, value interface{}) bool {
		source := value.(string)

		if strings.Contains(source, className) {
			used = true
			return false
		}

		return true
	})

	return used
}
