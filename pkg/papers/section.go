package papers

type Section struct {
	numbered  bool
	title     string
	paragraph string
	points    []Point
}

func NewSection(numbered bool, title string, paragraph string, points []Point) *Section {
	return &Section{
		numbered:  numbered,
		title:     title,
		paragraph: paragraph,
		points:    points,
	}
}

func (s *Section) Numbered() bool {
	return s.numbered
}

func (s *Section) Title() string {
	return s.title
}

func (s *Section) Paragraph() string {
	return s.paragraph
}

func (s *Section) Points() []Point {
	return s.points
}
