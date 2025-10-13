package papers

import (
	"github.com/fxtlabs/date"
)

type PaperMeta struct {
	PaperType PT        `yaml:"paper_type"`
	Title     string    `yaml:"title"`
	Author    string    `yaml:"author"`
	Office    string    `yaml:"office"`
	Contact   string    `yaml:"contact"`
	Typist    string    `yaml:"typist"`
	Date      date.Date `yaml:"date"`
	Comment   string    `yaml:"comment"`
}
