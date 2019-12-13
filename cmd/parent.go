package cmd

import(
	"fmt"
	"github.com/l1redd/cobraexample/cmd/children"
	"github.com/spf13/cobra"
)

type ParentArgs struct {
	one int32
	two int32
	childArgs children.ChildArgs
}

var ParentCmd = &cobra.Command{
	Use:   "parent",
	Short: "parent command",
	Run: func(cmd *cobra.Command, args []string) {
		ParentRun(cmd)
	},
}

func ParentRun(cmd *cobra.Command) {
	args := parseParentArgs(cmd)
	fmt.Printf("What's up! \n")

	if args.one != 0 && args.two != 0 {
		fmt.Printf("Parent -- first: %d, second: %d \n", args.one, args.two)
	}else{
		fmt.Printf("in parent, printing child \n")
		children.PrintTheThing(args.childArgs)
	}
}

func parseParentArgs(cmd *cobra.Command) (args ParentArgs) {
	first, err := cmd.Flags().GetInt32("one")
	if err != nil {
		return
	}

	second, err := cmd.Flags().GetInt32("two")
	if err != nil {
		return
	}

	childArgs := children.ParseChildArgs(cmd)

	return ParentArgs{
		one: first,
		two: second,
		childArgs: childArgs,
	}
}

func init () {
	ParentCmd.PersistentFlags().Int32P("one", "o", 0, "first arg")
	ParentCmd.PersistentFlags().Int32P("two", "t", 0, "second arg")

	children.DefineFlags(ParentCmd)
}
