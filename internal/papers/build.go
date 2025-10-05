package papers

import (
	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/pkg/errors"

	"github.com/derhabicht/writing-tools/internal/config"
	"github.com/derhabicht/writing-tools/internal/logging"
	"github.com/derhabicht/writing-tools/pkg/papers"
)

func configureLaTeXCompiler(logger logging.Logger) (*latex.Compiler, error) {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find cache directory")
	}

	compiler := latex.NewCompiler(latex.XeLaTeX, latex.Biber, *cacheDir)

	return &compiler, nil
}

func generateAssets(paper papers.Paper, output files.File) ([]files.File, error) {
	ol, ok := paper.(*papers.Outline)
	if !ok {
		return nil, nil
	}

	bibtex := ol.BibTeX()
	if bibtex == "" {
		return nil, nil
	}

	refs, err := output.Dir().CreateFile("references.bib")
	if err != nil {
		return nil, errors.WithMessage(err, "unable to create references.bib")
	}

	err = refs.WriteString(bibtex)
	if err != nil {
		return nil, errors.WithMessage(err, "unable to write references.bib")
	}

	return []files.File{refs}, nil
}

func BuildPaper(input, output files.File, logger logging.Logger) error {
	paper, err := papers.ParseFromFile(input)
	if err != nil {
		return errors.WithMessage(err, "failed to parse PMD")
	}

	assets, err := generateAssets(paper, output)
	if err != nil {
		return errors.WithStack(err)
	}

	compiler, err := configureLaTeXCompiler(logger)
	if err != nil {
		return errors.WithMessage(err, "failed to configure LaTeX compiler")
	}

	err = compiler.GenerateLaTeX(paper, output, assets)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.CompileLaTeX(output)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
