package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/inconshreveable/go-update"
	"github.com/manifoldco/promptui"
	"github.com/nilotpaul/gospur/config"
	"github.com/nilotpaul/gospur/ui"
)

// StackConfig represents a final stack configuration
// based on which project files will be made.
type StackConfig struct {
	// Echo, Fiber, etc...
	WebFramework string

	// CssStrategy can be tailwind, vanilla, etc.
	CssStrategy string
	// UI Library is pre-made styled libs like Preline.
	UILibrary string

	// RenderingStrategy defines how HTML is rendered.
	// Eg. templates, templ, seperate client.
	RenderingStrategy string

	// Flags Only
	// Extras are extra add-ons like css lib, HTMX etc.
	ExtraOpts []string
}

// ProjectPath represents destination or location
// where user want their project to be created.
type ProjectPath struct {
	// FullPath is the absolute path to the project directory.
	FullPath string

	// Path is the relative path to the project directory.
	Path string
}

type GitHubReleaseResponse struct {
	Version string `json:"tag_name"`
	Assets  []struct {
		ID                 int64  `json:"id"`
		Name               string `json:"name"`
		Size               int64  `json:"size"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

// GetStackConfig will give a series of prompts
// to the user to configure their project stack.
func GetStackConfig(cfg *StackConfig) error {
	// Framework options
	if len(cfg.WebFramework) == 0 {
		frameworkPrompt := promptui.Select{
			Label: "Choose a web framework",
			Items: config.WebFrameworkOpts,
		}
		_, framework, err := frameworkPrompt.Run()
		if err != nil {
			return fmt.Errorf("failed to select web framework")
		}
		cfg.WebFramework = framework
	}
	// CSS Strategy
	if len(cfg.CssStrategy) == 0 {
		extraPrompt := promptui.Select{
			Label: "Choose a CSS Strategy",
			Items: config.CssStrategyOpts,
		}
		_, css, err := extraPrompt.Run()
		if err != nil {
			return fmt.Errorf("failed to select CSS Strategy")
		}
		cfg.CssStrategy = css
	}
	// UI Library Options
	if len(cfg.UILibrary) == 0 {
		// Filtering the opts for UI Libs based on the css strategy chosen.
		filteredOpts := make([]string, 0)
		for lib, deps := range config.UILibraryOpts {
			if len(deps) == 0 {
				filteredOpts = append(filteredOpts, lib)
				continue
			}
			if contains(deps, cfg.CssStrategy) {
				filteredOpts = append(filteredOpts, lib)
			}
		}

		// Only ask anything if we have a compatible UI Lib for
		// the chosen CSS Strategy.
		if len(filteredOpts) != 0 {
			// Asking for UI Lib if we've any filtered opts.
			uiLibPrompt := promptui.Select{
				Label: "Choose a UI Library",
				Items: filteredOpts,
			}
			_, uiLib, err := uiLibPrompt.Run()
			if err != nil {
				return fmt.Errorf("failed to select UI Library")
			}
			cfg.UILibrary = uiLib
		}
	}
	// Rendering Strategy Options
	if len(cfg.RenderingStrategy) == 0 {
		renderingStratPrompt := promptui.Select{
			Label: "Choose a Rendering Strategy",
			Items: config.RenderingStrategy,
		}
		_, opts, err := renderingStratPrompt.Run()
		if err != nil {
			return fmt.Errorf("failed to select Rendering Strategy")
		}
		cfg.RenderingStrategy = opts
	}

	return nil
}

// GetGoModulePath will give a input prompt to the user
// for them to enter a go mod path.
func GetGoModulePath() (string, error) {
	pathPrompt := promptui.Prompt{
		Label:    "Enter go mod path (eg. github.com/username/repo)",
		Validate: validateGoModPath,
	}
	path, err := pathPrompt.Run()
	if err != nil {
		return "", fmt.Errorf("error getting the mod path %v", err)
	}

	return path, nil
}

// RunGoModInit takes the full project path and a name.
// It changes the cwd to the given path and run go mod init
// with the given name.
func RunGoModInit(fullProjectPath, name string) error {
	// Change the current working directory to the project directory
	if err := os.Chdir(fullProjectPath); err != nil {
		return fmt.Errorf("failed to change to project directory: %v", err)
	}

	cmd := exec.Command("go", "mod", "init", name)
	return cmd.Run()
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
		return nil, fmt.Errorf("error getting the current working directory %v", err)
	}

	fullPath := filepath.Join(cwd, targetPath)

	return &ProjectPath{FullPath: fullPath, Path: targetPath}, nil
}

func FetchRelease(ctx context.Context, v ...string) (GitHubReleaseResponse, error) {
	var (
		data         GitHubReleaseResponse
		givenVersion = "latest"
		releaseUrl   = fmt.Sprintf(config.GitHubReleaseAPIURL+"/%s", givenVersion)
	)
	if len(v) > 0 && v[0] != "latest" {
		releaseUrl = fmt.Sprintf(config.GitHubReleaseAPIURL+"/tags/%s", v[0])
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, releaseUrl, nil)
	if err != nil {
		return data, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return data, fmt.Errorf("request took too long or canceled")
		}

		return data, err
	}
	if res.StatusCode == http.StatusNotFound {
		return data, fmt.Errorf("version not found")
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

// HandleUpdateCLI updates the cli to a new version, handling the pending
// states and clean up.
func HandleUpdateCLI(url string, exePath string) error {
	var (
		errChan = make(chan error, 1)
		s       = ui.NewSpinner("updating...")
	)

	s.Start()
	defer func() {
		close(errChan)
		s.Stop()
		fmt.Printf("%s", "\n")
	}()

	go func() {
		errChan <- doUpdate(url, exePath)
	}()

	err := <-errChan
	return err
}

// HandleGetRelease handles gets the latest release, handles pending states and clean up.
func HandleGetRelease(ctx context.Context) (GitHubReleaseResponse, error) {
	var (
		releaseChan = make(chan GitHubReleaseResponse, 1)
		errChan     = make(chan error, 1)

		s = ui.NewSpinner("getting the latest version...")
	)

	s.Start()
	defer func() {
		close(releaseChan)
		close(errChan)
		s.Stop()
		fmt.Printf("%s", "\n")
	}()

	go func() {
		// Fetch the latest release from github.
		release, err := FetchRelease(ctx, "latest")

		releaseChan <- release
		errChan <- err
	}()

	release := <-releaseChan
	err := <-errChan

	return release, err
}

func PrintSuccessMsg(path string) {
	fmt.Println(config.SuccessMsg("\nProject Created! 🎉\n"))
	fmt.Println(config.NormalMsg("Please Run:"))

	// Post installation instructions
	if path == "." {
		fmt.Println(config.FaintMsg(`
go install github.com/bokwoon95/wgo@latest
go mod tidy
npm install
`))
	} else {
		fmt.Println(config.FaintMsg(fmt.Sprintf(`
cd %s
go install github.com/bokwoon95/wgo@latest
go mod tidy
npm install
`, path)))
	}

}

func doUpdate(url string, targetPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	binary, err := uncompress(resp.Body, url)
	if err != nil {
		return err
	}

	return update.Apply(binary, update.Options{
		TargetPath: targetPath,
		TargetMode: os.ModePerm,
	})
}

func validateGoModPath(path string) error {
	if len(path) < 3 {
		return fmt.Errorf("path cannot be less than 3 character(s)")
	}
	// Starts with https://
	if strings.HasPrefix(path, "https://") {
		return fmt.Errorf("invalid path '%s', should not contain https", path)
	}
	// Contains any of these -> :*?|
	if strings.ContainsAny(path, " :*?|") {
		return fmt.Errorf("invalid path '%s', contains reserved characters", path)
	}
	// Length exceedes 255 character(s)
	if len(path) > 255 {
		return fmt.Errorf("exceeded maximum length")
	}

	return nil
}
