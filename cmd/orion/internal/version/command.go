package version

import (
	"github.com/spf13/cobra"

	"github.com/varunnegi07/orion/cmd/orion/internal"
	"github.com/varunnegi07/orion/cmd/orion/internal/cliui"
	"github.com/varunnegi07/orion/pkg/config"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show version information",
		Run: func(_ *cobra.Command, _ []string) {
			printVersion()
		},
	}

	return cmd
}

func printVersion() {
	build, goVer := config.FormatBuildInfo()
	cliui.PrintVersion(internal.Logo, "orion "+config.FormatVersion(), build, goVer)
}
