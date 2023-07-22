package utils

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeSHA1(t *testing.T) {
	file1, err := ioutil.ReadFile("sha1_test.go")
	assert.Nil(t, err)

	hash1, err := ComputeSHA1(file1)
	assert.Nil(t, err)
	assert.NotEmpty(t, hash1)

	file2, err := ioutil.ReadFile("sha1.go")
	assert.Nil(t, err)

	hash2, err := ComputeSHA1(file2)
	assert.Nil(t, err)
	assert.NotEmpty(t, hash2)

	assert.NotEqual(t, hash1, hash2)

	sameFile, err := ioutil.ReadFile("sha1_test.go")
	assert.Nil(t, err)

	sameHash, err := ComputeSHA1(sameFile)
	assert.Nil(t, err)
	assert.NotEmpty(t, sameHash)
	assert.Equal(t, hash1, sameHash)
}
