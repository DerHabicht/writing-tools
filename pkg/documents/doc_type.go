package documents

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type DocType int

const (
	DocBook DocType = iota
	DocMemo
	DocMFR
	DocOutline
	DocPointPaper
	DocTalkingPaper
	DocBulletBackgroundPaper
	DocBackgroundPaper
	DocPositionPaper
	DocSTRATPLAN
	DocCAMPLAN
	DocBPLAN
	DocCONPLAN
	DocOPLAN
	DocOPORD
	DocFRAGORD
	DocSPLAN
	DocAAR
)

func ParseDocType(s string) (DocType, error) {
	switch strings.ToLower(s) {
	case "book":
		return DocBook, nil
	case "memo":
		return DocMemo, nil
	case "mfr":
		return DocMFR, nil
	case "outline":
		return DocOutline, nil
	case "point-paper":
		return DocPointPaper, nil
	case "talking-paper":
		return DocTalkingPaper, nil
	case "bullet-background-paper":
		return DocBulletBackgroundPaper, nil
	case "background-paper":
		return DocBackgroundPaper, nil
	case "position-paper":
		return DocPositionPaper, nil
	case "stratplan":
		return DocSTRATPLAN, nil
	case "camplan":
		return DocCAMPLAN, nil
	case "bplan":
		return DocBPLAN, nil
	case "conplan":
		return DocCONPLAN, nil
	case "oplan":
		return DocOPLAN, nil
	case "opord":
		return DocOPORD, nil
	case "fragord":
		return DocFRAGORD, nil
	case "splan":
		return DocSPLAN, nil
	case "aar":
		return DocAAR, nil
	default:
		return -1, errors.Errorf("unrecognized document type: %s", s)
	}
}

func (dt DocType) String() string {
	switch dt {
	case DocBook:
		return "book"
	case DocMemo:
		return "memo"
	case DocMFR:
		return "mfr"
	case DocOutline:
		return "outline"
	case DocPointPaper:
		return "point-paper"
	case DocTalkingPaper:
		return "talking-paper"
	case DocBulletBackgroundPaper:
		return "bullet-background-paper"
	case DocBackgroundPaper:
		return "background-paper"
	case DocPositionPaper:
		return "position-paper"
	case DocSTRATPLAN:
		return "stratplan"
	case DocCAMPLAN:
		return "camplan"
	case DocBPLAN:
		return "bplan"
	case DocCONPLAN:
		return "conplan"
	case DocOPLAN:
		return "oplan"
	case DocOPORD:
		return "opord"
	case DocFRAGORD:
		return "fragord"
	case DocSPLAN:
		return "splan"
	case DocAAR:
		return "aar"
	default:
		panic(errors.Errorf("invalid DocType value: %d", dt))
	}
}

func (dt DocType) MarshalJSON() ([]byte, error) {
	out := fmt.Sprintf(`"%s"`, dt.String())
	return []byte(out), nil
}

func (dt *DocType) UnmarshalJSON(raw []byte) error {
	str := strings.Trim(string(raw), "\"")

	val, err := ParseDocType(str)
	if err != nil {
		return errors.WithStack(err)
	}

	*dt = val
	return nil
}

func (dt DocType) MarshalYAML() (interface{}, error) {
	return dt.String(), nil
}

func (dt *DocType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw string

	err := unmarshal(&raw)
	if err != nil {
		return errors.WithStack(err)
	}

	parsed, err := ParseDocType(raw)
	if err != nil {
		return errors.WithStack(err)
	}

	*dt = parsed
	return nil
}

func (dt DocType) Value() (driver.Value, error) {
	return []byte(dt.String()), nil
}

func (dt *DocType) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseDocType(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*dt = val
	return nil
}
