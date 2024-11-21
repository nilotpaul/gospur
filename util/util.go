package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const maxNestingDepth = 3

// SanitizeDirPath takes a `path` and checks if the given
// project path is valid or not.
func ValidateDirPath(path string) (string, error) {
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

// CreateTargetDir takes a `path` and `strict`,
//
// In strict mode, it'll check if the directory is empty or not
// if the dir already exists. If the dir doesn't exist it'll create one.
//
// If not in strict mode, it'll ignore the directory status and
// create the necessary dir(s).
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

// doesTargetDirExistAndIsEmpty takes a `target` path, if it's
// not a directory, not empty or doesn't exist then it'll return
// false and an error, otherwise true and nil error.
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

	if len(entires) != 0 {
		return false, fmt.Errorf("'%s' is not empty", target)
	}

	return true, nil
}
