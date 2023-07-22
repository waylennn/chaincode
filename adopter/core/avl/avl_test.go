package avl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLSampleInsert(t *testing.T) {
	i1 := NewImmutable()
	i1.Insert(mockEntry("a"))
	i1.Insert(mockEntry("e"))
	i1.Insert(mockEntry("c"))
	i1.Insert(mockEntry("f"))
	i1.Insert(mockEntry("b"))
	i1.Insert(mockEntry("d"))
	i1.Insert(mockEntry("g"))
	i1.Insert(mockEntry("A"))

	assert.Equal(t, uint64(8), i1.number)
	entry := i1.root.entry.(mockEntry)
	assert.Equal(t, "c", string(entry))

	min := i1.root
	for min.children[0] != nil {
		min = min.children[0]
	}
	entry = min.entry.(mockEntry)
	assert.Equal(t, "A", string(entry))

	max := i1.root
	for max.children[1] != nil {
		max = max.children[1]
	}
	entry = max.entry.(mockEntry)
	assert.Equal(t, "g", string(entry))
}

func TestImmutableDelete(t *testing.T) {
	i1 := NewImmutable()
	i1.Insert(mockEntry("a"))
	i1.Insert(mockEntry("e"))
	i1.Insert(mockEntry("c"))
	i1.Insert(mockEntry("f"))
	i1.Insert(mockEntry("b"))
	i1.Insert(mockEntry("d"))
	i1.Insert(mockEntry("g"))
	i1.Insert(mockEntry("A"))

	i1.Delete(mockEntry("a"))

	array, err := i1.InOrder()
	assert.Nil(t, err)
	assert.Equal(t, 7, len(array))
}

func TestImmutableInOrder(t *testing.T) {
	i1 := NewImmutable()
	i1.Insert(mockEntry("a"))
	i1.Insert(mockEntry("e"))
	i1.Insert(mockEntry("c"))
	i1.Insert(mockEntry("f"))
	i1.Insert(mockEntry("b"))
	i1.Insert(mockEntry("d"))
	i1.Insert(mockEntry("g"))
	i1.Insert(mockEntry("A"))

	result, err := i1.InOrder()
	assert.Nil(t, err)

	array := make([]string, 0)
	for _, value := range result {
		array = append(array, string(value))
	}

	actual := strings.Join(array, ",")
	assert.Equal(t, "A,a,b,c,d,e,f,g", actual)
}

func BenchmarkImmutableInsert(b *testing.B) {
	s1 := NewImmutable()
	for i := 0; i < b.N; i++ {
		entry := fmt.Sprintf("Item%d", i)
		s1.Insert(mockEntry(entry))
	}
}

func BenchmarkImmutableGet(b *testing.B) {
	s1 := NewImmutable()
	for i := 0; i < b.N; i++ {
		entry := fmt.Sprintf("Item%d", i)
		s1.Insert(mockEntry(entry))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		entry := fmt.Sprintf("Item%d", i)
		s1.Get(mockEntry(entry))
	}
}
