package papers

import (
	"testing"

	"github.com/ag7if/go-files"
	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

const expectedLongLaTeXOutput = `\documentclass[outline]{usafpaper}

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

const expectedShortLaTeXOutput = `\documentclass[outline,short]{usafpaper}

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
    \item This means "short" outlines should be supported
    \item Short papers start with arabic numerals instead of capital roman numerals
\end{itemize}

\subsection*{Starship Troopers~\autocite{starship_troopers}}

\subsubsection*{ch. 2}

\begin{quote}
    Anyone who clings to the historically untrue---and thoroughly immoral---doctrine that 'violence never settles anything' I would advise to conjure up the ghosts of Napoleon Bonaparte and of the Duke of Wellington and let them debate it. The ghost of Hitler could referee, and the jury might well be the Dodo, the Great Auk, and the Passenger Pigeon. Violence, naked force, has settled more issues in history than has any other factor, and the contrary opinion is wishful thinking at its worst. Breeds that forget this basic truth have always paid for it with their lives and freedoms.
\end{quote}

\subsubsection*{ch. 12}

\begin{quote}
    Morals---all correct moral laws---derive from the instinct to survive. Moral behavior is survival behavior above the individual level.
\end{quote}

\pagebreak

\printbibliography

\end{document}
`

const expectedBibTeXOutput = `
@manual{cmos,
    title = {The Chicago Manual of Style},
    author = {University of Chicago},
    date = {2017},
	address = {Chicago, IL},
	organization = {University of Chicago},
	publisher = {University of Chicago Press},
}

@book(starship_troopers,
	title = {Starship Troopers},
	author = {Robert A. Heinlein},
	date = {1959},
	publisher = {G.P. Putnam and Sons},
`

func generateExpectedLongOutline() *Outline {
	pm := PaperMeta{
		PaperType: OutlinePaper,
		Title:     "Sample Outline",
		Author:    "University of Chicago",
		Office:    "UC/CMOS",
		Contact:   "(555) 555-1234",
		Typist:    "cmos",
		Date:      date.New(2017, 1, 1),
		Comment:   "Sample outline found on p. 415 of CMOS 17.",
	}

	i := NewOPoint(0, "Historical Introduction")
	ii := NewOPoint(0, "Dentition in various groups of vertebrates")
	iiA := NewOPoint(1, "Reptilia")
	iiA1 := NewOPoint(2, "Histology and development of reptilian teeth")
	iiA2 := NewOPoint(2, "Survey of forms")
	iiB := NewOPoint(1, "Mammalia")
	iiB1 := NewOPoint(2, "Histology and development of mammalian teeth")
	iiB2 := NewOPoint(2, "Survey of forms")
	iiB2a := NewOPoint(3, "Primates")
	iiB2a1 := NewOPoint(4, "Lemuroidea")
	iiB2a2 := NewOPoint(4, "Anthropoidea")
	iiB2a2a := NewOPoint(5, "Platyrrhini")
	iiB2a2b := NewOPoint(5, "Catarrhini")
	iiB2a2b_i := NewOPoint(6, "Cercopithecidae")
	iiB2a2b_ii := NewOPoint(6, "Pongidea")
	iiB2b := NewOPoint(3, "Carnivora")
	iiB2b1 := NewOPoint(4, "Creodonta")
	iiB2b2 := NewOPoint(4, "Fissipedia")
	iiB2b2a := NewOPoint(5, "Ailuroidea")
	iiB2b2b := NewOPoint(5, "Arctoidea")
	iiB2b3 := NewOPoint(4, "Pinnipedia")
	iiB2c := NewOPoint(3, "Etc.")

	ii.AddSubpoints(iiA, iiB)
	iiA.AddSubpoints(iiA1, iiA2)
	iiB.AddSubpoints(iiB1, iiB2)
	iiB2.AddSubpoints(iiB2a, iiB2b, iiB2c)
	iiB2a.AddSubpoints(iiB2a1, iiB2a2)
	iiB2a2.AddSubpoints(iiB2a2a, iiB2a2b)
	iiB2a2b.AddSubpoints(iiB2a2b_i, iiB2a2b_ii)
	iiB2b.AddSubpoints(iiB2b1, iiB2b2, iiB2b3)
	iiB2b2.AddSubpoints(iiB2b2a, iiB2b2b)

	outline := NewOutline(pm, nil)
	outline.AddOPoints(i, ii)

	return outline
}

func generateExpectedShortOutline() *Outline {
	pm := PaperMeta{
		PaperType: OutlinePaper,
		Title:     "Sample Short Outline",
		Author:    "University of Chicago",
		Office:    "UC/CMOS",
		Contact:   "(555) 555-1234",
		Typist:    "cmos",
		Date:      date.New(2017, 1, 1),
		Comment:   "Sample short outline found on p. 415 of CMOS 17.",
	}

	refs := map[string]Reference{
		"cmos": {
			RefType: "manual",
			Title:   "The Chicago Manual of Style",
			Author:  "University of Chicago",
			Date:    "2017",
			Bibdata: map[string]string{
				"organization": "University of Chicago",
				"publisher":    "University of Chicago Press",
				"address":      "Chicago, IL",
			},
			Notes: []Note{
				{
					Citation: "p. 416",
					Quote: "In a list with fewer levels, " +
						"one might dispense with the capital roman numerals and capital letters and instead begin with arabic numerals. What is important is that readers see at a glance the level to which each item belongs. Note that each division and subdivision should normally contain at least two items.",
					Remarks: []string{
						"This means \"short\" outlines should be supported",
						"Short papers start with arabic numerals instead of capital roman numerals",
					},
				},
			},
		},
		"starship_troopers": {
			RefType: "book",
			Title:   "Starship Troopers",
			Author:  "Robert A. Heinlein",
			Date:    "1959",
			Bibdata: map[string]string{
				"publisher": "G. P. Putnam's Sons",
			},
			Notes: []Note{
				{
					Citation: "ch. 2",
					Quote:    "Anyone who clings to the historically untrue---and thoroughly immoral---doctrine that 'violence never settles anything' I would advise to conjure up the ghosts of Napoleon Bonaparte and of the Duke of Wellington and let them debate it. The ghost of Hitler could referee, and the jury might well be the Dodo, the Great Auk, and the Passenger Pigeon. Violence, naked force, has settled more issues in history than has any other factor, and the contrary opinion is wishful thinking at its worst. Breeds that forget this basic truth have always paid for it with their lives and freedoms.",
				},
				{
					Citation: "ch. 12",
					Quote:    "Morals---all correct moral laws---derive from the instinct to survive. Moral behavior is survival behavior above the individual level.",
				},
			},
		},
	}

	one := NewOPoint(0, "Punctuation")
	one_a := NewOPoint(1, "Using commas appropriately")
	one_b := NewOPoint(1, "Deleting unnecessary quotation marks")
	one_c := NewOPoint(1, "Distinguishing colons from semicolons")
	one.AddSubpoints(one_a, one_b, one_c)

	two := NewOPoint(0, "Spelling")
	two_a := NewOPoint(1, "Using a dictionary appropriately")
	two_b := NewOPoint(1, "Recognizing homonyms")
	two_c := NewOPoint(1, "Hyphenating correctly")
	two.AddSubpoints(two_a, two_b, two_c)

	three := NewOPoint(0, "Syntax")
	three_a := NewOPoint(1, "Matching verb to subject")
	three_b := NewOPoint(1, "Recognizing and eliminating misplaced modifiers")
	three_c := NewOPoint(1, "Distinguishing phrases from clauses while singing the \"Conjunction Junction\" song")
	three.AddSubpoints(three_a, three_b, three_c)

	outline := NewShortOutline(pm, refs)
	outline.AddOPoints(one, two, three)

	return outline
}

func TestParseOutline_Long(t *testing.T) {
	expected := generateExpectedLongOutline()

	f, err := files.NewFile("../../test/data/sample_outline.pmd")
	assert.NoError(t, err)

	outline, err := ParseFromFile(f)
	assert.NoError(t, err)

	assert.Equal(t, expected, outline)
}

func TestParseOutline_Short(t *testing.T) {
	expected := generateExpectedShortOutline()

	f, err := files.NewFile("../../test/data/sample_short_outline.pmd")
	assert.NoError(t, err)

	outline, err := ParseFromFile(f)
	assert.NoError(t, err)

	assert.Equal(t, expected, outline)
}

func TestOutlineLaTeX(t *testing.T) {
	fLongIn, err := files.NewFile("../../test/data/sample_outline.pmd")
	assert.NoError(t, err)

	long, err := ParseFromFile(fLongIn)
	assert.NoError(t, err)

	longTeX := long.LaTeX()
	assert.Equal(t, expectedLongLaTeXOutput, longTeX)

	fShortIn, err := files.NewFile("../../test/data/sample_short_outline.pmd")
	assert.NoError(t, err)

	short, err := ParseFromFile(fShortIn)
	assert.NoError(t, err)

	shortTeX := short.LaTeX()
	assert.Equal(t, expectedShortLaTeXOutput, shortTeX)
}

func TestBibliography(t *testing.T) {
	fShortIn, err := files.NewFile("../../test/data/sample_short_outline.pmd")
	assert.NoError(t, err)

	short, err := ParseFromFile(fShortIn)
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
