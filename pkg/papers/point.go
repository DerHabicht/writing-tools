package papers

type Point interface {
	// PTT returns the point type (bullet or outline point).
	PTT() PTT
	
	// Level returns the nesting level of this point in the tree.
	Level() int
	
	// RenderLaTeX renders this point as a LaTeX fragment.
	// This implements the LaTeXRenderer interface.
	RenderLaTeX() []byte

	// Subpoints returns all children nodes as points.
	Subpoints() []Point
}
