package main

import (
	"flyawayhub-cli/cmd"
	"fmt"
	"github.com/spf13/cobra"
)

const appName = "flyawayhub"
const version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "flyawayhub-cli",
	Short:   "Flyawayhub CLI application",
	Version: version,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + appName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appName + " version " + version)
	},
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule return your flying reservations.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appName + " version " + version)
	},
}

func init() {
	cmd.InitCommands(rootCmd)
	rootCmd.AddCommand(versionCmd, scheduleCmd)
	//rootCmd.AddCommand(listReservationsCmd)
	//var rootCmd = &cobra.Command{Use: "flyawayhub-cli"}
	//rootCmd.AddCommand(versionCmd)
	//rootCmd.AddCommand(loginCmd)
	//rootCmd.AddCommand(logoutCmd)
	//rootCmd.AddCommand(scheduleCmd)
	/*err := rootCmd.Execute()
	if err != nil {
		return
	}*/
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
