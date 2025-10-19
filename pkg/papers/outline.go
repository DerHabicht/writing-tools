package papers

import (
	"github.com/fxtlabs/date"
)

type Outline struct {
	pm      PaperMeta
	refs    []*Reference
	outline *Section
}

func NewOutline(pm PaperMeta, refs []*Reference, points []Point) *Outline {
	var sec *Section
	if refs != nil {
		sec = NewSection(false, "Outline", "", points)
	} else {
		sec = NewSection(false, "", "", points)
	}

	return &Outline{
		pm:      pm,
		refs:    refs,
		outline: sec,
	}
}

func (o *Outline) PT() PT {
	return OutlinePaper
}

func (o *Outline) Title() string {
	return o.pm.Title
}

func (o *Outline) Author() string {
	return o.pm.Author
}

func (o *Outline) Office() string {
	return o.pm.Office
}

func (o *Outline) Contact() string {
	return o.pm.Contact
}

func (o *Outline) Typist() string {
	return o.pm.Typist
}

func (o *Outline) Date() date.Date {
	return o.pm.Date
}

func (o *Outline) References() []*Reference {
	return o.refs
}

func (o *Outline) Sections() []*Section {
	if o.refs != nil {
		panic("implement refs section")
	}

	return []*Section{o.outline}
}
