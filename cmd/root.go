package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
	exportCmd.Flags().String("type", "", "Export file type, must be one of the following values: key-value-json, csv, android-xml, apple-strings, arb, properties, ini, i18next-json, excel-xlsx, yaml")
	exportCmd.Flags().String("languages", "", "Export languages, comma separated, e.g. en,zh-CN (If not provided, all languages will be exported)")
	exportCmd.Flags().String("out", "", "Output file path")
	exportCmd.MarkFlagRequired("team")
	exportCmd.MarkFlagRequired("project")
	exportCmd.MarkFlagRequired("type")
	exportCmd.MarkFlagRequired("out")
}
