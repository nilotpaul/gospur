package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetProjectPath(args []string) (string, error) {
	targetPath := "gospur"

	if len(args) > 0 {
		// Santize the given path.
		finalPath, err := SanitizeDirPath(args[0])
		if err != nil {

			return "", err
		}
		// Now it's safe to use the `targetPath`.
		targetPath = finalPath
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Error getting the current working directory: ", err)
	}

	fullPath := filepath.Join(cwd, targetPath)

	return fullPath, nil
}
