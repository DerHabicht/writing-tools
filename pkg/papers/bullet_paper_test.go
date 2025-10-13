package papers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

const expectedPointPaperOutput = `\documentclass[point]{usafpaper}

\title{Writing Point Papers}
\date{22 Jul 2014}
\author{SSgt Hines-Davis}
\authorOffice{SOC/ES}
\authorPhone{(555) 555-1234}
\typist{ahd}

\begin{document}
\maketitle

\begin{usafitem}
    \item Function: Minimal text outline of a single issue to quickly inform others
    \begin{usafitem}
        \item Short: Supports extemporaneous speaking opportunities (e.g. "elevator speech")
        \item Subject matter: requires prior preparation/immersion in background and details
    \end{usafitem}
    \item Format: Baseline standards below are flexible to save space and/or conform to user's needs
    \begin{usafitem}
        \item Overall: single page with short telegraphic bullets
        \item Page setup
        \begin{usafitem}
            \item Title: centered; all capital letters; long titles wrap single-spaced under third line (FYI: Use manual line breaks for long title readability or to visually balance the lines.)
            \item Margins: One (1) inch all around
            \item Headings (e.g. PURPOSE, DISCUSSION) are optional
            \item Dashes: single before major thoughts; multiple for subordinate thoughts
            \item Line spacing and text wrapping: single-space within bullets and double-space between bullets; wrap bullets as in this example (wrapped bullets are rare in Point Papers)
            \item Punctuation: open punctuation style---ending punctuation not required
            \item Indentification line: One (1) inch from bottom, flush left; *alternately placed in the footer, one half (1/2) inch from the bottom, flush left as in this example*
		    \begin{usafitem}
                \item Rank/Title and last name of the point of contact (POC)
			    \item Organization/office symbol
			    \item Telephone: Full DSN or 10-digit commercial
			    \item Typist's initials ("ahd" in the example) (may be the POC or someone else)
			    \item Date in "DD Mmm YY" Format
		    \end{usafitem}
	    \end{usafitem}
	    \item Classifed content: See DOD 5200.1-R/AFI 31-401 to prepare classified Papers
    \end{usafitem}
    \item Recommendations or conclusion: Give your point a solid sense of the way ahead or closure
\end{usafitem}
\end{document}
`

const expectedTalkingPaperOutput = `\documentclass[point]{usafpaper}

\title{Writing Talking Papers}
\date{4 Apr 2014}
\author{Ms. Kenney}
\authorOffice{Holm Center/CCS}
\authorPhone{(555) 555-1234}
\typist{jjk}

\begin{document}
\maketitle

\begin{usafitem}
    \item Function: Speaking notes that outline and narrate a single issue to inform others during planned/scheduled oral presentations
    \begin{usafitem}
        \item Provides both the outline of a single issue and quick-reference content on key points, facts, data, positions, or frequently asked questions
        \item Can stand alone for basic understanding; better with knowledge topic and related issues
    \end{usafitem}
    \item Format: Baseline standards below are flexible to save space and/or conform to user's needs
    \begin{usafitem}
        \item Normally a single page (FYI: avoid lengthy chronologies and excessive detail)
        \item Headings (e.g. PURPOSE, DISCUSSION) are optional; save space by eliminating headings, by using run-in headings (e.g. Format, Flow), or both as in this example
        \item Bullets are short phrases or statements; telegraphic wording saves space
        \item Format the title, margins, dashes, line spacing, and text wrapping as in the Point Paper
        \item Punctuation: Normal rules apply for complete sentences and paragraphs; bullets may have internal punctuation but do not require closing punctuation
        \item Identification line: Format the same as for the Point Paper and place only on the first page, one (1) inch from bottom, fluh left; *alternatively placed in the footer, one half (1/1) inch from the bottom, flush left*
        \item Page numbering (if longer than one page): Place the number for page 2 onwards at the top of page, one-half inch from the top and flush with the right margins
        \item Classified content: See DOD 5200.1-R/AFI 31-401 to prepare classified Papers
    \end{usafitem}
    \item Flow: Clear statement, logical support, and closure
    \begin{usafitem}
        \item Make it memorable with a literary device and solid sense of closure
        \item Include additional information as for your information (FYI) note at the appropriate place in the text/attached background paper (FYI: This is an example FYI note.)
    \end{usafitem}
    \item Recommendations or conclusion: Give your talk a solid sense of the way ahead or closure
\end{usafitem}
\end{document}
`

