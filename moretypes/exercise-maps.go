//go:build OMIT

package main

import (
	"fmt"
	"reflect"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	Test(WordCount)
}

// Test runs a test suite against f.
func Test(f func(string) map[string]int) {
	for _, c := range testCases {
		got := f(c.in)

		// test par itération
		ok := true
		if len(c.want) != len(got) {
			ok = false
		} else {
			for k := range c.want {
				if c.want[k] != got[k] {
					ok = false
				}
			}
		}
		if !ok {
			fmt.Printf("FAIL\n f(%q) =\n  %#v\n want:\n  %#v",
				c.in, got, c.want)
			break
		}

		// OU

		// représentation map string ordonnée
		if fmt.Sprint(c.want) != fmt.Sprint(got) {
			fmt.Printf("FAIL\n f(%q) =\n  %#v\n want:\n  %#v",
				c.in, got, c.want)
			break
		}

		// OU

		// package reflect
		if !reflect.DeepEqual(c.want, got) {
			fmt.Printf("FAIL\n f(%q) =\n  %#v\n want:\n  %#v",
				c.in, got, c.want)
			break
		}

		fmt.Printf("PASS\n f(%q) = \n  %#v\n", c.in, got)
	}
}

var testCases = []struct {
	in   string
	want map[string]int
}{
	{"I am learning Go!", map[string]int{
		"I": 1, "am": 1, "learning": 1, "Go!": 1,
	}},
	{"The quick brown fox jumped over the lazy dog.", map[string]int{
		"The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1,
		"over": 1, "the": 1, "lazy": 1, "dog.": 1,
	}},
	{"I ate a donut. Then I ate another donut.", map[string]int{
		"I": 2, "ate": 2, "a": 1, "donut.": 2, "Then": 1, "another": 1,
	}},
	{"A man a plan a canal panama.", map[string]int{
		"A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama.": 1,
	}},
}
