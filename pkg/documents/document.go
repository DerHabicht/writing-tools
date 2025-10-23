package documents

type Document interface {
	// 
	Meta() DocMeta

	// LaTeX implements the LaTeXer interface from github.com/ag7if/go-latex
	LaTeX() []byte

	// BibTeX implements the BibTeXer interface from github.com/ag7if/go-latex
	BibTeX() []byte

	// String implements the bulit-in Stringer interface. It's output should be the raw Markdown from which this document was parsed.
	String() string
}
