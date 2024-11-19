package config

import "github.com/manifoldco/promptui"

var (
	ErrMsg     = promptui.Styler(promptui.FGRed)
	SuccessMsg = promptui.Styler(promptui.FGGreen)
	NormalMsg  = promptui.Styler(promptui.FGBlack)
)

var (
	WebFrameworkOpts = []string{
		"Echo",
	}
	UILibraryOpts = []string{
		"Preline (requires tailwind)",
	}
	ExtraOpts = []string{
		"Tailwind",
		"HTMX",
	}
)
