package util

import (
	"fmt"
	"os"
	"path/filepath"
)

type ProjectPath struct {
	FullPath string
	Path     string
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
		return nil, fmt.Errorf("Error getting the current working directory: %v", err)
	}

	fullPath := filepath.Join(cwd, targetPath)

	return &ProjectPath{FullPath: fullPath, Path: targetPath}, nil
}
