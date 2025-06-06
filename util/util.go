package util

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nilotpaul/gospur/config"
)

const maxNestingDepth = 3

// SanitizeDirPath takes a `path` and checks if the given
// project path is valid or not.
func ValidateDirPath(path string) (string, error) {
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

func MakeProjectCtx(cfg StackConfig, modPath string) map[string]any {
	return map[string]any{
		"ModPath": modPath,
		"IsLinux": strings.Split(runtime.GOOS, "/")[0] == "linux",
		"Web": map[string]bool{
			"IsEcho":  cfg.WebFramework == "Echo",
			"IsFiber": cfg.WebFramework == "Fiber",
			"IsChi":   cfg.WebFramework == "Chi",
		},
		"UI": map[string]bool{
			// CSS Strategy
			"HasTailwind":  strings.HasPrefix(cfg.CssStrategy, "Tailwind"),
			"HasTailwind4": cfg.CssStrategy == "Tailwind (v4)",
			"HasTailwind3": cfg.CssStrategy == "Tailwind (v3)",

			// CSS Library
			"HasPreline": cfg.UILibrary == "Preline",
			"HasDaisy":   cfg.UILibrary == "DaisyUI",
		},
		"Render": map[string]bool{
			"IsTemplates": cfg.RenderingStrategy == "Templates",
			"IsSeperate":  cfg.RenderingStrategy == "Seperate",
		},
		"Extras": map[string]bool{
			"HasHTMX": contains(cfg.ExtraOpts, "HTMX"),
		},
	}
}

// AutoDetectBinaryURL loops over assets (binary links) and returns one
// compatible with the current system.
func FindMatchingBinary(names []string, os string, arch string) string {
	os, arch = mapRuntimeOSAndArch(os, arch)
	expected := fmt.Sprintf("gospur_%s_%s", os, arch)

	for _, name := range names {
		if strings.Contains(name, expected) {
			return name
		}
	}

	return ""
}

func GetRenderingOpts(actual bool) []string {
	opts := make([]string, len(config.RenderingStrategy))
	idx := 0

	for name, actualName := range config.RenderingStrategy {
		if actual {
			opts[idx] = actualName
		} else {
			opts[idx] = name
		}
		idx++
	}

	return opts
}

func GetMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0)
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func uncompress(src io.Reader, url string) (io.Reader, error) {
	if strings.HasSuffix(url, ".zip") {
		buf, err := io.ReadAll(src)
		if err != nil {
			return nil, fmt.Errorf("failed to read the release .zip file: %v", err)
		}

		r := bytes.NewReader(buf)
		z, err := zip.NewReader(r, r.Size())
		if err != nil {
			return nil, fmt.Errorf("failed to uncompress the .zip file: %v", err)
		}

		for _, file := range z.File {
			_, name := filepath.Split(file.Name)
			if !file.FileInfo().IsDir() && matchBinaryFile(name) {
				return file.Open()
			}
		}
	} else if strings.HasSuffix(url, ".tar.gz") {
		gz, err := gzip.NewReader(src)
		if err != nil {
			return nil, fmt.Errorf("failed to uncompress the .tar.gz file: %v", err)
		}

		return unarchiveTarGZ(gz)
	}

	return nil, fmt.Errorf("given file is not .tar.gz or .zip format")
}

func unarchiveTarGZ(src io.Reader) (io.Reader, error) {
	t := tar.NewReader(src)
	for {
		h, err := t.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to unarchive .tar.gz file: %v", err)
		}
		_, name := filepath.Split(h.Name)
		if matchBinaryFile(name) {
			return t, nil
		}
	}

	return nil, fmt.Errorf("binary not found after uncompressing")
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

// Map Go's OS to GoReleaser's naming convention.
func mapRuntimeOSAndArch(os string, arch string) (mappedOS string, mappedArch string) {
	switch os {
	case "darwin":
		mappedOS = "Darwin"
	case "linux":
		mappedOS = "Linux"
	case "windows":
		mappedOS = "Windows"
	default:
		mappedOS = os
	}

	switch arch {
	case "amd64":
		mappedArch = "x86_64"
	case "386":
		mappedArch = "i386"
	case "arm64":
		mappedArch = "arm64"
	default:
		mappedArch = arch
	}

	return mappedOS, mappedArch
}

// matchBinaryFile checks returns true if the binary name is correct.
func matchBinaryFile(name string) bool {
	switch name {
	case config.WinBinaryName:
		return true
	case config.OtherBinaryName:
		return true
	default:
		return false
	}
}

// contains checks if a slice of string contains the given item.
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func removeLinesStartEnd(s string, start, end int) string {
	lines := strings.Split(s, "\n")
	if len(lines) > 2 {
		lines = lines[start : len(lines)-end]
	}

	return strings.Join(lines, "\n")
}
