package main

import (
	"fmt"
	"os"

	"github.com/nilotpaul/gospur/config"
	"github.com/nilotpaul/gospur/util"
	"github.com/spf13/cobra"
)

func handleInitCmd(cmd *cobra.Command, args []string) {
	isEarlyStage := os.Getenv("EARLY_STAGE") == "True"

	targetPath, err := util.GetProjectPath(args)
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

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
	}

	fmt.Println("Stack config: ", cfg)
	fmt.Println("Final resolved path: ", targetPath.FullPath)
}
