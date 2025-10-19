package papers

import (
	"testing"

	"github.com/ag7if/go-latex"
	"github.com/yuin/goldmark/parser"
)

func TestPapersImplementsAllInterfaces(t *testing.T) {
	var _ parser.Parser = (Paper)(nil)
	var _ latex.LaTeXer = (Paper)(nil)
	var _ latex.BibTeXer = (Paper)(nil)
}
