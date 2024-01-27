module github.com/yundict/yundict-cli

go 1.21.5

require (
	github.com/spf13/cobra v1.8.0
	github.com/yundict/yundict-go v0.0.0-20240106131841-c3e26932b76c
)

require (
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/yundict/yundict-cli/pkg/cmd/auth => ../pkg/cmd/auth
