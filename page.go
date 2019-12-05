package crawler

import (
	"fmt"
	"io"
)

const defaultLinksDepth = 10

// Page represents page struct, it has info about page url and map of links on this page
type Page struct {
	Link       string
	CrossPages []*Page
}

// Print prints hierarchical structure of page and related links
func (p *Page) Print(out io.Writer, depth int) {
	if depth == 0 {
		depth = defaultLinksDepth
	}
	p.print(out, depth, 1)
}

func (p *Page) print(out io.Writer, depth, level int) {
	if level > depth {
		return
	}
	var prefix string
	i := 0
	for i < level {
		prefix += "\t"
		i++
	}
	fmt.Fprint(out, prefix+p.Link+"\n")
	level++
	for _, cp := range p.CrossPages {
		cp.print(out, depth, level)
	}
	return
}
