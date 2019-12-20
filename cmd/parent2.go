package cmd

import(
	"fmt"
	"github.com/l1redd/cobraexample/cmd/children"
	"github.com/spf13/cobra"
)

type Parent2Args struct {
	one int32
	two int32
	childArgs children.ChildArgs
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
	fmt.Println("What's up!")
	fmt.Printf("Parent2 -- first: %d, second: %d \n", args.one, args.two)

	childArgs := args.childArgs
	fmt.Println("List!")
	for _, item := range childArgs.List{
		fmt.Println(item)
	}
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

	childArgs := children.ParseChildArgs(cmd)

	return Parent2Args{
		one: first,
		two: second,
		childArgs: childArgs,
	}
}

func init () {
	Parent2Cmd.PersistentFlags().Int32P("one", "o", 0, "first arg")
	Parent2Cmd.PersistentFlags().Int32P("two", "t", 0, "second arg")

	children.DefineFlags(Parent2Cmd)
}
