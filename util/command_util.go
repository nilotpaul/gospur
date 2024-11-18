package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/nilotpaul/gospur/config"
)

type StackConfig struct {
	WebFramework string
	UILibrary    string

	// Extras are extra add-ons like css lib, HTMX etc.
	Extras []string
}

type ProjectPath struct {
	FullPath string
	Path     string
}

func GetStackConfig() (*StackConfig, error) {
	var cfg StackConfig

	// Framework options
	frameworkPrompt := promptui.Select{
		Label: "Choose a web framework",
		Items: config.WebFrameworkOpts,
	}
	_, framework, err := frameworkPrompt.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to select web framework: %v", err)
	}
	cfg.WebFramework = framework

	// UI Library Options
	uiLibPrompt := promptui.Select{
		Label: "Choose a UI Library",
		Items: config.UILibraryOpts,
	}
	_, uiLib, err := uiLibPrompt.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to select web framework: %v", err)
	}
	cfg.UILibrary = uiLib

	// Extra Add-Ons
	extrasChosen := make([]string, 0)
	for _, extra := range config.ExtraOpts {
		extraPrompt := promptui.Select{
			Label: "Add " + extra,
			Items: []string{"No", "Yes"},
		}
		_, choice, err := extraPrompt.Run()
		if err != nil {
			return nil, fmt.Errorf("Failed to select extras: %w", err)
		}
		if choice == "Yes" {
			extrasChosen = append(extrasChosen, extra)
		}
	}
	cfg.Extras = extrasChosen

	return &cfg, nil
}

func GetProjectPath(args []string) (*ProjectPath, error) {
	targetPath := "gospur"

	if len(args) > 0 {
		// Santize the given path.
		finalPath, err := SanitizeDirPath(args[0])
		if err != nil {

			return nil, err
		}
		// Now it's safe to use the `targetPath`.
		targetPath = finalPath
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Error getting the current working directory: %v", err)
	}

	fullPath := filepath.Join(cwd, targetPath)

	return &ProjectPath{FullPath: fullPath, Path: targetPath}, nil
}
