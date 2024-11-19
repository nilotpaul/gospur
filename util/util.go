package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const maxNestingDepth = 3

func SanitizeDirPath(path string) (string, error) {
	dir := filepath.Clean(path)

	// Check for invalid paths like `/../`.
	if strings.Contains(dir, "..") {
		return "", fmt.Errorf("Invalid directory path: '%s' contains '..'", dir)
	}

	// Check the nesting depth.
	depth := strings.Count(dir, string(filepath.Separator))
	// Avoid deep nesting for paths more than 3 depth.
	if depth > maxNestingDepth {
		return "", fmt.Errorf("Invalid directory path: exceeds maximum allowed depth of %d", maxNestingDepth)

	}

	return dir, nil
}

func RunGoModInit(fullProjectPath, name string) error {
	// Change the current working directory to the project directory
	if err := os.Chdir(fullProjectPath); err != nil {
		return fmt.Errorf("Failed to change to project directory: %v", err)
	}

	cmd := exec.Command("go", "mod", "init", name)
	return cmd.Run()
}

func CreateTargetDir(path string, strict bool) error {
	if strict {
		_, err := doesTargetDirExistAndIsEmpty(path)
		if err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	// Target dir doesn't exist, we need to create it.
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func doesTargetDirExistAndIsEmpty(target string) (bool, error) {
	file, err := os.Stat(target)
	if err != nil {
		return false, err
	}
	if !file.IsDir() {
		return false, fmt.Errorf("'%s' is not a directory", target)
	}

	entires, err := os.ReadDir(target)
	if err != nil {
		return false, err
	}

	return len(entires) == 0, fmt.Errorf("'%s' is not empty", target)
}
