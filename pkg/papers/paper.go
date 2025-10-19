package papers

import (
	"github.com/fxtlabs/date"
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

	// References returns the refs and notes attached to this paper
	References() []*Reference

	// Points returns the papers the tree of points in the paper
	Sections() []*Section
}
