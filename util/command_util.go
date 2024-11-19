package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
		return nil, fmt.Errorf("Failed to select web framework %v", err)
	}
	cfg.WebFramework = framework

	// UI Library Options
	uiLibPrompt := promptui.Select{
		Label: "Choose a UI Library",
		Items: config.UILibraryOpts,
	}
	_, uiLib, err := uiLibPrompt.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to select web framework %v", err)
	}
	cfg.UILibrary = uiLib

	// Extra Add-Ons
	extrasChosen := make([]string, 0)
	for _, extra := range config.ExtraOpts {
		extraPrompt := promptui.Select{
			Label: "Add " + extra,
			Items: []string{"No", "Yes"},
		}

		// If Preline which depends on tailwind is selected as a UI Lib, we skip the
		// current iteration and add tailwind in `extras` by default.
		if extra == "Tailwind" && cfg.UILibrary == "Preline (requires tailwind)" {
			extrasChosen = append(extrasChosen, "Tailwind")
			continue
		}

		_, choice, err := extraPrompt.Run()
		if err != nil {
			return nil, fmt.Errorf("Failed to select extras %v", err)
		}
		if choice == "Yes" {
			extrasChosen = append(extrasChosen, extra)
		}
	}
	cfg.Extras = extrasChosen

	return &cfg, nil
}

func GetGoModulePath() (string, error) {
	pathPrompt := promptui.Prompt{
		Label: "Enter go mod path (eg. github.com/username/repo)",
		Validate: func(givenPath string) error {
			if len(givenPath) < 3 {
				return fmt.Errorf("Path cannot be less than 3 character(s)")
			}
			if strings.HasPrefix(givenPath, "https://") {
				return fmt.Errorf("Invalid path '%s', should not contain https", givenPath)
			}
			if strings.ContainsAny(givenPath, " :*?|") {
				return fmt.Errorf("Invalid path '%s', contains reserved characters", givenPath)
			}
			return nil
		},
	}
	path, err := pathPrompt.Run()
	if err != nil {
		return "", fmt.Errorf("Error getting the mod path %v", err)
	}

	return path, nil
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
		return nil, fmt.Errorf("Error getting the current working directory %v", err)
	}

	fullPath := filepath.Join(cwd, targetPath)

	return &ProjectPath{FullPath: fullPath, Path: targetPath}, nil
}
