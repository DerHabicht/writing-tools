package papers

import (
	"bytes"
	"regexp"

	"github.com/ag7if/go-files"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

func ParseFromFile(file files.File) (Paper, error) {
	re := regexp.MustCompile(`(?s)---(.+)---`)

	raw, err := file.ReadFile()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res := re.FindSubmatch(raw)
	if len(res) < 2 {
		return nil, errors.New("file header not found")
	}

	ydocs := bytes.Split(res[1], []byte("\n---"))

	pm := PaperMeta{}
	err = yaml.Unmarshal(ydocs[0], &pm)
	if err != nil {
		return nil, errors.WithMessage(err, "invalid file header")
	}

	var refs map[string]Reference
	if len(ydocs) > 1 {
		refs = make(map[string]Reference)
		err = yaml.Unmarshal(ydocs[1], &refs)
		if err != nil {
			return nil, errors.WithMessage(err, "invalid refs document")
		}
	}

	body := re.ReplaceAll(raw, make([]byte, 0))

	switch pm.PaperType {
	case TalkingPaper:
		panic("not implemented")
	case BulletBackgroundPaper:
		panic("not implemented")
	case OutlinePaper:
		outline := NewOutline(pm, refs)
		err = outline.Parse(body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return outline, nil
	default:
		panic(errors.Errorf("invalid paper type: %v", pm.PaperType))
	}
}
