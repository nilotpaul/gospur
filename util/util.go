package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const maxNestingDepth = 3

func createTargerDir(path string) error {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Target dir doesn't exist, we need create it.
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func SanitizeDirPath(path string) (string, error) {
	dir := filepath.Clean(path)

	// Check for invalid paths like `/../`.
	if strings.Contains(dir, "..") {
		return "", fmt.Errorf("invalid directory path: '%s' contains '..'", dir)
	}

	// Check the nesting depth.
	depth := strings.Count(dir, string(filepath.Separator))
	// Avoid deep nesting for paths more than 3 depth.
	if depth > maxNestingDepth {
		return "", fmt.Errorf("invalid directory path: exceeds maximum allowed depth of %d", maxNestingDepth)

	}

	return dir, nil
}
