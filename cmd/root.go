package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	authCmd "github.com/yundict/yundict-cli/pkg/cmd/auth"
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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(exportCmd)
	rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "Yundict API Token")
	rootCmd.MarkPersistentFlagRequired("token")

	// export command
	exportCmd.Flags().String("team", "", "Yundict team name")
	exportCmd.Flags().String("project", "", "Yundict project name")
	exportCmd.MarkFlagRequired("team")
	exportCmd.MarkFlagRequired("project")

	rootCmd.AddCommand(authCmd.NewCmdAuth())
}
