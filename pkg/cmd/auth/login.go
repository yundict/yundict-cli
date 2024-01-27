/*
Copyright © 2024 Yundict
*/
package auth

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func NewCmdLogin() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "login",
		Args:  cobra.ExactArgs(0),
		Short: "Log in to a Yundict account",
		Long: heredoc.Docf(`
			Authenticate with a Yundict host.

			The default authentication mode is a web-based browser flow. After completion, an
			authentication token will be stored securely in the system credential store.
			If a credential store is not found or there is an issue using it gh will fallback
			to writing the token to a plain text file. See %[1]sgh auth status%[1]s for its
			stored location.

			Alternatively, use %[1]s--with-token%[1]s to pass in a token on standard input.
			The minimum required scopes for the token are: %[1]srepo%[1]s, %[1]sread:org%[1]s, and %[1]sgist%[1]s.

			Alternatively, gh will use the authentication token found in environment variables.
			This method is most suitable for "headless" use of gh such as in automation. See
			%[1]sgh help environment%[1]s for more info.

			To use gh in Yundict Actions, add %[1]sGH_TOKEN: ${{ github.token }}%[1]s to %[1]senv%[1]s.

			The git protocol to use for git operations on this host can be set with %[1]s--git-protocol%[1]s,
			or during the interactive prompting. Although login is for a single account on a host, setting
			the git protocol will take effect for all users on the host.
		`, "`"),
		Example: heredoc.Doc(`
			# Start interactive setup
			$ gh auth login

			# Authenticate against github.com by reading the token from a file
			$ gh auth login --with-token < mytoken.txt

			# Authenticate with specific host
			$ gh auth login --hostname enterprise.internal
		`),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("login called")
		},
	}

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
