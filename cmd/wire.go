package cmd

import (
	"flyawayhub-cli/logging"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

func InitializeLoginCommand() *cobra.Command {
	wire.Build(
		logging.NewZapLogger, // Provides logging.Logger
		NewLoginCommand,      // Constructor for *cobra.Command
	)
	return nil // Wire fills this in
}
