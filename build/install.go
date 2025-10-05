package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/derhabicht/writing-tools/internal/config"
	"github.com/derhabicht/writing-tools/internal/logging"
)

func CreateConfigDirectories() error {
	cfg, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}

	cfgDir := cfg.Path()

	logging.Info().Str("dir", cfgDir).Msg("creating config directory")
	err = os.Mkdir(cfgDir, 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cfgDir).Msg("failed to create config directory")
	}

	logging.Info().Str("subdir", "cfg").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "cfg"), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "cfg").Msg("failed to create subdirectory")
	}

	logging.Info().Str("subidr", "assets").Msg("creating subdirectory")
	err = os.Mkdir(filepath.Join(cfgDir, "assets"), 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("subdir", "assets").Msg("failed to create subdirectory")
	}

	return nil
}

func CreateCacheDirectory() error {
	cache, err := config.CacheDir()
	if err != nil {
		return errors.WithStack(err)
	}

	cacheDir := cache.Path()

	logging.Info().Str("dir", cacheDir).Msg("creating cache directory")
	err = os.Mkdir(cacheDir, 0700)
	err = ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("dir", cacheDir).Msg("failed to create directory")
	}

	return nil
}

func ClearFileExistsError(err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()
	if strings.Contains(msg, "file exists") {
		return nil
	}

	return err
}

func ClearFileDoesNotExistError(err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()
	if strings.Contains(msg, "no such file or directory") {
		return nil
	}

	return err
}

func main() {
	logging.InitLogging("info", true)

	err := CreateConfigDirectories()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create config directories")
		os.Exit(1)
	}

	err = CreateCacheDirectory()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create cache directories")
		os.Exit(1)
	}
}
