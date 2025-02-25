package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchBinary(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	// Invalid binary name
	ok := matchBinaryFile("bad")
	a.False(ok)

	// Invalid binary name for windows
	ok = matchBinaryFile("bad.exe")
	a.False(ok)

	// Correct binary name
	ok = matchBinaryFile("gospur")
	a.True(ok)

	// Correct binary name for windows
	ok = matchBinaryFile("gospur.exe")
	a.True(ok)
}

func TestFindMatchingBinary(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	binaries := []string{
		"gospur_0.7.1_checksums.txt",
		"gospur_Darwin_arm64.tar.gz",
		"gospur_Darwin_x86_64.tar.gz",
		"gospur_Linux_arm64.tar.gz",
		"gospur_Linux_i386.tar.gz",
		"gospur_Linux_x86_64.tar.gz",
		"gospur_Windows_arm64.zip",
		"gospur_Windows_i386.zip",
		"gospur_Windows_x86_64.zip",
	}

	// Test for different OS and architectures
	tests := []struct {
		os       string
		arch     string
		expected string
	}{
		{"darwin", "arm64", "gospur_Darwin_arm64.tar.gz"},
		{"darwin", "amd64", "gospur_Darwin_x86_64.tar.gz"},
		{"linux", "arm64", "gospur_Linux_arm64.tar.gz"},
		{"linux", "386", "gospur_Linux_i386.tar.gz"},
		{"linux", "amd64", "gospur_Linux_x86_64.tar.gz"},
		{"windows", "arm64", "gospur_Windows_arm64.zip"},
		{"windows", "386", "gospur_Windows_i386.zip"},
		{"windows", "amd64", "gospur_Windows_x86_64.zip"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s-%s", test.os, test.arch), func(t *testing.T) {
			binary := FindMatchingBinary(binaries, test.os, test.arch)
			a.Equal(test.expected, binary)
		})
	}
}
