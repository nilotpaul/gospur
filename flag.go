package main

import (
	"fmt"
	"strings"

	"github.com/nilotpaul/gospur/config"
	"github.com/nilotpaul/gospur/util"
)

func registerInitCmdFlags() {
	initCmd.Flags().StringVar(
		&stackConfig.WebFramework, "framework", "",
		strings.Join(config.WebFrameworkOpts, ", "),
	)
	initCmd.Flags().StringVar(
		&stackConfig.CssStrategy, "styling", "",
		strings.Join(config.CssStrategyOpts, ", "),
	)
	initCmd.Flags().StringVar(
		&stackConfig.UILibrary, "ui", "",
		strings.Join(util.GetMapKeys(config.UILibraryOpts), ", "),
	)
	initCmd.Flags().StringVar(
		&stackConfig.RenderingStrategy, "render", "",
		strings.Join(util.GetRenderingOpts(true), ", "),
	)
	initCmd.Flags().StringSliceVar(
		&stackConfig.ExtraOpts, "extra", []string{},
		fmt.Sprintf("One or Many: %s", strings.Join(config.ExtraOpts, ", ")),
	)
}
