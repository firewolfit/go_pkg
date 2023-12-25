package file

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewFile(t *testing.T) {
	fileName := "a/b/aa.file"
	x := NewFile(fileName, false)
	assert.Nil(t, x)
	assert.True(t, IsExisted(fileName, false))
	os.RemoveAll("a")
}

func TestNewFolder(t *testing.T) {
	folder := "a/b/c"
	x := NewFolder(folder)
	assert.Nil(t, x)
	assert.True(t, IsExisted(folder, true))
	os.RemoveAll("a")
}

func TestIsExisted(t *testing.T) {
	assert.False(t, IsExisted("func1.go", false))
	assert.True(t, IsExisted("func.go", false))
	assert.False(t, IsExisted("func.go", true))
}
