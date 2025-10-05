package papers

import (
	"strings"

	"github.com/pkg/errors"
)

type PT int

const (
	TalkingPaper PT = iota
	BulletBackgroundPaper
	OutlinePaper
)

func ParsePT(s string) (PT, error) {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "talking":
		return TalkingPaper, nil
	case "bullet":
		return BulletBackgroundPaper, nil
	case "outline":
		return OutlinePaper, nil
	default:
		return -1, errors.Errorf("unknown paper type: %s", s)
	}
}

func (p PT) String() string {
	switch p {
	case TalkingPaper:
		return "talking"
	case BulletBackgroundPaper:
		return "bullet"
	case OutlinePaper:
		return "outline"
	default:
		panic(errors.Errorf("invalid paper type: %d", p))
	}
}

func (p PT) MarshalYAML() (interface{}, error) {
	return p.String(), nil
}

func (p *PT) UnmarshalYAML(b []byte) error {
	pt, err := ParsePT(string(b))
	if err != nil {
		return errors.WithStack(err)
	}

	*p = pt
	return nil
}

type PTT int

const (
	BulletPoint PTT = iota
	OutlinePoint
)
