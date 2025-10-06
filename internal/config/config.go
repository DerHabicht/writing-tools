package config

import (
	"os"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Configuration keys

// Directory functions
const progName = "writing-tools"

// ConfigDir returns the directory where configuration files and assets are stored.
func ConfigDir() (*files.Directory, error) {
	hd, err := os.UserConfigDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find user home directory")
	}

	return files.NewDirectory(filepath.Join(hd, progName)), nil
}

// CacheDir returns the directory where LaTeX files are built.
func CacheDir() (*files.Directory, error) {
	cd, err := os.UserCacheDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find user cache directory")
	}

	return files.NewDirectory(filepath.Join(cd, progName)), nil
}

// Config setup functions

// SetConfigType wraps Viper's SetConfigType function
func SetConfigType(t string) {
	viper.SetConfigType(t)
}

// SetConfigName wraps Viper's SetConfigName function
func SetConfigName(name string) {
	viper.SetConfigName(name)
}

// AddConfigPath wraps Viper's AddConfigPath function
func AddConfigPath(path string) {
	viper.AddConfigPath(path)
}

// SetDefault wraps Viper's SetDefault function.
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

// ReadInConfig wraps Viper's ReadInConfig function.
func ReadInConfig() error {
	return viper.ReadInConfig()
}

// WriteConfigAs wraps Viper's WriteConfigAs function.
func WriteConfigAs(path string) error {
	return viper.WriteConfigAs(path)
}

// Configuration access functions

func Set(key string, value interface{}) {
	viper.Set(key, value)
}

// GetFloat64 wraps Viper's GetFloat64 function.
func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

// GetInt wraps Viper's GetInt function.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetString wraps Viper's GetString function.
func GetString(key string) string {
	return viper.GetString(key)
}
