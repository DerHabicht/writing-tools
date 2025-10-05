package papers

type Point interface {
	PTT() PTT
	Level() int
	Text() string
	Subpoints() []Point
}
