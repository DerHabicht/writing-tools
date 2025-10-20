package papers

type PointList interface {
	ListType() LT

	Level() int

	Points() []Point

	// RenderLaTeX this list (and its points) as a 
	RenderLaTeX() []byte
}

type Point interface {
	// RenderLaTeX renders this point as a LaTeX fragment.
	// This implements the LaTeXRenderer interface.
	RenderLaTeX() []byte

	// Subpoints returns all children nodes as points.
	Subpoints() PointList
}
