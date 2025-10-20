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

type LT int

const (
	EnumeratedList LT = iota
	ItemizedList
	UsafEnumeratedList
	UsafItemizedList
	OutlineList
)

func (lt LT) RenderLaTeX() []byte {
	switch lt {
	case EnumeratedList:
		return []byte("enumerate")
	case ItemizedList:
		return []byte("itemize")
	case UsafEnumeratedList:
		return []byte("usafenum")
	case UsafItemizedList:
		return []byte("usafitem")
	case OutlineList:
		return []byte("outline")
	default:
		panic(errors.Errorf("invalid list type value: %d", lt))
	}
}