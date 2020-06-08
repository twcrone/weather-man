package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	wstring "github.com/twcrone/wman/pkg/string"
)

type printOptions struct {
	Reverse bool
}

var (
	printExample = `  # Prints a hello message to the passed in name
  wman print nofluff`
)

// newPrintCmd returns a new initialized instance of the print sub command
func newPrintCmd() *cobra.Command {
	opts := &printOptions{}
	printCmd := &cobra.Command{
		Use:     "print",
		Short:   "Print hello to the passed in argument",
		Example: printExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return PrintCmd(opts, args)
		},
	}
	printCmd.Flags().BoolVar(&opts.Reverse, "reverse", false, "If set to true, it reverses the argument passed in (default \"false\")")
	return printCmd
}

// PrintCmd performs the print sub command
func PrintCmd(options *printOptions, args []string) error {
	if options.Reverse {
		fmt.Printf("Hello %q\n", wstring.Reverse(args[0]))

	} else {
		fmt.Printf("Hello %q\n", args[0])
	}
	return nil
}
