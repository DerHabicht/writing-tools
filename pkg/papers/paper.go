package papers

import (
	"github.com/fxtlabs/date"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Paper interface {
	// PT returns the paper type
	PT() PT
	// Title returns the title of this paper
	Title() string
	// Author returns the paper's author
	Author() string
	// Office returns the office symbol of the paper's author
	Office() string
	// Contact returns the contact info (phone/email address) of the paper's author
	Contact() string
	// Typist returns the initals of the paper's typist
	Typist() string
	// Date returns the date of the paper
	Date() date.Date
	// Points returns the papers the tree of points in the paper
	Points() []Point
	// Parse implements Goldmark's parse.Parser interface
	Parse(reader text.Reader, opts ...parser.ParseOption) ast.Node
	// AddOptions implements Goldmark's parse.Parser interface
	AddOptions(...parser.Option)
	// LaTeX implements the LaTeXer interface
	LaTeX() []byte
	// BibTeX implements the BibTeXer interface
	BibTeX() []byte
}
