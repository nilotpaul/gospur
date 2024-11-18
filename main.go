package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Execute(); err != nil {
		fmt.Println("Go Spur exited with an error: ", err)
		os.Exit(1)
	}
}
