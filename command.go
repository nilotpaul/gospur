package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nilotpaul/gospur/util"
	"github.com/spf13/cobra"
)

func handleInitCmd(cmd *cobra.Command, args []string) {
	targetPath := "gospur"

	if len(args) > 0 {
		// Santize the given path.
		finalPath, err := util.SanitizeDirPath(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		// Now it's safe to use the `targetPath`.
		targetPath = finalPath
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the current working directory: ", err)
		return
	}

	fullPath := filepath.Join(cwd, targetPath)
	fmt.Println("Final resolved path: ", fullPath)
}
