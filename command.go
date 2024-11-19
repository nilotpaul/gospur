package main

import (
	"fmt"
	"os"

	"github.com/nilotpaul/gospur/config"
	"github.com/nilotpaul/gospur/util"
	"github.com/spf13/cobra"
)

func handleInitCmd(cmd *cobra.Command, args []string) {
	// isEarlyStage depicts the CLI is still in early stages
	// and we don't have enough options for user prompts.
	isEarlyStage := os.Getenv("EARLY_STAGE") == "True"

	targetPath, err := util.GetProjectPath(args)
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	// Based isEarlyStage we are building the `StackConfig`
	// If isEarlyStage is not true, we use start taking user prompts.
	// If isEarlyStage is true, we use a default config.
	var cfg util.StackConfig
	if !isEarlyStage {
		stackCfg, err := util.GetStackConfig()
		if err != nil {
			fmt.Println(config.ErrMsg(err))
			return
		}
		cfg = *stackCfg
	} else {
		cfg = util.StackConfig{
			WebFramework: config.WebFrameworkOpts[0],
			UILibrary:    config.UILibraryOpts[0],
			Extras:       config.ExtraOpts,
		}
		fmt.Println(config.NormalMsg("The CLI is still WIP 🚧, you'll get a default stack for now."))
	}

	// Asking for the go mod path from user.
	goModPath, err := util.GetGoModulePath()
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	fmt.Println("Go mod path: ", goModPath)
	fmt.Println("Stack config: ", cfg)
	fmt.Println("Final resolved path: ", targetPath.FullPath)
}
