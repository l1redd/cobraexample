package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)

type Parent2Args struct {
	one int32
	two int32
}

var Parent2Cmd = &cobra.Command{
	Use:   "parent2",
	Short: "parent2 command",
	Run: func(cmd *cobra.Command, args []string) {
		Parent2Run(cmd)
	},
}

func Parent2Run(cmd *cobra.Command) {
	args := parseParent2Args(cmd)
	fmt.Printf("What's up! \n")
	fmt.Printf("Parent2 -- first: %d, second: %d \n", args.one, args.two)
}

func parseParent2Args(cmd *cobra.Command) (args Parent2Args) {
	first, err := cmd.Flags().GetInt32("one")
	if err != nil {
		return
	}

	second, err := cmd.Flags().GetInt32("two")
	if err != nil {
		return
	}
	return Parent2Args{
		one: first,
		two: second,
	}
}

func init () {
	Parent2Cmd.PersistentFlags().Int32P("one", "o", 0, "first arg")
	Parent2Cmd.PersistentFlags().Int32P("two", "t", 0, "second arg")

	childCmd := NewChildCmd()
	Parent2Cmd.AddCommand(childCmd)
}
