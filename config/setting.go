package config

import (
	"fmt"
	"runtime/debug"

	"github.com/manifoldco/promptui"
)

const logo string = `
   _____       _____                  
  / ____|     / ____|                 
 | |  __  ___| (___  _ __  _   _ _ __ 
 | | |_ |/ _ \\___ \| '_ \| | | | '__|
 | |__| | (_) |___) | |_) | |_| | |   
  \_____|\___/_____/| .__/ \__,_|_|   
                    | |               
                    |_|               
`

const (
	WinBinaryName   = "gospur.exe"
	OtherBinaryName = "gospur" // for linux and darwin
)

var LogoColoured string = promptui.Styler(promptui.FGCyan, promptui.FGBold)(logo)

// GoSpur CLI version info
var (
	version string
	commit  string
	date    string
)

func GetVersion() (string, error) {
	noInfoErr := fmt.Errorf("No version information available")

	// goreleaser has embeded the version via ldflags.
	if len(version) != 0 {
		return version, nil
	}

	// Try to get the version from the go.mod build info.
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", noInfoErr
	}
	if info.Main.Version != "(devel)" {
		return info.Main.Version, nil
	}

	return "", noInfoErr
}

func GetSafeVersion() string {
	version, _ := GetVersion()
	return NormalMsg(version)
}
