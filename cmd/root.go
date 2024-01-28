package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	authCmd "github.com/yundict/yundict-cli/pkg/cmd/auth"
	projectCmd "github.com/yundict/yundict-cli/pkg/cmd/project"
	"github.com/yundict/yundict-cli/pkg/util"
)

const (
	Version = "0.0.1"
)

var (
	Token string
)

var rootCmd = &cobra.Command{
	Use:     "yundict",
	Short:   "Yundict CLI v" + Version,
	Version: Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// check auth token
		if util.IsAuthCheckEnabled(cmd) && !util.CheckAuth(Token) {
			cmd.SilenceUsage = true
			fmt.Println(util.AuthHelp())
			return errors.New("Auth Error")
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type exitCode int

const (
	exitOK      exitCode = 0
	exitError   exitCode = 1
	exitCancel  exitCode = 2
	exitAuth    exitCode = 4
	exitPending exitCode = 8
)

func init() {
	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "Yundict API Token")

	rootCmd.AddCommand(authCmd.NewCmdAuth())
	rootCmd.AddCommand(projectCmd.NewCmdProject())
}