const expectedBulletBackgroundPaperOutput = `\documentclass[bullet-background]{usafpaper}

\usepackage[authordate,backend=biber]{biblatex-chicago}
\addbibresource{references.bib}

\title{Writing Bullet Background Papers}
\date{20 Jun 2014}
\author{Ms. McKitt}
\authorOffice{SAASS/AS}
\authorPhone{(555) 555-1234}
\typist{cdb}

\begin{document}
\maketitle

\section*{Purpose\footcite[p.224]{afh33-337}}

Discuss the functions, format, and tips for building a Bullet Background Paper (BBP). The purpose statement informs readers on both the purpose and main points of the paper. It may be a single sentence, as in this example, or a short paragraph in length.

\section*{Functions}

\begin{usafitem}
    \item Concise background information on a single idea
    \item Summary of a staff package
    \item Accomplishment Summary
    \item Chronological of a problem
    \item Support for an attached paper
    \item Program information
    \item Information to provide a response to just about any task
\end{usafitem}

\section*{Format}

\begin{usafitem}
    \item Flexible to save space and/or conform to user's needs
    \item Main ideas may be presented as headings (e.g. FUNCTION, FORMAT, TIPS) or as subordinate ideas to broader generic headings (e.g. PURPOSE, DISCUSSION, FINDINGS, RECOMMENDATIONS) using dashes and indentation.
    \item Secondary items follow with a single dash; tertiary and further subordinate items follow with multiple indented dashes. Any item can be as short as a word or as long as several sentences.
    \item Page setup
    \begin{usafitem}
        \item Title: centered; all capital letters; long titles wrap single-spaced under third line (FYI: Use manual line breaks for long title readability or to visually balance the lines.)
        \item Margins: One (1) inch all around
        \item Dashes: single before major thoughts; multiple for subordinate thoughts
        \item Line spacing and text wrapping: single-space within bullets and double-space between bullets; wrap bullets as in this example
        \item Punctuation: Normal rules apply for complete sentences and paragraphs; bullets may have internal punctuation but do not require closing punctuation
        \item Identification line: Format the same as for the Point Paper and place only on the first page, one (1) inch from bottom, flush left; *alternatively placed in the footer, one half (1/2) inch from the bottom, flush left*
        \item Page numbering (if longer than one page): Place the page number for page 2 onwards at the top of page, one-half inch from the top and flush with the right margins
        \item Classified content: See DOD 5200.1-R/AFI 31-401 to prepare classified Papers
    \end{usafitem}
\end{usafitem}

\section*{Tips}

\begin{usafitem}
    \item Write the BBP according to the knowledge level of the expected readers
    \item Emphasize main points by using them as headings or short bullets. If additional information is needed, refer to it in the text of the BBP and attach the referenced documents.
    \item Strive to minimize the length to communicate quickly with impact
\end{usafitem}

\section*{Conclusion}

The BBP serves many purposes and is no longer than it needs to be to convey the message. The formats for many of the BBP elements are the same as they are for other Air Force papers. End with recommendations or conclusions that bring the discussion to a close.

\pagebreak

\printbibliography

\end{document}
`

const expectedBulletBackgroundPaperBibTeXOutput = `
@manual{afh33-337,
    title = {AFH 33-337},
    author = {Department of the Air Force},
    date = {2015-05-27},
    subtitle = {The Tongue and Quill},
    organization = {United States Air Force},
    publiher = {Department of the Air Force},
    address = {Washington, DC},
}
`

func TestPointPaperImplementsInterfaces(t *testing.T) {
	var _ Paper = (*BulletPaper)(nil)
	var _ Point = (*BPoint)(nil)
}

func TestSamplePapersCompile(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	pp := latex.LaTeXString(expectedPointPaperOutput)
	tp := latex.LaTeXString(expectedTalkingPaperOutput)
	bbp := latex.LaTeXString(expectedBulletBackgroundPaperOutput)

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
	dataDir := files.NewDirectory(filepath.Join(projDir, "test", "data"))

	ppOut, err := dataDir.NewFile("sample_point_paper.pdf")
	assert.NoError(t, err)

	tpOut, err := dataDir.NewFile("sample_talking_paper.pdf")
	assert.NoError(t, err)

	bbpOut, err := dataDir.NewFile("sample_bullet_background_paper.pdf")
	assert.NoError(t, err)

	c := latex.NewCompiler(latex.XeLaTeX, latex.NoBib, *cacheDir)

	err = c.GenerateLaTeX(pp, ppOut, assets)
	assert.NoError(t, err)
	err = c.CompileLaTeX(ppOut)
	assert.NoError(t, err)
	assert.FileExists(t, ppOut.FullPath())
	err = ppOut.Remove()
	assert.NoError(t, err)

	err = c.GenerateLaTeX(tp, tpOut, assets)
	assert.NoError(t, err)
	err = c.CompileLaTeX(tpOut)
	assert.NoError(t, err)
	assert.FileExists(t, tpOut.FullPath())
	err = tpOut.Remove()
	assert.NoError(t, err)

	c = latex.NewCompiler(latex.XeLaTeX, latex.Biber, *cacheDir)
	ref, err := cacheDir.CreateFile("references.bib")
	assert.NoError(t, err)
	err = ref.WriteString(expectedBulletBackgroundPaperBibTeXOutput)
	assert.NoError(t, err)
	err = c.GenerateLaTeX(bbp, bbpOut, assets)
	assert.NoError(t, err)
	err = c.CompileLaTeX(bbpOut)
	assert.NoError(t, err)
	assert.FileExists(t, bbpOut.FullPath())
	err = bbpOut.Remove()
	assert.NoError(t, err)
}

