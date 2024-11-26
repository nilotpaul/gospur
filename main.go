package main

import (
	"fmt"
	"os"

	"github.com/nilotpaul/gospur/config"
)

func main() {
	if err := Execute(); err != nil {
		fmt.Println(config.ErrMsg("Go Spur exited"))
		os.Exit(1)
	}
}
