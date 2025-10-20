package papers

var ListIndent = []byte("    ")

type BPointList struct {
	listType LT
	level int
	points []*BPoint
}

func NewBPointList(listType LT, level int, points ...*BPoint) *BPointList {
	return &BPointList{
		listType: listType,
		level: level,
		points: points,
	}
}

func (bpl *BPointList) ListType() LT {
	return bpl.listType
}

func (bpl *BPointList) Level() int {
	return bpl.level
}

func (bpl *BPointList) Points() []Point {
	var p []Point

	for _, v := range bpl.points {
		p = append(p, v)
	}

	return p
}

func (bpl *BPointList) RenderLaTeX() []byte {
	var indent []byte
	for i := 0; i < bpl.level; i++ {
		indent = append(indent, ListIndent...)
	}

	latex := append(indent, []byte(`\begin{`)...)
	latex = append(latex, bpl.listType.RenderLaTeX()...)
	latex = append(latex, '}', '\n')

	for _, v := range bpl.points {
		p := v.RenderLaTeX()
		latex = append(latex, p...)
		latex = append(latex, '\n')
	}

	latex = append(latex, indent...)
	latex = append(latex, []byte(`\end{`)...)
	latex = append(latex, bpl.listType.RenderLaTeX()...)
	latex = append(latex, '}', '\n')

	return latex
}

type BPoint struct {
	text string
	subpoints *BPointList
}

func NewBPoint(text string, subpoints *BPointList) *BPoint {
	return &BPoint{
		text: text,
		subpoints: subpoints,
	}
}

func (bp *BPoint) RenderLaTeX() []byte {
	latex := append(ListIndent, []byte(`\item `)...)
	latex = append(latex, []byte(bp.text)...)
	return latex
}

func (bp *BPoint) Subpoints() PointList {
	return bp.subpoints
}