package avl

import (
	"container/list"
	"math"
)

// Immutable represents an immutable AVL tree. This is acheived
// by branch copying
type Immutable struct {
	root   *node
	number uint64
	dummy  node // helper for inserts.
}

// copy returns a copy of this immutable tree with a copy
// of the root and a new dummy helper for the insertion operation
func (immutable *Immutable) copy() *Immutable {
	var root *node
	if immutable.root != nil {
		root = immutable.root.copy()
	}
	cp := &Immutable{
		root:   root,
		number: immutable.number,
		dummy:  *newNode(nil),
	}
	return cp
}

func (immutable *Immutable) resetDummy() {
	immutable.dummy.children[0], immutable.dummy.children[1] = nil, nil
	immutable.dummy.balance = 0
}

func (immutable *Immutable) init() {
	immutable.dummy = node{
		children: [2]*node{},
	}
}

// InOrder implementing the in-order traversal of the binary tree
func (immutable *Immutable) InOrder() ([][]byte, error) {
	tree := immutable
	result := make([][]byte, 0)

	if tree == nil {
		return result, nil
	}

	stack := list.New()
	n := immutable.root

	for n != nil || stack.Len() != 0 {
		for n != nil {
			stack.PushBack(n)
			n = n.children[0]
		}

		if stack.Len() != 0 {
			v := stack.Back()
			n = v.Value.(*node)
			bs, err := n.entry.ToBytes()
			if err != nil {
				return nil, err
			}

			result = append(result, bs)

			n = n.children[1]
			stack.Remove(v)
		}
	}

	return result, nil
}

// Get will get the provided entry from the tree. If no matching
// entry is found, a empty string is returned in its place.
func (immutable *Immutable) Get(entry Entry) Entry {
	n := immutable.root
	var result int
	for n != nil {
		switch result = n.entry.Compare(entry); {
		case result == 0:
			return n.entry
		case result > 0:
			n = n.children[0]
		case result < 0:
			n = n.children[1]
		}
	}

	return nil
}

// Len returns the number of items in this immutable
func (immutable *Immutable) Len() uint64 {
	return immutable.number
}

// Insert will add the provided entry into the tree.
func (immutable *Immutable) Insert(entry Entry) {
	if immutable.root == nil {
		immutable.root = newNode(entry)
		immutable.number++
		return
	}

	immutable.resetDummy()
	var (
		dummy           = immutable.dummy
		p, s, q         *node
		dir, normalized int
		helper          = &dummy
	)

	helper.children[1] = immutable.root
	for s, p = helper.children[1], helper.children[1]; ; {
		dir = p.entry.Compare(entry)

		normalized = normalizeComparison(dir)
		if dir > 0 {
			if p.children[0] != nil {
				q = p.children[0].copy()
				p.children[0] = q
			} else {
				q = nil
			}
		} else if dir < 0 {
			if p.children[1] != nil {
				q = p.children[1].copy()
				p.children[1] = q
			} else {
				q = nil
			}
		} else {
			p.entry = entry
			return
		}

		if q == nil {
			break
		}

		if q.balance != 0 {
			helper = p
			s = q
		}
		p = q
	}

	immutable.number++
	q = newNode(entry)
	p.children[normalized] = q

	immutable.root = dummy.children[1]
	for p = s; p != q; p = p.children[normalized] {
		normalized = normalizeComparison(p.entry.Compare(entry))
		if normalized == 0 {
			p.balance--
		} else {
			p.balance++
		}
	}

	q = s

	if math.Abs(float64(s.balance)) > 1 {
		normalized = normalizeComparison(s.entry.Compare(entry))
		s = insertBalance(s, normalized)
	}

	if q == dummy.children[1] {
		immutable.root = s
	} else {
		helper.children[intFromBool(helper.children[1] == q)] = s
	}
}

