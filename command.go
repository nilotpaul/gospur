package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

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
		fmt.Println("Version:", config.ErrMsg(err))
		return
	}

	fmt.Println("Version:", config.NormalMsg(version))
}

// handleUpdateCmd handles the update command for gospur CLI.
func handleUpdateCmd(cmd *cobra.Command, args []string) {
	currVersion, err := config.GetVersion()
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	// Gets the currently running binary location.
	installedExePath, err := os.Executable()
	if err != nil {
		fmt.Println(config.ErrMsg("Failed to get the installation path: " + err.Error()))
		return
	}

	// Cancel fetch to github releases if took more than 2 seconds.
	ctx, cancel := context.WithDeadline(cmd.Context(), time.Now().Add(2*time.Second))
	defer cancel()

	// Fetiching latest release from github api
	release, err := util.HandleGetRelease(ctx)
	if err != nil {
		fmt.Println(config.ErrMsg(err))
		return
	}

	if currVersion == release.Version {
		fmt.Println(config.ErrMsg("Latest version is already installed"))
		return
	}

	var binaries []string
	for _, asset := range release.Assets {
		binaries = append(binaries, asset.BrowserDownloadURL)
	}

	// Finding the compatible binary from currently installed one.
	targetBinaryUrl := util.FindMatchingBinary(binaries, runtime.GOOS, runtime.GOARCH)
	if err := util.HandleUpdateCLI(targetBinaryUrl, installedExePath); err != nil {
		fmt.Printf("Update failed: %v\n", config.ErrMsg(err))
		return
	}

	fmt.Printf("CLI has been updated to the latest version (%s)\n", config.SuccessMsg(release.Version))
}