func generateExpectedPointPaper() *BulletPaper {
	pm := PaperMeta{
		PaperType: PointPaper,
		Title:     "Writing Point Papers",
		Author:    "SSgt Hines-Davis",
		Office:    "SOC/ES",
		Contact:   "(555) 555-1234",
		Typist:    "ahd",
		Date:      date.New(2014, 7, 22),
		Comment:   "Sample point paper found on p. 222 of AFH 33-337 (the Tongue & Quill).",
	}

	l1Points := make([]*BPoint, 3)

	l1Points[0] = NewBPoint(0, "Function: Minimal text outline of a single issue to quickly inform others")
	l1Points[0].AddSubpoints(NewBPoint(1, "Short: Supports extemporaneous speaking opportunities (e.g. \"elevator speech\")"))
	l1Points[0].AddSubpoints(NewBPoint(1, "Subject matter: requires prior preparation/immersion in background and details"))

	l1Points[1] = NewBPoint(0, "Format: Baseline standards below are flexible to save space and/or conform to user's needs")
	l1Points[1].AddSubpoints(NewBPoint(1, "Overall: single page with short telegraphic bullets"))

	pagePoint := NewBPoint(1, "Page setup")
	pagePoint.AddSubpoints(NewBPoint(2, "Title: centered; all capital letters; long titles wrap single-spaced under third line (FYI: Use manual line breaks for long title readability or to visually balance the lines.)"))
	pagePoint.AddSubpoints(NewBPoint(2, "Margins: One (1) inch all around"))
	pagePoint.AddSubpoints(NewBPoint(2, "Headings (e.g. PURPOSE, DISCUSSION) are optional"))
	pagePoint.AddSubpoints(NewBPoint(2, "Dashes: single before major thoughts; multiple for subordinate thoughts"))
	pagePoint.AddSubpoints(NewBPoint(2, "Line spacing and text wrapping: single-space within bullets and double-space between bullets; wrap bullets as in this example (wrapped bullets are rare in Point Papers)"))
	pagePoint.AddSubpoints(NewBPoint(2, "Punctuation: open punctuation style---ending punctuation not required"))

	identPoint := NewBPoint(2, "Indentification line: One (1) inch from bottom, flush left; *alternately placed in the footer, one half (1/2) inch from the bottom, flush left as in this example*")
	identPoint.AddSubpoints(NewBPoint(3, "Rank/Title and last name of the point of contact (POC)"))
	identPoint.AddSubpoints(NewBPoint(3, "Organization/office symbol"))
	identPoint.AddSubpoints(NewBPoint(3, "Telephone: Full DSN or 10-digit commercial"))
	identPoint.AddSubpoints(NewBPoint(3, "Typist's initials (\"ahd\" in the example) (may be the POC or someone else)"))
	identPoint.AddSubpoints(NewBPoint(3, "Date in \"DD Mmm YY\" Format"))
	pagePoint.AddSubpoints(identPoint)

	l1Points[1].AddSubpoints(pagePoint)

	l1Points[1].AddSubpoints(NewBPoint(1, "Classifed content: See DOD 5200.1-R/AFI 31-401 to prepare classified Papers"))

	l1Points[2] = NewBPoint(0, "Recommendations or conclusion: Give your point a solid sense of the way ahead or closure")

	paper := NewBulletPaper(pm, nil)
	paper.AddBPoints(l1Points...)

	return paper
}

func TestParsePointPaper(t *testing.T) {
	expected := generateExpectedPointPaper()

	f, err := files.NewFile("../../test/data/sample_point_paper.pmd")
	assert.NoError(t, err)

	paper, err := ParseFromFile(f)
	assert.NoError(t, err)

	assert.Equal(t, expected, paper)
}
