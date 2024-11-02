package cmd

import (
	"strings"

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
		exportType, _ := cmd.Flags().GetString("type")
		languages, _ := cmd.Flags().GetString("languages")
		outPath, _ := cmd.Flags().GetString("out")
		export(token, teamName, projectName, exportType, languages, outPath)
	},
}

func export(token string, teamName string, projectName string, exportType string, languages string, outPath string) {
	// Make HTTP Request
	client := yundict.NewClient(token)

	// Parse languages
	langs := strings.Split(languages, ",")

	// export all keys & translations
	res, err := client.Keys.Export(teamName, projectName, exportType, langs)
	if err != nil {
		panic(err)
	}

	fileUrl := res.Data

	// check fileUrl format
	if !strings.HasPrefix(fileUrl, "http") {
		panic("Invalid output file url!")
	}

	// Download file
	downloadFile(fileUrl, outPath)
}
