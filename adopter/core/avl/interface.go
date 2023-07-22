package avl

// Entry represents all items that can be placed into the AVL tree
type Entry interface {
	// Compare should return a value indicating the relationship
	// of thie Entry to the provided Entry. A -1 means this entry
	// is less than, 0 means equality, and 1 means greater than.
	Compare(Entry) int

	// ToBytes serializes the instance to bytes
	ToBytes() ([]byte, error)
}