// Delete will remove the provided entry from this AVL tree
func (immutable *Immutable) Delete(entry Entry) bool {
	if immutable.root == nil {
		return false
	}

	var (
		cache                      = make(nodes, 64)
		it, p, q                   *node
		top, done, dir, normalized int
		dirs                       = make([]int, 64)
		found                      bool
	)

	it = immutable.root

	for {
		if it == nil {
			return false
		}

		dir = it.entry.Compare(entry)
		if dir == 0 {
			found = true
			break
		}

		normalized = normalizeComparison(dir)
		dirs[top] = normalized
		cache[top] = it
		top++
		it = it.children[normalized]
	}
	immutable.number++

	for i := 0; i < top; i++ {
		p = cache[i]
		if p.children[dirs[i]] != nil {
			q = p.children[dirs[i]].copy()
			p.children[dirs[i]] = q
			if i != top-1 {
				cache[i+1] = q
			}
		}
	}
	it = it.copy()

	oldTop := top
	if it.children[0] == nil || it.children[1] == nil {
		dir = intFromBool(it.children[0] == nil)
		if top != 0 {
			cache[top-1].children[dirs[top-1]] = it.children[dir]
		} else {
			immutable.root = it.children[dir]
		}
	} else {
		heir := it.children[1]
		dirs[top] = 1
		cache[top] = it
		top++

		for heir.children[0] != nil {
			dirs[top] = 0
			cache[top] = heir
			top++
			heir = heir.children[0]
		}

		it.entry = heir.entry
		if oldTop != 0 {
			cache[oldTop-1].children[dirs[oldTop-1]] = it
		} else {
			immutable.root = it
		}
		cache[top-1].children[intFromBool(cache[top-1] == it)] = heir.children[1]
	}

	for top-1 >= 0 && done == 0 {
		top--
		if dirs[top] != 0 {
			cache[top].balance--
		} else {
			cache[top].balance++
		}

		if math.Abs(float64(cache[top].balance)) == 1 {
			break
		} else if math.Abs(float64(cache[top].balance)) > 1 {
			cache[top] = removeBalance(cache[top], dirs[top], &done)

			if top != 0 {
				cache[top-1].children[dirs[top-1]] = cache[top]
			} else {
				immutable.root = cache[0]
			}
		}
	}

	return found
}

func removeBalance(root *node, dir int, done *int) *node {
	n := root.children[takeOpposite(dir)].copy()
	root.children[takeOpposite(dir)] = n
	var bal int8
	if dir == 0 {
		bal = -1
	} else {
		bal = 1
	}

	if n.balance == -bal {
		root.balance, n.balance = 0, 0
		root = rotate(root, dir)
	} else if n.balance == bal {
		adjustBalance(root, takeOpposite(dir), int(-bal))
		root = doubleRotate(root, dir)
	} else {
		root.balance = -bal
		n.balance = bal
		root = rotate(root, dir)
		*done = 1
	}

	return root
}

func intFromBool(value bool) int {
	if value {
		return 1
	}

	return 0
}

func insertBalance(root *node, dir int) *node {
	n := root.children[dir]
	var bal int8
	if dir == 0 {
		bal = -1
	} else {
		bal = 1
	}

	if n.balance == bal {
		root.balance, n.balance = 0, 0
		root = rotate(root, takeOpposite(dir))
	} else {
		adjustBalance(root, dir, int(bal))
		root = doubleRotate(root, takeOpposite(dir))
	}

	return root
}

func adjustBalance(root *node, dir, bal int) {
	n := root.children[dir]
	nn := n.children[takeOpposite(dir)]

	if nn.balance == 0 {
		root.balance, n.balance = 0, 0
	} else if int(nn.balance) == bal {
		root.balance = int8(-bal)
		n.balance = 0
	} else {
		root.balance = 0
		n.balance = int8(bal)
	}
	nn.balance = 0
}

func rotate(parent *node, dir int) *node {
	otherDir := takeOpposite(dir)

	child := parent.children[otherDir]
	parent.children[otherDir] = child.children[dir]
	child.children[dir] = parent

	return child
}

func doubleRotate(parent *node, dir int) *node {
	otherDir := takeOpposite(dir)

	parent.children[otherDir] = rotate(parent.children[otherDir], otherDir)
	return rotate(parent, dir)
}

func takeOpposite(value int) int {
	return 1 - value
}

func normalizeComparison(i int) int {
	if i < 0 {
		return 1
	}

	if i > 0 {
		return 0
	}

	return -1
}

// NewImmutable allocates, initializesm, and returns a new immutable
// AVL tree.
func NewImmutable() *Immutable {
	immutable := &Immutable{}
	immutable.init()
	return immutable
}
