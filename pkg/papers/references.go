package papers

import (
	"fmt"
)

type Reference struct {
	RefType string            `yaml:"type"`
	Title   string            `yaml:"title"`
	Author  string            `yaml:"author"`
	Date    string            `yaml:"date"`
	Bibdata map[string]string `yaml:"bibdata"`
	Notes   []Note            `yaml:"notes"`
}

func (r *Reference) RenderBibTeX(key string) string {
	bibtex := fmt.Sprintf(`@%s{%s,%s`, r.RefType, key, "\n")

	bibtex += `    title = {` + r.Title + "},\n"
	bibtex += `    author = {` + r.Author + "},\n"
	bibtex += `    date = {` + r.Date + "},\n"

	for k, v := range r.Bibdata {
		bibtex += fmt.Sprintf(`    %s = {%s},%s`, k, v, "\n")
	}

	bibtex += "}\n"

	return bibtex
}

func (r *Reference) RenderLaTeX(key string) string {
	latex := fmt.Sprintf(`\subsection*{%s~\autocite{%s}}`, r.Title, key) + "\n\n"

	for _, v := range r.Notes {
		latex += v.RenderLaTeX(key)
	}

	return latex
}

type Note struct {
	Citation string   `yaml:"citation"`
	Quote    string   `yaml:"quote"`
	Remarks  []string `yaml:"remarks"`
}

func (n *Note) RenderLaTeX(refKey string) string {
	latex := `\subsubsection*{` + n.Citation + "}\n\n"

	if n.Quote != "" {
		latex += `\begin{quote}` + "\n"
		latex += "    " + n.Quote + "\n"
		latex += `\end{quote}` + "\n\n"
	}

	if len(n.Remarks) > 0 {
		latex += `\begin{itemize}` + "\n"
		for _, remark := range n.Remarks {
			latex += "    " + `\item ` + remark + "\n"
		}
		latex += `\end{itemize}` + "\n\n"
	}

	return latex
}
