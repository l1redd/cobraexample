package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)

type ChildArgs struct {
	one int32
	two int32
	three int32
	four int32
}
var ChildCommand = &cobra.Command{
	Use:   "child",
	Short: "Child command",
	Run: func(cmd *cobra.Command, args []string) {
		ChildRun(cmd)
	},
}

func ChildRun(cmd *cobra.Command) {
	args := parseChildArgs(cmd)
	fmt.Printf("Child -- first: %d, second: %d, third: %d, fourth: %d", args.one, args.two, args.three, args.four)

}

func parseChildArgs(cmd *cobra.Command) (args ChildArgs) {
	first, err := cmd.Flags().GetInt32("one")
	if err != nil {
		return
	}

	second, err := cmd.Flags().GetInt32("two")
	if err != nil {
		return
	}

	third, err := cmd.PersistentFlags().GetInt32("three")
	if err != nil {
		return
	}

	fourth, err := cmd.PersistentFlags().GetInt32("four")
	if err != nil {
		return
	}

	return ChildArgs{
		one: first,
		two: second,
		three: third,
		four: fourth,
	}
}

func init() {
	ChildCommand.Flags().Int32P("three", "th", 0, "third arg")
	ChildCommand.Flags().Int32P("four", "f", 0, "fourth arg")
}
