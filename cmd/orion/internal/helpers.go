package internal

import (
	"os"
	"path/filepath"

	"github.com/varunnegi07/orion/pkg"
	"github.com/varunnegi07/orion/pkg/config"
	"github.com/varunnegi07/orion/pkg/logger"
)

const Logo = pkg.Logo

// GetPicoclawHome returns the orion home directory.
// Priority: $ORION_HOME > ~/.orion
func GetPicoclawHome() string {
	return config.GetHome()
}

func GetConfigPath() string {
	if configPath := os.Getenv(config.EnvConfig); configPath != "" {
		return configPath
	}
	return filepath.Join(GetPicoclawHome(), "config.json")
}

func LoadConfig() (*config.Config, error) {
	cfg, err := config.LoadConfig(GetConfigPath())
	if err != nil {
		return nil, err
	}
	logger.SetLevelFromString(cfg.Gateway.LogLevel)
	return cfg, nil
}

// FormatVersion returns the version string with optional git commit
// Deprecated: Use pkg/config.FormatVersion instead
func FormatVersion() string {
	return config.FormatVersion()
}

// FormatBuildInfo returns build time and go version info
// Deprecated: Use pkg/config.FormatBuildInfo instead
func FormatBuildInfo() (string, string) {
	return config.FormatBuildInfo()
}

// GetVersion returns the version string
// Deprecated: Use pkg/config.GetVersion instead
func GetVersion() string {
	return config.GetVersion()
}
