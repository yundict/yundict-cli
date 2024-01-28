package util

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func CheckAuth(token string) bool {
	// 配置中有token
	// if cfg.Authentication().HasEnvToken() {
	// 	return true
	// }

	// 参数中有-t
	if token == "" {
		return false
	}

	return true
}

func IsAuthCheckEnabled(cmd *cobra.Command) bool {
	switch cmd.Name() {
	case "help", cobra.ShellCompRequestCmd, cobra.ShellCompNoDescRequestCmd:
		return false
	}

	for c := cmd; c.Parent() != nil; c = c.Parent() {
		if c.Annotations != nil && c.Annotations["skipAuthCheck"] == "true" {
			return false
		}
	}

	return true
}

func AuthHelp() string {
	return heredoc.Doc(`
		To get started with Yundict CLI, token parameter (-t, --token) must be provided.
		Alternatively, you can run: yundict auth login first to bind the token.
	`)
}
