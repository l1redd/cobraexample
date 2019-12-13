package main

import (
	"fmt"
	"github.com/l1redd/cobraexample/cmd"
	"github.com/l1redd/cobraexample/cmd/children"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use: "rootCmd",
	}

	rootCmd.AddCommand(cmd.ParentCmd)
	rootCmd.AddCommand(cmd.Parent2Cmd)
	rootCmd.AddCommand(children.ChildCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("errorrrrr: %s \n", err.Error())
	}

}