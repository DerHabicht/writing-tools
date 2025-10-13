package papers

import (
	"strings"

	"github.com/pkg/errors"
)

type PT int

const (
	PointPaper PT = iota
	TalkingPaper
	BulletBackgroundPaper
	OutlinePaper
)

func ParsePT(s string) (PT, error) {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "point":
		return PointPaper, nil
	case "talking":
		return TalkingPaper, nil
	case "bullet-background":
		return BulletBackgroundPaper, nil
	case "outline":
		return OutlinePaper, nil
	default:
		return -1, errors.Errorf("unknown paper type: %s", s)
	}
}

func (p PT) String() string {
	switch p {
	case PointPaper:
		return "point"
	case TalkingPaper:
		return "talking"
	case BulletBackgroundPaper:
		return "bullet-background"
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
