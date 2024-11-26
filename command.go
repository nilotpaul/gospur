package main

import (
	"fmt"

	"github.com/nilotpaul/gospur/config"
	"github.com/nilotpaul/gospur/util"
	"github.com/spf13/cobra"
)

// handleInitCmd handles the `init` command for gospur CLI.
func handleInitCmd(cmd *cobra.Command, args []string) {
	targetPath, err := util.GetProjectPath(args)
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	// Building the stack config by talking user prompts.
	stackCfg, err := util.GetStackConfig()
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}
	cfg := *stackCfg

	// Asking for the go mod path from user.
	goModPath, err := util.GetGoModulePath()
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	// Creating the target project directory.
	// It'll check if the dir already exist and is empty or not (strict).
	if err := util.CreateTargetDir(targetPath.Path, true); err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	// Creating the project files in the target directory.
	// Passing the go mod path for resolving Go imports.
	err = util.CreateProject(
		targetPath.Path,
		cfg,
		util.MakeProjectCtx(cfg, goModPath),
	)
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	// Running `go mod init` with the specified name.
	if err := util.RunGoModInit(targetPath.FullPath, goModPath); err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	util.PrintSuccessMsg(targetPath.Path)
}

// handleVersionCmd handles the `version` command for gospur CLI.
func handleVersionCmd(cmd *cobra.Command, args []string) {
	version, err := config.GetVersion()
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	fmt.Println(config.NormalMsg("version: " + version))
}
