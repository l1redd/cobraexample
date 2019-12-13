package children

import (
	"fmt"
	"github.com/spf13/cobra"
)

func DefineFlags(cmd *cobra.Command) {
	cmd.Flags().Int32P("three", "r", 0, "third arg")
	cmd.Flags().Int32P("four", "f", 0, "fourth arg")
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

	return ChildArgs{
		three: third,
		four: fourth,
	}
}


