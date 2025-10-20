package papers

import (
	"github.com/fxtlabs/date"
)

type BulletPaper struct {
	paperType PT
	meta      PaperMeta
	refs      []*Reference
	sections  []*Section
}

func NewBulletPaper(pt PT, meta PaperMeta, refs []*Reference, sections []*Section) *BulletPaper {
	return &BulletPaper{
		paperType: pt,
		meta:      meta,
		refs:      refs,
		sections:  sections,
	}
}

func (b *BulletPaper) PT() PT {
	return b.paperType
}

func (b *BulletPaper) Title() string {
	return b.meta.Title
}

func (b *BulletPaper) Author() string {
	return b.meta.Author
}

func (b *BulletPaper) Office() string {
	return b.meta.Office
}

func (b *BulletPaper) Contact() string {
	return b.meta.Contact
}

func (b *BulletPaper) Typist() string {
	return b.meta.Typist
}

func (b *BulletPaper) Date() date.Date {
	return b.meta.Date
}

func (b *BulletPaper) References() []*Reference {
	return b.refs
}

func (b *BulletPaper) Sections() []*Section {
	return b.sections
}
