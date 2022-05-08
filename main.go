package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const appName = "flyawayhub"
const version = "1.0.0"

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of " + appName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appName + " version " + version)
	},
}

var scheduleCmd = &cobra.Command{
	Use: "schedule",
	Short: "Schedule return your flying reservations.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appName + " version " + version)
	},
}

func init() {
	var rootCmd = &cobra.Command{Use: "flyawayhub-cli"}
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(scheduleCmd)
	rootCmd.Execute()
}

func login() {

}

func Hello() string {
	return "Hello World!"
}

func main() {
	fmt.Println(Hello())
}
