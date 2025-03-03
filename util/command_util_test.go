package util

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetProjectPath(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	defaultProjectPath := "gospur"

	// With no given arg
	pp, err := GetProjectPath([]string{})
	a.NoError(err)
	a.NotNil(pp)
	a.NotEmpty(pp.FullPath)
	a.Equal(defaultProjectPath, pp.Path)

	// With given arg `.` (current dir)
	pp, err = GetProjectPath([]string{"."})
	a.NoError(err)
	a.NotNil(pp)
	a.NotEmpty(pp.FullPath)
	a.Equal(".", pp.Path)

	// With given arg `new-project`
	pp, err = GetProjectPath([]string{"new-project"})
	a.NoError(err)
	a.NotNil(pp)
	a.NotEmpty(pp.FullPath)
	a.Equal("new-project", pp.Path)

	// With given arg `./new-project`
	pp, err = GetProjectPath([]string{"./new-project"})
	a.NoError(err)
	a.NotNil(pp)
	a.NotEmpty(pp.FullPath)
	a.Equal("new-project", pp.Path)

	// With given arg `../new-project`
	pp, err = GetProjectPath([]string{"../new-project"})
	a.Error(err)
	a.ErrorContains(err, "invalid directory path: '../new-project' contains '..'")
	a.Nil(pp)
}

func TestValidateGoModPath(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	// Given path is less than 3 character(s)
	err := validateGoModPath("ww")
	a.EqualError(err, "path cannot be less than 3 character(s)")

	// Given path contains https://
	err = validateGoModPath("https://something")
	a.EqualError(err, "invalid path 'https://something', should not contain https")

	// Given path contains a space
	err = validateGoModPath("some thing")
	a.EqualError(err, "invalid path 'some thing', contains reserved characters")

	// Given path contains a :
	err = validateGoModPath("some:thing")
	a.EqualError(err, "invalid path 'some:thing', contains reserved characters")

	// Given path contains *
	err = validateGoModPath("some*thing")
	a.EqualError(err, "invalid path 'some*thing', contains reserved characters")

	// Given path contains ?
	err = validateGoModPath("github.com/paul?key=value")
	a.EqualError(err, "invalid path 'github.com/paul?key=value', contains reserved characters")

	// Given path contains |
	err = validateGoModPath("github.com/paul|repo")
	a.EqualError(err, "invalid path 'github.com/paul|repo', contains reserved characters")

	// Given path exceedes 255 character(s)
	err = validateGoModPath(generateRandomString(500))
	a.EqualError(err, "exceeded maximum length")

	// Given path is valid
	err = validateGoModPath("github.com/nilotpaul/gospur")
	a.NoError(err)

	// Given another path is valid
	err = validateGoModPath("gospur")
	a.NoError(err)
}

// Helper func only for testing
//
// GenerateRandomString generates a random string of a given length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.NewSource(time.Now().UnixNano())

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
