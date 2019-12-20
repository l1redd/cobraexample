package children

import(
	"fmt"
	"github.com/spf13/cobra"
)

type ChildArgs struct {
	one int32
	two int32
	three int32
	four int32
	List []int64
}
var ChildCmd = &cobra.Command{
	Use:   "child",
	Short: "Child command",
	Run: func(cmd *cobra.Command, args []string) {
		ChildRun(cmd)
	},
}

func ChildRun(cmd *cobra.Command) {
	args := ParseChildArgs(cmd)
	fmt.Println("in child, printing child")
	PrintTheThing(args)
}

func PrintTheThing(args ChildArgs) {
	fmt.Printf("Child -- third: %d, fourth: %d \n", args.three, args.four)
}

func init() {
	DefineFlags(ChildCmd)
}
