package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/nilotpaul/gospur/config"
)

// StackConfig represents a final stack configuration
// based on which project files will be made.
type StackConfig struct {
	// Echo, Fiber, etc...
	WebFramework string

	// UI Library is pre-made styled libs like Preline.
	UILibrary string

	// Extras are extra add-ons like css lib, HTMX etc.
	Extras []string
}

// ProjectPath represents destination or location
// where user want their project to be created.
type ProjectPath struct {
	// FullPath is the absolute path to the project directory.
	FullPath string

	// Path is the relative path to the project directory.
	Path string
}

// GetStackConfig will give a series of prompts
// to the user to configure their project stack.
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

// GetGoModulePath will give a input prompt to the user
// for them to enter a go mod path.
func GetGoModulePath() (string, error) {
	pathPrompt := promptui.Prompt{
		Label: "Enter go mod path (eg. github.com/username/repo)",
		Validate: func(givenPath string) error {
			// Less than 3 character(s).
			if len(givenPath) < 3 {
				return fmt.Errorf("Path cannot be less than 3 character(s)")
			}
			// Starts with https://
			if strings.HasPrefix(givenPath, "https://") {
				return fmt.Errorf("Invalid path '%s', should not contain https", givenPath)
			}
			// Contains any of these -> :*?|
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

// GetProjectPath takes a slice of args (all provided args), validates
// and determines the absolute project path depending on the cwd.
// If no args provided, we fallback to the default set path 'gospur'.
func GetProjectPath(args []string) (*ProjectPath, error) {
	targetPath := "gospur"

	if len(args) > 0 {
		// Santize the given path.
		finalPath, err := ValidateDirPath(args[0])
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

func PrintSuccessMsg(path string) {
	fmt.Println(config.SuccessMsg("\nProject Created! 🎉\n"))
	fmt.Println(config.NormalMsg("Please Run:"))

	// Post installation instructions
	if path == "." {
		fmt.Println(config.FaintMsg(fmt.Sprintf(`
go install github.com/bokwoon95/wgo@latest
go mod tidy
npm install
`)))
	} else {
		fmt.Println(config.FaintMsg(fmt.Sprintf(`
cd %s
go install github.com/bokwoon95/wgo@latest
go mod tidy
npm install
`, path)))
	}

}
