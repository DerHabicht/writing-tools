package papers

import (
	"strings"

	"github.com/pkg/errors"
)

var Short = false

type OPoint struct {
	level     int
	text      string
	subpoints []*OPoint
}

func NewOPoint(level int, text string) *OPoint {
	return &OPoint{
		level: level,
		text:  text,
	}
}

func (o *OPoint) PTT() PTT {
	return OutlinePoint
}

func (o *OPoint) Level() int {
	return o.level
}

func (o *OPoint) Text() string {
	return o.text
}

func (o *OPoint) AddSubpoints(n ...*OPoint) {
	o.subpoints = append(o.subpoints, n...)
}

func (o *OPoint) Subpoints() []Point {
	var sp []Point

	for _, sub := range o.subpoints {
		sp = append(sp, sub)
	}

	return sp
}

func chompLines(lines []string) []string {
	var chompedLines []string
	for _, line := range lines {
		if !(len(line) == 0 || line[0] == '#') {
			chompedLines = append(chompedLines, line)
		}
	}

	return chompedLines
}

func isShort(line string) (bool, error) {
	line = strings.TrimSpace(line)
	if l2re(true).MatchString(line) {
		return true, nil
	}

	if l0re().MatchString(line) {
		return false, nil
	}

	return false, errors.Errorf("invalid first line of outline: %s", line)
}

func parseLine(lines []string, level int, short bool) (*OPoint, []string, error) {
	var text string
	var subpoints []*OPoint

	for len(lines) > 0 {
		line := strings.TrimRight(lines[0], " \n")

		l := determineLevel(line, short)
		if l == -1 {
			if text == "" {
				return nil, nil, errors.Errorf("invalid first line of outline: %s", line)
			}
			text += " " + strings.TrimSpace(line)
			lines = lines[1:]
		} else if l < level {
			break
		} else if l == level {
			if text != "" {
				break
			}
			text = lvlRe(level, short).ReplaceAllString(line, "")
			lines = lines[1:]
		} else if l == level+1 {
			p, lns, err := parseLine(lines, level+1, short)
			if err != nil {
				return nil, nil, errors.WithStack(err)
			}
			subpoints = append(subpoints, p)
			lines = lns
		} else {
			return nil, nil, errors.Errorf("invalid jump to next outline level: %s (current: %d, next: %d)", line, level, l)
		}
	}

	p := NewOPoint(level, text)
	p.AddSubpoints(subpoints...)

	return p, lines, nil
}

// ParseOutlinePoints takes a set of line as input and parses them into a set of
// points, recursively adding subpoints. This function assumes that YAML docs
// have been stripped out. ParseOutlinePoints will detect if the outline is
// using the short format indicated in CMOS 6.132 (i.e., starting with arabic
// numerals instead of capital roman numerals). Lines where the first non-space
// character is '#' are treated as empty lines (i.e. comments).
//
// Lines are added to the text of the OPoint until a tag indicating the next
// level down is encountered or the end of the line set is encountered. The
// "next level down" is determined using the `level` and `short` params.
func ParseOutlinePoints(lines []string) ([]*OPoint, bool, error) {
	lines = chompLines(lines)

	short, err := isShort(lines[0])
	if err != nil {
		return nil, false, errors.WithStack(err)
	}

	var points []*OPoint

	for len(lines) > 0 {
		point, l, err := parseLine(lines, 0, short)
		if err != nil {
			return nil, false, errors.WithStack(err)
		}
		points = append(points, point)
		lines = l
	}

	return points, short, nil
}
