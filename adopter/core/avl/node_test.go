package avl

type mockEntry string

func (me mockEntry) Compare(other Entry) int {
	otherMe := other.(mockEntry)

	if me > otherMe {
		return 1
	}

	if me < otherMe {
		return -1
	}

	return 0
}

func (me mockEntry) ToBytes() ([]byte, error) {
	return []byte(me), nil
}
