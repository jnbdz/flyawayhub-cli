package cmd

import (
	"flyawayhub-cli/auth"
	"flyawayhub-cli/logging"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

func InitializeLoginCommand() *cobra.Command {
	wire.Build(
		auth.NewAuthService,
		logging.NewZapLogger, // Provides logging.Logger
		NewLoginCommand,      // Constructor for *cobra.Command
	)
	return nil // Wire fills this in
}
