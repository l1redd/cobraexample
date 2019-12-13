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
	rootCmd.AddCommand(cmd.Parent2Cmd)
	//childCmd := cmd.NewChildCmd()
	//rootCmd.AddCommand(childCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("errorrrrr: %s \n", err.Error())
	}

}