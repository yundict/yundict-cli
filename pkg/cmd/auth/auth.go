/*
Copyright © 2024 Yundict
*/
package auth

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth <command>",
		Short: "Authenticate with Yundict",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("auth called")
		},
	}

	// cmdutil.DisableAuthCheck(cmd)

	cmd.AddCommand(NewCmdLogin())

	return cmd
}
