// Orion - Ultra-lightweight personal AI agent
// License: MIT
//
// Copyright (c) 2026 Orion contributors

package config

import (
	"os"
	"path/filepath"

	"github.com/varunnegi07/orion/pkg"
)

// Runtime environment variable keys for the orion process.
// These control the location of files and binaries at runtime and are read
// directly via os.Getenv / os.LookupEnv. All orion-specific keys use the
// ORION_ prefix. Reference these constants instead of inline string
// literals to keep all supported knobs visible in one place and to prevent
// typos.
const (
	// EnvHome overrides the base directory for all orion data
	// (config, workspace, skills, auth store, …).
	// Default: ~/.orion
	EnvHome = "ORION_HOME"

	// EnvConfig overrides the full path to the JSON config file.
	// Default: $ORION_HOME/config.json
	EnvConfig = "ORION_CONFIG"

	// EnvBuiltinSkills overrides the directory from which built-in
	// skills are loaded.
	// Default: <cwd>/skills
	EnvBuiltinSkills = "ORION_BUILTIN_SKILLS"

	// EnvBinary overrides the path to the orion executable.
	// Used by the web launcher when spawning the gateway subprocess.
	// Default: resolved from the same directory as the current executable.
	EnvBinary = "ORION_BINARY"

	// EnvGatewayHost overrides the host address for the gateway server.
	// Default: "localhost"
	EnvGatewayHost = "ORION_GATEWAY_HOST"
)

func GetHome() string {
	homePath, _ := os.UserHomeDir()
	if orionHome := os.Getenv(EnvHome); orionHome != "" {
		homePath = orionHome
	} else if homePath != "" {
		homePath = filepath.Join(homePath, pkg.DefaultOrionHome)
	}
	if homePath == "" {
		homePath = "."
	}
	return homePath
}
