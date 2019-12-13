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
var childCommand = &cobra.Command{
	Use:   "child",
	Short: "Child command",
	Run: func(cmd *cobra.Command, args []string) {
		ChildRun(cmd)
	},
}

func NewChildCmd() *cobra.Command {
	//return &cobra.Command{
	//	Use:   "child",
	//	Short: "Child command",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		ChildRun(cmd)
	//	},
	//}
	return childCommand
}

func ChildRun(cmd *cobra.Command) {
	args := parseChildArgs(cmd)
	fmt.Printf("Child -- first: %d, second: %d, third: %d, fourth: %d \n", args.one, args.two, args.three, args.four)

}

func parseChildArgs(cmd *cobra.Command) (args ChildArgs) {
	fmt.Printf("parsing args \n")
	first, err := cmd.Flags().GetInt32("one")
	if err != nil {
		fmt.Printf("error parsing one: %s \n", err.Error())
		return
	}
	fmt.Printf("one: %d", first)

	second, err := cmd.Flags().GetInt32("two")
	if err != nil {
		fmt.Printf("error parsing two: %s \n", err.Error())
		return
	}

	fmt.Printf("two: %d", second)

	third, err := cmd.Flags().GetInt32("three")
	if err != nil {
		fmt.Printf("error parsing three: %s \n", err.Error())
		return
	}

	fmt.Printf("three: %d", third)

	fourth, err := cmd.Flags().GetInt32("four")
	if err != nil {
		fmt.Printf("error parsing four: %s \n", err.Error())
		return
	}

	fmt.Printf("four: %d \n", fourth)

	return ChildArgs{
		one: first,
		two: second,
		three: third,
		four: fourth,
	}
}

func init() {
	childCommand.Flags().Int32P("three", "r", 0, "third arg")
	childCommand.Flags().Int32P("four", "f", 0, "fourth arg")
}
