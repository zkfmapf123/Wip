package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fs(t *testing.T) {

	p1 := MustGetPwdPath()
	p2 := MustGetHomePath()

	assert.NotNil(t, p1, true)
	assert.NotNil(t, p2, true)
}

func Test_IsCheckGitCredentials(t *testing.T) {
	path, _ := GetGitPath()

	assert.NotNil(t, path, true)
}
