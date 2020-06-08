package cmd

import (
	"github.com/spf13/cobra"
)

// NewWmanCmd creates a new root command for wman
func NewWmanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "wman",
		Short:        "Weather Man CLI",
		Long:         `Weather Man (wman) CLI for capturing weather information.`,
		SilenceUsage: true,
		Example: `  # Run reverse function
  wman reverse nofluff
`,
	}
	cmd.AddCommand(newPrintCmd())
	return cmd
}
