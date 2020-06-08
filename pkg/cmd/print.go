package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	printExample = `  # Prints a hello message to the passed in name
  wman print nofluff`
)

// newPrintCmd returns a new initialized instance of the print sub command
func newPrintCmd() *cobra.Command {
	printCmd := &cobra.Command{
		Use:     "print",
		Short:   "Print hello to the passed in argument",
		Example: printExample,
		RunE:    PrintCmd,
	}

	return printCmd
}

// PrintCmd performs the print sub command
func PrintCmd(cmd *cobra.Command, args []string) error {
	// might be worth err checking args
	fmt.Printf("Hello %q\n", args[0])
	return nil
}
