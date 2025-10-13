package papers

import (
	"strings"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"
)

type BulletPaper struct {
	meta       PaperMeta
	references map[string]Reference
	points     []*BPoint
}

func NewBulletPaper(meta PaperMeta, refs map[string]Reference) *BulletPaper {
	return &BulletPaper{
		meta:       meta,
		references: refs,
	}
}

func (bp *BulletPaper) AddReference(k string, r Reference) {
	if bp.references == nil {
		bp.references = make(map[string]Reference)
	}

	bp.references[k] = r
}

func (bp *BulletPaper) AddBPoints(p ...*BPoint) {
	bp.points = append(bp.points, p...)
}

func (bp *BulletPaper) PT() PT {
	return bp.meta.PaperType
}

func (bp *BulletPaper) Title() string {
	return bp.meta.Title
}

func (bp *BulletPaper) Author() string {
	return bp.meta.Author
}

func (bp *BulletPaper) Office() string {
	return bp.meta.Office
}

func (bp *BulletPaper) Contact() string {
	return bp.meta.Contact
}

func (bp *BulletPaper) Typist() string {
	return bp.meta.Typist
}

func (bp *BulletPaper) Date() date.Date {
	return bp.meta.Date
}

func (bp *BulletPaper) Parse(b []byte) error {
	raw := string(b)
	lines := strings.Split(raw, "\n")

	points, err := ParseBulletPoints(lines)
	if err != nil {
		return errors.WithStack(err)
	}

	bp.points = points

	return err
}

func (bp *BulletPaper) Points() []Point {
	var p []Point

	for _, point := range bp.points {
		p = append(p, point)
	}

	return p
}

func (bp *BulletPaper) BibTeX() string {
	if bp.references == nil {
		return ""
	}

	bibtex := ""
	for k, v := range bp.references {
		bibtex += v.RenderBibTeX(k)
	}

	return bibtex
}

func (bp *BulletPaper) LaTeX() string {
	panic("implement me")
}
