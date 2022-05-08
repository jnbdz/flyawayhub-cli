package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	appName = "Flyawayhub"
	// Used for flags.
	cfgFile string

	loginCmd = &cobra.Command{
		Use: "login",
		Short: "Login will sign you in " + appName + " (will generate an authorization bearer).",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(appName + " version " + version)
		},
	}

	logoutCmd = &cobra.Command{
		Use: "logout",
		Short: "Logout will informe " + appName + " to end your session (will remove the authorization bearer).",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(appName + " version " + version)
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)


}