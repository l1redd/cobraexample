package main

import (
	"fmt"
	"github.com/l1redd/cobraexample/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use: "rootCmd",
	}

	rootCmd.AddCommand(cmd.ParentCmd)
	rootCmd.AddCommand(cmd.ChildCommand)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Print("errorrrrr")
	}

}