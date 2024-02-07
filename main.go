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

var organizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "Fetch and display organization information",
	RunE: func(c *cobra.Command, args []string) error {
		sessionData, err := cmd.LoadSession()
		if err != nil {
			return fmt.Errorf("loading session: %w", err)
		}

		return cmd.FetchOrganizationInfo(sessionData.AccessToken)
	},
}

var reservationsCmd = &cobra.Command{
	Use:   "reservations",
	Short: "Fetch flying reservations for your organization",
	Run: func(c *cobra.Command, args []string) {
		cmd.HandleReservationsCommand() // Use cmd. to reference the function from the cmd package
	},
}

var membersCmd = &cobra.Command{
	Use:   "members",
	Short: "Fetch flying members for your organization",
	Run: func(c *cobra.Command, args []string) {
		cmd.HandleMembersCommand() // Use cmd. to reference the function from the cmd package
	},
}

var sunInfoCmd = &cobra.Command{
	Use:   "suninfo",
	Short: "Fetch flying sun info for your organization",
	Run: func(c *cobra.Command, args []string) {
		cmd.HandleSunInfoCommand() // Use cmd. to reference the function from the cmd package
	},
}

var schedulesCmd = &cobra.Command{
	Use:   "schedules",
	Short: "Fetch flying schedules for your organization",
	Run: func(c *cobra.Command, args []string) {
		cmd.HandleSchedulesCommand() // Use cmd. to reference the function from the cmd package
	},
}

var notificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "Fetch flying notifications for your organization",
	Run: func(c *cobra.Command, args []string) {
		cmd.HandleNotificationsCommand() // Use cmd. to reference the function from the cmd package
	},
}

func init() {
	cmd.InitCommands(rootCmd)
	rootCmd.AddCommand(versionCmd, scheduleCmd, organizationsCmd, reservationsCmd, membersCmd, sunInfoCmd, schedulesCmd, notificationsCmd)
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
