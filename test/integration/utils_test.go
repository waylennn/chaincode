package integration

import (
	"io/ioutil"
	"path"
	"testing"

	"git.querycap.com/cloudchain/chaincode/utils"
	"github.com/stretchr/testify/assert"
)

func TestSHA1Compare(t *testing.T) {
	file1, err := ioutil.ReadFile(path.Join(GetTestFixturePath(), "sample.docx"))
	assert.Nil(t, err)

	sha1, err := utils.ComputeSHA1(file1)
	assert.Nil(t, err)

	file2, err := ioutil.ReadFile(path.Join(GetTestFixturePath(), "sample_new.docx"))
	assert.Nil(t, err)

	sha2, err := utils.ComputeSHA1(file2)
	assert.Nil(t, err)

	assert.Equal(t, sha1, sha2)
}
