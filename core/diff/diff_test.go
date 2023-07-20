package diff

import (
	"io/ioutil"
	"path"
	"testing"

	"git.querycap.com/cloudchain/chaincode/utils"
	"github.com/stretchr/testify/assert"
)

func TestDiff1(t *testing.T) {
	src, err := ioutil.ReadFile(path.Join(testFixturePath(), "src.xlsx"))
	assert.Nil(t, err)

	dst, err := ioutil.ReadFile(path.Join(testFixturePath(), "dst1.xlsx"))
	assert.Nil(t, err)

	a, u, d, err := Do(dst, src, []string{"姓名"})
	assert.Nil(t, err)
	assert.Equal(t, uint32(0), a)
	assert.Equal(t, uint32(0), u)
	assert.Equal(t, uint32(1), d)
}

func TestDiff2(t *testing.T) {
	src, err := ioutil.ReadFile(path.Join(testFixturePath(), "src.xlsx"))
	assert.Nil(t, err)

	dst, err := ioutil.ReadFile(path.Join(testFixturePath(), "dst2.xlsx"))
	assert.Nil(t, err)

	a, u, d, err := Do(dst, src, []string{"姓名"})
	assert.Nil(t, err)
	assert.Equal(t, uint32(0), a)
	assert.Equal(t, uint32(1), u)
	assert.Equal(t, uint32(0), d)
}

func TestDiff3(t *testing.T) {
	src, err := ioutil.ReadFile(path.Join(testFixturePath(), "src.xlsx"))
	assert.Nil(t, err)

	dst, err := ioutil.ReadFile(path.Join(testFixturePath(), "dst3.xlsx"))
	assert.Nil(t, err)

	a, u, d, err := Do(dst, src, []string{"姓名"})
	assert.Nil(t, err)
	assert.Equal(t, uint32(0), a)
	assert.Equal(t, uint32(3), u)
	assert.Equal(t, uint32(1), d)
}

func TestDiff4(t *testing.T) {
	src, err := ioutil.ReadFile(path.Join(testFixturePath(), "src.xlsx"))
	assert.Nil(t, err)

	dst, err := ioutil.ReadFile(path.Join(testFixturePath(), "dst4.xlsx"))
	assert.Nil(t, err)

	a, u, d, err := Do(dst, src, []string{"姓名"})
	assert.Nil(t, err)
	assert.Equal(t, uint32(0), a)
	assert.Equal(t, uint32(0), u)
	assert.Equal(t, uint32(4), d)
}

func TestDiff5(t *testing.T) {
	src, err := ioutil.ReadFile(path.Join(testFixturePath(), "src.xlsx"))
	assert.Nil(t, err)

	dst, err := ioutil.ReadFile(path.Join(testFixturePath(), "dst5.xlsx"))
	assert.Nil(t, err)

	a, u, d, err := Do(dst, src, []string{"姓名"})
	assert.Nil(t, err)
	assert.Equal(t, uint32(1), a)
	assert.Equal(t, uint32(1), u)
	assert.Equal(t, uint32(1), d)
}

func testFixturePath() string {
	return path.Join(utils.GoPath(), "src", "git.querycap.com/cloudchain/chaincode", "test/fixtures/diff")
}
