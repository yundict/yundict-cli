/*
Copyright © 2024 Yundict
*/
package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project <command>",
		Short: "Authenticate with Yundict",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("auth called")
		},
	}

	// cmdutil.DisableAuthCheck(cmd)

	cmd.AddCommand(NewCmdExport())

	return cmd
}
