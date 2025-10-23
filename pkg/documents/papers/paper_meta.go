package documents

import (
	"strings"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/writing-tools/pkg/references"
)

const (
	PaperOffice     string = "office"
	PaperContact    string = "contact"
	PaperTypist     string = "typist"
	PaperReferences string = "references"
)

type PaperMeta struct {
	PaperType   DocType                         `yaml:"type"`
	PaperTitle  string                          `yaml:"title"`
	PaperAuthor string                          `yaml:"author"`
	PaperDate   date.Date                       `yaml:"date"`
	Office      string                          `yaml:"office"`
	Contact     string                          `yaml:"contact"`
	Typist      string                          `yaml:"typist"`
	References  map[string]references.Reference `yaml:"references"`
}

func (pm PaperMeta) DocType() DocType {
	return pm.PaperType
}

func (pm PaperMeta) Title() string {
	return pm.PaperTitle
}

func (pm PaperMeta) Author() string {
	return pm.PaperAuthor
}

func (pm PaperMeta) Date() date.Date {
	return pm.PaperDate
}

func (pm PaperMeta) Data(key string) any {
	switch strings.ToLower(key) {
	case PaperOffice:
		return pm.Office
	case PaperContact:
		return pm.Contact
	case PaperTypist:
		return pm.Typist
	case PaperReferences:
		return pm.References
	default:
		return nil
	}
}
