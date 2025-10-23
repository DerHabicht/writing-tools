package rexy

/*

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/derhabicht/writing-tools/internal/logging"
	"github.com/derhabicht/writing-tools/internal/papers"
)

var logLevel string

// TODO: Consider breaking this out into subcommands for compiling PDFs and generating BibTeX files.
var rootCmd = &cobra.Command{
	Use:   "rexy <paper.omd> <output.pdf>",
	Short: "Generate outlines, talking, and bullet background papers from special markdown files.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.Logger{}

		pmd, err := files.NewFile(args[0])
		if err != nil {
			logger.Error().Err(err).Str("filename", args[0]).Msg("failed to create reference to input file")
			os.Exit(1)
		}

		var outputFilePath string
		if len(args) == 2 {
			outputFilePath = args[1]
		} else {
			outputFilePath = filepath.Join(pmd.Dir().Path(), fmt.Sprintf("%s.%s", pmd.Base(), "pdf"))
		}
		outputFile, err := files.NewFile(outputFilePath)
		if err != nil {
			logger.Error().Err(err).Str("filename", outputFilePath).Msg("failed to create reference to output file")
			os.Exit(1)
		}

		papers.BuildPaper(pmd, outputFile, logger)
	},
}

func Execute(version string) {
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "")

	logging.InitLogging(logLevel, true)
}
*/
