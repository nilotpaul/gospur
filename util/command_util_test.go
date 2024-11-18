package util

import (
	"testing"

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
