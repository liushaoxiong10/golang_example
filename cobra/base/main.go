package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var version bool
	var cmd = &cobra.Command{
		Use:   "root [sub]",
		Short: "rout command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("args: %v\n", args)
			if version {
				fmt.Println("version: 1.0")
			}
		},
	}
	flags := cmd.Flags()
	flags.BoolVarP(&version, "version", "v", false, "echo version")
	// 调用Run
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

}
