package children

import (
	"fmt"
	"github.com/spf13/cobra"
)

func DefineFlags(cmd *cobra.Command) {
	cmd.Flags().Int32P("three", "r", 0, "third arg")
	cmd.Flags().Int32P("four", "f", 0, "fourth arg")
	val := make([]int64, 0)
	cmd.Flags().Int64SliceP("list", "l", val, "pass in a list of ints")
}

func ParseChildArgs(cmd *cobra.Command) (args ChildArgs) {
	third, err := cmd.Flags().GetInt32("three")
	if err != nil {
		//fmt.Printf("error parsing three: %s \n", err.Error())
		return
	}

	fmt.Printf("three: %d", third)

	fourth, err := cmd.Flags().GetInt32("four")
	if err != nil {
		//fmt.Printf("error parsing four: %s \n", err.Error())
		return
	}

	fmt.Printf("four: %d \n", fourth)

	list, err := cmd.Flags().GetInt64Slice("list")
	if err != nil {
		return
	}

	return ChildArgs{
		three: third,
		four: fourth,
		List: list,
	}
}


