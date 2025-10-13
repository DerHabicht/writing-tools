package papers

import (
	"github.com/fxtlabs/date"
)

type Paper interface {
	PT() PT
	Title() string
	Author() string
	Office() string
	Contact() string
	Typist() string
	Date() date.Date
	Parse(b []byte) error
	Points() []Point
	LaTeX() string
	BibTeX() string
}
