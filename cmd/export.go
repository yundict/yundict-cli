package cmd

import (
	"github.com/spf13/cobra"

	"github.com/yundict/yundict-go"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export key & translations to a file",
	Run: func(cmd *cobra.Command, args []string) {
		token := cmd.Flags().Lookup("token").Value.String()
		teamName, _ := cmd.Flags().GetString("team")
		projectName, _ := cmd.Flags().GetString("project")
		export(token, teamName, projectName)
	},
}

func export(token string, teamName string, projectName string) {
	// Make HTTP Request
	client := yundict.NewClient(token)

	// export all keys & translations
	res, err := client.Keys.Export(teamName, projectName, "json")
	if err != nil {
		panic(err)
	}

	// Print result
	println(res.Data)
}
