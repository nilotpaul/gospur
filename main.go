package main

import (
	"fmt"
	"os"

	"github.com/nilotpaul/gospur/config"
)

func main() {
	// This sets a `EARLY_STAGE` env which we use to
	// initialize a default project without any options.
	// This is done as currently we don't have any extra
	// options other than the default template.
	os.Setenv("EARLY_STAGE", "True")

	// Stars the CLI
	if err := Execute(); err != nil {
		fmt.Println(config.NormalMsg("Go Spur exited"))
		os.Exit(1)
	}
}
