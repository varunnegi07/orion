// Orion - Ultra-lightweight personal AI agent
// Inspired by and based on nanobot: https://github.com/HKUDS/nanobot
// License: MIT
//
// Copyright (c) 2026 Orion contributors

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/varunnegi07/orion/cmd/orion/internal"
	"github.com/varunnegi07/orion/cmd/orion/internal/agent"
	"github.com/varunnegi07/orion/cmd/orion/internal/auth"
	"github.com/varunnegi07/orion/cmd/orion/internal/cliui"
	"github.com/varunnegi07/orion/cmd/orion/internal/cron"
	"github.com/varunnegi07/orion/cmd/orion/internal/gateway"
	"github.com/varunnegi07/orion/cmd/orion/internal/mcp"
	"github.com/varunnegi07/orion/cmd/orion/internal/migrate"
	"github.com/varunnegi07/orion/cmd/orion/internal/model"
	"github.com/varunnegi07/orion/cmd/orion/internal/onboard"
	"github.com/varunnegi07/orion/cmd/orion/internal/skills"
	"github.com/varunnegi07/orion/cmd/orion/internal/status"
	"github.com/varunnegi07/orion/cmd/orion/internal/version"
	"github.com/varunnegi07/orion/pkg/config"
	"github.com/varunnegi07/orion/pkg/updater"
)

var rootNoColor bool

func syncCliUIColor(root *cobra.Command) {
	no, _ := root.PersistentFlags().GetBool("no-color")
	cliui.Init(no || os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb")
}

// earlyColorDisabled matches lipgloss/banner behavior from env and argv before Cobra parses flags.
func earlyColorDisabled() bool {
	if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		return true
	}
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "--no-color" || arg == "--no-color=true" || arg == "--no-color=1" {
			return true
		}
	}
	return false
}

func NewOrionCommand() *cobra.Command {
	short := fmt.Sprintf("%s Orion — personal AI assistant", internal.Logo)
	long := fmt.Sprintf(`%s Orion is a lightweight personal AI assistant.

Version: %s`, internal.Logo, config.FormatVersion())

	cmd := &cobra.Command{
		Use:   "orion",
		Short: short,
		Long:  long,
		Example: `orion version
orion onboard
orion --no-color status`,
		SilenceErrors: true,
		// Avoid plain UsageString() on stderr/stdout when a command fails; cliui
		// renders matching panels on stderr instead.
		SilenceUsage: true,
		PersistentPreRun: func(c *cobra.Command, _ []string) {
			syncCliUIColor(c.Root())
		},
	}

	cmd.PersistentFlags().BoolVar(&rootNoColor, "no-color", false,
		"Disable colors (boxed layout unchanged)")

	cmd.SetHelpFunc(func(c *cobra.Command, _ []string) {
		syncCliUIColor(c.Root())
		fmt.Fprint(c.OutOrStdout(), cliui.RenderCommandHelp(c))
	})

	cmd.AddCommand(
		onboard.NewOnboardCommand(),
		agent.NewAgentCommand(),
		auth.NewAuthCommand(),
		gateway.NewGatewayCommand(),
		status.NewStatusCommand(),
		cron.NewCronCommand(),
		mcp.NewMCPCommand(),
		migrate.NewMigrateCommand(),
		skills.NewSkillsCommand(),
		model.NewModelCommand(),
		updater.NewUpdateCommand("orion"),
		version.NewVersionCommand(),
	)

	return cmd
}

const (
	colorBlue = "\033[1;38;2;62;93;185m"
	colorRed  = "\033[1;38;2;213;70;70m"
	banner    = "\r\n" +
		colorBlue + " ██████╗ ██████╗ ██╗ ██████╗ " + colorRed + "███╗   ██╗\n" +
		colorBlue + "██╔═══██╗██╔══██╗██║██╔═══██╗" + colorRed + "████╗  ██║\n" +
		colorBlue + "██║   ██║██████╔╝██║██║   ██║" + colorRed + "██╔██╗ ██║\n" +
		colorBlue + "██║   ██║██╔══██╗██║██║   ██║" + colorRed + "██║╚██╗██║\n" +
		colorBlue + "╚██████╔╝██║  ██║██║╚██████╔╝" + colorRed + "██║ ╚████║\n" +
		colorBlue + " ╚═════╝ ╚═╝  ╚═╝╚═╝ ╚═════╝ " + colorRed + "╚═╝  ╚═══╝\n " +
		"\033[0m\r\n"
	plainBanner = "\r\n" +
		" ██████╗ ██████╗ ██╗ ██████╗ ███╗   ██╗\n" +
		"██╔═══██╗██╔══██╗██║██╔═══██╗████╗  ██║\n" +
		"██║   ██║██████╔╝██║██║   ██║██╔██╗ ██║\n" +
		"██║   ██║██╔══██╗██║██║   ██║██║╚██╗██║\n" +
		"╚██████╔╝██║  ██║██║╚██████╔╝██║ ╚████║\n" +
		" ╚═════╝ ╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝\n " +
		"\r\n"
)

func main() {
	cliui.Init(earlyColorDisabled())

	if earlyColorDisabled() {
		fmt.Print(plainBanner)
	} else {
		fmt.Printf("%s", banner)
	}

	tzEnv := os.Getenv("TZ")
	if tzEnv != "" {
		fmt.Println("TZ environment:", tzEnv)
		zoneinfoEnv := os.Getenv("ZONEINFO")
		fmt.Println("ZONEINFO environment:", zoneinfoEnv)
		loc, err := time.LoadLocation(tzEnv)
		if err != nil {
			fmt.Println("Error loading time zone:", err)
		} else {
			fmt.Println("Time zone loaded successfully:", loc)
			time.Local = loc //nolint:gosmopolitan // We intentionally set local timezone from TZ env
		}
	}

	cmd := NewOrionCommand()
	last, err := cmd.ExecuteC()
	if err != nil {
		syncCliUIColor(cmd)
		fmt.Fprint(os.Stderr, cliui.FormatCLIError(err.Error(), last))
		os.Exit(1)
	}
}
