package main

import (
	"fmt"

	"github.com/nilotpaul/gospur/util"
	"github.com/spf13/cobra"
)

func handleInitCmd(cmd *cobra.Command, args []string) {
	targetPath, err := util.GetProjectPath(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Final resolved path: ", targetPath.FullPath)
}
