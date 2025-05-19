package main

import (
	"fmt"

	"github.com/nilotpaul/gospur/config"
	"github.com/spf13/cobra"
)

var (
	// Root command
	// On run -> gospur.
	rootCmd = &cobra.Command{
		Use:     "gospur",
		Short:   "Go Spur: Build web applications with Go, without the hassle of JavaScript",
		Long:    "Go Spur is a CLI tool that helps you quickly bootstrap Go web applications without worrying about JavaScript. Focus solely on the backend, while we handle the small repetitive tasks for you.",
		Version: config.GetSafeVersion(),
	}

	// Project init command
	// On run -> gospur init.
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a Full-Stack Go Web Project",
		Args:  cobra.MaximumNArgs(1),
		Run:   handleInitCmd,
	}

	// Project Update CLI command
	// On run -> gospur update (latest).
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates the CLI to the latest version",
		Args:  cobra.NoArgs,
		Run:   handleUpdateCmd,
	}

	// Project version command
	// On run -> gospur version.
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Shows the current installed version",
		Args:  cobra.NoArgs,
		Run:   handleVersionCmd,
	}
)

func Execute() error {
	fmt.Println(config.LogoColoured)
	return rootCmd.Execute()
}

func init() {
	// Flags for init cmd.
	registerInitCmdFlags()

	rootCmd.AddCommand(
		initCmd,
		updateCmd,
		versionCmd,
	)

}
