package project

import (
	"github.com/spf13/cobra"

	"github.com/yundict/yundict-go"
)

// var exportCmd = &cobra.Command{
// 	Use:   "export",
// 	Short: "Export key & translations to a file",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		token := cmd.Flags().Lookup("token").Value.String()
// 		teamName, _ := cmd.Flags().GetString("team")
// 		projectName, _ := cmd.Flags().GetString("project")
// 		export(token, teamName, projectName)
// 	},
// }

var (
	Token string
)

func NewCmdExport() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export key & translations to a file",
		Run: func(cmd *cobra.Command, args []string) {
			token := cmd.Flags().Lookup("token").Value.String()
			teamName, _ := cmd.Flags().GetString("team")
			projectName, _ := cmd.Flags().GetString("project")
			export(token, teamName, projectName)
		},
	}

	cmd.Flags().StringVarP(&Token, "token", "t", "", "Yundict API Token")
	cmd.Flags().String("team", "", "Yundict team name")
	cmd.Flags().String("project", "", "Yundict project name")
	cmd.MarkFlagRequired("token")
	cmd.MarkFlagRequired("team")
	cmd.MarkFlagRequired("project")

	// cmd.Flags().StringVarP(&opts.Hostname, "hostname", "h", "", "The hostname of the Yundict instance to authenticate with")
	// cmd.Flags().StringSliceVarP(&opts.Scopes, "scopes", "s", nil, "Additional authentication scopes to request")
	// cmd.Flags().BoolVar(&tokenStdin, "with-token", false, "Read token from standard input")
	// cmd.Flags().BoolVarP(&opts.Web, "web", "w", false, "Open a browser to authenticate")
	// cmdutil.StringEnumFlag(cmd, &opts.GitProtocol, "git-protocol", "p", "", []string{"ssh", "https"}, "The protocol to use for git operations on this host")

	// // secure storage became the default on 2023/4/04; this flag is left as a no-op for backwards compatibility
	// var secureStorage bool
	// cmd.Flags().BoolVar(&secureStorage, "secure-storage", false, "Save authentication credentials in secure credential store")
	// _ = cmd.Flags().MarkHidden("secure-storage")

	// cmd.Flags().BoolVar(&opts.InsecureStorage, "insecure-storage", false, "Save authentication credentials in plain text instead of credential store")

	return cmd
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
