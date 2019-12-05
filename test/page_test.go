package test

import (
	"bytes"
	"testing"

	"github.com/sbashilov/crawler"
)

func TestPage_Print(t *testing.T) {
	expectedPrint := `	https://gobyexample.com/
		https://gobyexample.com/range
			https://gobyexample.com/functions
`
	p := &crawler.Page{
		Link: "https://gobyexample.com/",
		CrossPages: []*crawler.Page{
			&crawler.Page{
				Link: "https://gobyexample.com/range",
				CrossPages: []*crawler.Page{
					&crawler.Page{Link: "https://gobyexample.com/functions"},
				},
			},
		},
	}
	out := new(bytes.Buffer)
	p.Print(out, 0)
	result := out.String()
	if result != expectedPrint {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, expectedPrint)
	}
}
