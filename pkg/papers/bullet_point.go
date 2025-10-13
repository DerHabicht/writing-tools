package papers

type BPoint struct {
	level     int
	text      string
	subpoints []*BPoint
}

func NewBPoint(level int, text string) *BPoint {
	return &BPoint{
		level: level,
		text:  text,
	}
}

func (bp *BPoint) AddSubpoints(sp ...*BPoint) {
	bp.subpoints = append(bp.subpoints, sp...)
}

func (bp *BPoint) PTT() PTT {
	return BulletPoint
}

func (bp *BPoint) Level() int {
	panic("implement me")
}

func (bp *BPoint) Text() string {
	panic("implement me")
}

func (bp *BPoint) Subpoints() []Point {
	var sp []Point

	for _, sub := range bp.subpoints {
		sp = append(sp, sub)
	}

	return sp
}

func (bp *BPoint) RenderLaTeX() string {
	panic("implement me")
}

func ParseBulletPoints(lines []string) ([]*BPoint, error) {
	
}
