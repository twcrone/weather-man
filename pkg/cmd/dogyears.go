package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

var (
	dogYearExample = `  # Calculate dog years
  wman dogyears`
)

func newDogYearCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dogyear",
		Short:   "Calculates dogyears",
		Example: dogYearExample,
		RunE:    DogYearCmd,
	}
	return cmd
}

func DogYearCmd(cmd *cobra.Command, args []string) error {
	fmt.Println("Enter Age:")
	reader := bufio.NewReader(os.Stdin)
	a, err := reader.ReadString('\n')
	if errors.Is(err, bufio.ErrBufferFull) {
		return fmt.Errorf("buffer full %w", err)
	}
	if err != nil {
		return err
	}
	a = strings.TrimSpace(a)
	age, err := strconv.Atoi(a)
	if err != nil {
		return err
	}
	dogyears := age * 7

	fmt.Printf("your age of %d is %v in dog years\n", age, dogyears)
	return nil
}
