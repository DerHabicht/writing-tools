package documents

import (
	"github.com/fxtlabs/date"

)

type DocMeta interface {
	DocType() DocType
	LaTeX() []byte
	BibTeX() []byte
}

type PaperMeta struct {
	Title     string    `yaml:"title"`
	Author    string    `yaml:"author"`
	Office    string    `yaml:"office"`
	Contact   string    `yaml:"contact"`
	Typist    string    `yaml:"typist"`
	Date      date.Date `yaml:"date"`
	Comment   string    `yaml:"comment"`
}