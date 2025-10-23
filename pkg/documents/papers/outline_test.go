package papers

/*

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/stretchr/testify/assert"

	ipapers "github.com/derhabicht/writing-tools/internal/papers"
)

const expectedLongOutlineLaTeXOutput = `\documentclass[outline]{usafpaper}

\title{Sample Outline}
\date{2017-01-01}
\author{University of Chicago}
\authorOffice{UC/CMOS}
\authorPhone{(555) 555-1234}
\typist{cmos}

\begin{document}
\maketitle

\begin{outline}
    \item Historical Introduction
    \item Dentition in various groups of vertebrates
    \begin{outline}
        \item Reptilia
        \begin{outline}
            \item Histology and development of reptilian teeth
            \item Survey of forms
        \end{outline}
        \item Mammalia
        \begin{outline}
            \item Histology and development of mammalian teeth
            \item Survey of forms
            \begin{outline}
                \item Primates
                \begin{outline}
                    \item Lemuroidea
                    \item Anthropoidea
                    \begin{outline}
                        \item Platyrrhini
                        \item Catarrhini
                        \begin{outline}
                            \item Cercopithecidae
                            \item Pongidea
                        \end{outline}
                    \end{outline}
                \end{outline}
                \item Carnivora
                \begin{outline}
                    \item Creodonta
                    \item Fissipedia
                    \begin{outline}
                        \item Ailuroidea
                        \item Arctoidea
                    \end{outline}
                    \item Pinnipedia
                \end{outline}
                \item Etc.
            \end{outline}
        \end{outline}
    \end{outline}
\end{outline}

\end{document}
`

const expectedShortOutlineLaTeXOutput = `\documentclass[outline,short]{usafpaper}

\usepackage[authordate,backend=biber]{biblatex-chicago}
\addbibresource{references.bib}

\title{Sample Short Outline}
\date{2017-01-01}
\author{University of Chicago}
\authorOffice{UC/CMOS}
\authorPhone{(555) 555-1234}
\typist{cmos}

\begin{document}
\maketitle

\section*{Outline}

\begin{outline}
    \item Punctuation
    \begin{outline}
        \item Using commas appropriately
        \item Deleting unnecessary quotation marks
        \item Distinguishing colons from semicolons
    \end{outline}
    \item Spelling
    \begin{outline}
        \item Using a dictionary appropriately
        \item Recognizing homonyms
        \item Hyphenating correctly
    \end{outline}
    \item Syntax
    \begin{outline}
        \item Matching verb to subject
        \item Recognizing and eliminating misplaced modifiers
        \item Distinguishing phrases from clauses while singing the "Conjunction Junction" song
    \end{outline}
\end{outline}

\pagebreak

\section*{Notes}

\subsection*{The Chicago Manual of Style~\autocite{cmos}}

\subsubsection*{p. 416}

\begin{quote}
    In a list with fewer levels, one might dispense with the capital roman numerals and capital letters and instead begin with arabic numerals. What is important is that readers see at a glance the level to which each item belongs. Note that each division and subdivision should normally contain at least two items.
\end{quote}

\begin{itemize}
    \item This means ` + "``" + `short'' outlines should be supported
    \item Short papers start with arabic numerals instead of capital roman numerals
\end{itemize}

\subsection*{Starship Troopers~\autocite{starship_troopers}}

\subsubsection*{ch. 2}

\begin{quote}
    Anyone who clings to the historically untrue---and thoroughly immoral---doctrine that ` + "`" + `violence never settles anything' I would advise to conjure up the ghosts of Napoleon Bonaparte and of the Duke of Wellington and let them debate it. The ghost of Hitler could referee, and the jury might well be the Dodo, the Great Auk, and the Passenger Pigeon. Violence, naked force, has settled more issues in history than has any other factor, and the contrary opinion is wishful thinking at its worst. Breeds that forget this basic truth have always paid for it with their lives and freedoms.
\end{quote}

\subsubsection*{ch. 12}

\begin{quote}
    Morals---all correct moral laws---derive from the instinct to survive. Moral behavior is survival behavior above the individual level.
\end{quote}

\pagebreak

\printbibliography

\end{document}
`

const expectedOutlineBibTeXOutput = `
@manual{cmos,
    title = {The Chicago Manual of Style},
    author = {University of Chicago},
    date = {2017},
	address = {Chicago, IL},
	organization = {University of Chicago},
	publisher = {University of Chicago Press},
}

@book{starship_troopers,
	title = {Starship Troopers},
	author = {Robert A. Heinlein},
	date = {1959},
	publisher = {G.P. Putnam and Sons},
}
`

func generateExpectedLongOutline() *ppapers.BulletPaper {
	panic("rewrite me")
}

func generateExpectedShortOutline() *ppapers.BulletPaper {
	panic("rewrite me")
}

func TestSampleOutlinesCompile(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	so := latex.LaTeXString(expectedShortOutlineLaTeXOutput)
	lo := latex.LaTeXString(expectedLongOutlineLaTeXOutput)

	cwd, err := os.Getwd()
	assert.NoError(t, err)
	projDir := filepath.Join(cwd, "..", "..")

	clsDir := files.NewDirectory(filepath.Join(projDir, "assets", "ag7if-tex", "cls"))
	styDir := files.NewDirectory(filepath.Join(projDir, "assets", "ag7if-tex", "sty"))

	usafpaper, err := clsDir.NewFile("usafpaper.cls")
	assert.NoError(t, err)

	usafpub, err := styDir.NewFile("usafpub.sty")
	assert.NoError(t, err)

	assets := []files.File{usafpaper, usafpub}

	cacheDir := files.NewDirectory(filepath.Join(projDir, "test", "cache"))
	err = cacheDir.Create()
	assert.NoError(t, err)
	dataDir := files.NewDirectory(filepath.Join(projDir, "test", "data"))

	soOut, err := dataDir.NewFile("sample_short_outline.pdf")
	assert.NoError(t, err)

	loOut, err := dataDir.NewFile("sample_outline.pdf")
	assert.NoError(t, err)

	c := latex.NewCompiler(latex.XeLaTeX, latex.Biber, *cacheDir)
	ref, err := cacheDir.CreateFile("references.bib")
	assert.NoError(t, err)
	err = ref.WriteString(expectedOutlineBibTeXOutput)
	assert.NoError(t, err)
	err = c.GenerateLaTeX(so, soOut, assets)
	assert.NoError(t, err)
	err = c.CompileLaTeX(soOut)
	assert.NoError(t, err)
	assert.FileExists(t, soOut.FullPath())
	err = soOut.Remove()
	assert.NoError(t, err)

	c = latex.NewCompiler(latex.XeLaTeX, latex.NoBib, *cacheDir)
	err = c.GenerateLaTeX(lo, loOut, assets)
	assert.NoError(t, err)
	err = c.CompileLaTeX(loOut)
	assert.NoError(t, err)
	assert.FileExists(t, loOut.FullPath())
	err = loOut.Remove()
	assert.NoError(t, err)
}

func TestOutlineLaTeX(t *testing.T) {
	t.Skip()
	fLongIn, err := files.NewFile("../../test/data/sample_outline.pmd")
	assert.NoError(t, err)

	long, err := ipapers.ParseFromFile(fLongIn)
	assert.NoError(t, err)

	longTeX := long.LaTeX()
	assert.Equal(t, expectedLongOutlineLaTeXOutput, longTeX)

	fShortIn, err := files.NewFile("../../test/data/sample_short_outline.pmd")
	assert.NoError(t, err)

	short, err := ipapers.ParseFromFile(fShortIn)
	assert.NoError(t, err)

	shortTeX := short.LaTeX()
	assert.Equal(t, expectedShortOutlineLaTeXOutput, shortTeX)
}

func TestBibliography(t *testing.T) {
	fShortIn, err := files.NewFile("../../test/data/sample_short_outline.pmd")
	assert.NoError(t, err)

	short, err := ipapers.ParseFromFile(fShortIn)
	assert.NoError(t, err)

	shortO, ok := short.(*Outline)
	assert.True(t, ok)

	bibtex := shortO.BibTeX()

	assert.Regexp(t, `@manual\{cmos,`, bibtex)
	assert.Regexp(t, `    title = \{The Chicago Manual of Style\},`, bibtex)
	assert.Regexp(t, `    author = \{University of Chicago\},`, bibtex)
	assert.Regexp(t, `    date = \{2017\},`, bibtex)
	assert.Regexp(t, `    address = \{Chicago, IL\},`, bibtex)
	assert.Regexp(t, `    organization = \{University of Chicago\},`, bibtex)
	assert.Regexp(t, `    publisher = \{University of Chicago Press\},`, bibtex)

	assert.Regexp(t, `@book\{starship_troopers,`, bibtex)
	assert.Regexp(t, `    title = \{Starship Troopers\},`, bibtex)
	assert.Regexp(t, `    author = \{Robert A. Heinlein\},`, bibtex)
	assert.Regexp(t, `    date = \{1959\},`, bibtex)
	assert.Regexp(t, `    publisher = \{G. P. Putnam's Sons\},`, bibtex)
}
*/
