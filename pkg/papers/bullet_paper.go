package papers

import (
	"github.com/fxtlabs/date"
)

type BulletPaper struct {
	meta PaperMeta
}

func (b *BulletPaper) PT() PT {
	return b.meta.PaperType
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

func (b *BulletPaper) Typist() string {
	return b.meta.Typist
}

func (b *BulletPaper) Date() date.Date {
	return b.meta.Date
}

func (b *BulletPaper) Parse(raw []byte) error {
	//TODO implement me
	panic("implement me")
}

func (b *BulletPaper) Points() []Point {
	//TODO implement me
	panic("implement me")
}

func (b *BulletPaper) LaTeX() string {
	//TODO implement me
	panic("implement me")
}
