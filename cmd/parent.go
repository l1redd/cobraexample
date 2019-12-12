package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)

type ParentArgs struct {
	one int32
	two int32
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
	fmt.Printf("Parent -- first: %d, second: %d", args.one, args.two)
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
	return ParentArgs{
		one: first,
		two: second,
	}
}

//func Execute() error {
//	return rootCmd.Execute()
//}

func init () {
	ParentCmd.PersistentFlags().Int32P("one", "o", 0, "first arg")
	ParentCmd.PersistentFlags().Int32P("two", "t", 0, "second arg")

	ParentCmd.AddCommand(ChildCommand)
}
