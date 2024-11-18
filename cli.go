package main

import (
	"github.com/spf13/cobra"
)

var (
	// Root command
	// On run -> gospur.
	rootCmd = &cobra.Command{
		Use:   "gospur",
		Short: "Go Spur: Build web applications with Go, without the hassle of JavaScript",
		Long: `Go Spur is a CLI tool that helps you quickly bootstrap Go web applications without worrying about JavaScript.
With Go Spur, you can focus solely on the backend, while it handles the frontend tasks like bundling JavaScript libraries for you.`,
	}

	// Project init command
	// On run -> gospur init.
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a Full-Stack Go Web Project",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
}
