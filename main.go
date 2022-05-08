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

func init() {
	var rootCmd = &cobra.Command{Use: "flyawayhub-cli"}
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}

func Hello() string {
	return "Hello World!"
}

func main() {
	fmt.Println(Hello())
}
