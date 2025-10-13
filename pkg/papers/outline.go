package papers

import (
	"fmt"
	"strings"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"
)

type Outline struct {
	meta       PaperMeta
	short      bool
	references map[string]Reference
	points     []*OPoint
}

func NewOutline(meta PaperMeta, refs map[string]Reference) *Outline {
	return &Outline{
		meta:       meta,
		references: refs,
	}
}

func NewShortOutline(meta PaperMeta, refs map[string]Reference) *Outline {
	return &Outline{
		meta:       meta,
		short:      true,
		references: refs,
	}
}

func (o *Outline) AddReference(k string, r Reference) {
	if o.references == nil {
		o.references = make(map[string]Reference)
	}
	
	o.references[k] = r
}

func (o *Outline) AddOPoints(p ...*OPoint) {
	o.points = append(o.points, p...)
}

func (o *Outline) PT() PT {
	return OutlinePaper
}

func (o *Outline) Title() string {
	return o.meta.Title
}

func (o *Outline) Author() string {
	return o.meta.Author
}

func (o *Outline) Office() string {
	return o.meta.Office
}

func (o *Outline) Contact() string {
	return o.meta.Contact
}

func (o *Outline) Typist() string {
	return o.meta.Typist
}

func (o *Outline) Date() date.Date {
	return o.meta.Date
}

func (o *Outline) Short() bool {
	return o.short
}

func (o *Outline) Parse(b []byte) error {
	raw := string(b)
	lines := strings.Split(raw, "\n")

	points, short, err := ParseOutlinePoints(lines)
	if err != nil {
		return errors.WithStack(err)
	}

	o.short = short
	o.points = points

	return nil
}

func (o *Outline) Points() []Point {
	var p []Point

	for _, point := range o.points {
		p = append(p, point)
	}

	return p
}

func latexOutlineLists(points []Point) string {
	listIndent := strings.Repeat(" ", 4*points[0].Level())
	itemIndent := strings.Repeat(" ", 4*(points[0].Level()+1))

	latex := listIndent + `\begin{outline}` + "\n"
	for _, point := range points {
		latex += itemIndent + `\item ` + point.RenderLaTeX() + "\n"

		if len(point.Subpoints()) > 0 {
			latex += latexOutlineLists(point.Subpoints())
		}
	}

	latex += listIndent + `\end{outline}` + "\n"

	return latex
}

func (o *Outline) BibTeX() string {
	if o.references == nil {
		return ""
	}

	bibtex := ""
	for k, v := range o.references {
		bibtex += v.RenderBibTeX(k)
	}

	return bibtex
}

func (o *Outline) LaTeX() string {
	latex := `\documentclass[outline%s]{usafpaper}
%s
\title{%s}
\date{%s}
\author{%s}
\authorOffice{%s}
\authorPhone{%s}
\typist{%s}

\begin{document}
\maketitle

%s
\end{document}
`
	short := ""
	usebibtex := ""
	body := ""

	if o.short {
		short = ",short"
	}

	if o.references != nil {
		usebibtex = `
\usepackage[authordate,backend=biber]{biblatex-chicago}
\addbibresource{references.bib}
`
		body = `\section*{Outline}` + "\n\n"
		body += latexOutlineLists(o.Points())
		body += "\n" + `\pagebreak` + "\n"
		body += "\n" + `\section*{Notes}` + "\n\n"
		for k, v := range o.references {
			body += v.RenderLaTeX(k)
		}

		body += `\pagebreak` + "\n\n"

		body += `\printbibliography` + "\n"

	} else {
		body = latexOutlineLists(o.Points())
	}

	latex = fmt.Sprintf(
		latex,
		short,
		usebibtex,
		o.meta.Title,
		o.meta.Date,
		o.meta.Author,
		o.meta.Office,
		o.meta.Contact,
		o.meta.Typist,
		body,
	)

	return latex
}
