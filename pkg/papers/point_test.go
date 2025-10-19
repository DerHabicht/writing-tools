package papers

import (
	"testing"

	"github.com/ag7if/go-latex"
	"github.com/yuin/goldmark/ast"
)

func TestPointsImplementsAllInterfaces(t *testing.T) {
	var _ ast.Node = (Point)(nil)
	var _ latex.LaTeXRenderer = (Point)(nil)
}
