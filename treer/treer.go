//Simple red-black tree data structure. It's a modified version of sakeven's red-black tree implementation.
package goproto

import "fmt"

//Constants.
const (
	RED   = 0
	BLACK = 1
)

//Interface TLesser to be implemented for objects that will be stored in the tree.
type TLesser interface {
	//TLess takes an interface{} and returns a bool. Compare between the receiver and some interface.
	TLess(other interface{}) bool
}

//Helper node data structure of which trees consist from.
type node struct {
	left *node
	right *node
	parent *node
	color int
	Key TLesser
	Value interface{}
}

//Tree data structure.
type Tree struct {
	root *node
	size int
}

//Create a new tree with this function.
func NewTree() *Tree {
	return &Tree{}
}

//Find the value corresponding to the given key.
func (t *Tree) Find(key TLesser) interface{} {
	n := t.findNode(key)
	if n != nil {
		return n.Value
	}
	return nil
}

//Return an iterator to the key.
func (t *Tree) FindIt(key TLesser) *node {
	return t.findNode(key)
}

//Check whether the tree is empty. Returns bool.
func (t *Tree) IsEmpty() bool {
	if t.root == nil {
		return true
	}
	return false
}

//Get the iterator that points to the minimum value in the tree.
func (t *Tree) Iterator() *node {
	return minimum(t.root)
}

//Get the length of the tree.
func (t *Tree) Len() int {
	return t.size
}

//Erase everything from the tree.
func (t *Tree) Empty() {
	t.root = nil
	t.size = 0
}

//Insert any object with this function with its key.
func (t *Tree) Insert(key TLesser, value interface{}) {
	x := t.root
	var y *node

	for x != nil {
		y = x
		if key.TLess(x.Key) {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &node{parent: y, color: RED, Key: key, Value: value}
	t.size += 1

	if y == nil {
		z.color = BLACK
		t.root = z
		return
	} else if z.Key.TLess(y.Key) {
		y.left = z
	} else {
		y.right = z
	}
	t.rbInsertFixUp(z)
}

//Delete any object by its key.
func (t *Tree) Delete(key TLesser) {
	z := t.findNode(key)
	if z == nil {
		return
	}

	var x, y, parent *node
	y = z
	y_original_color := y.color
	parent = z.parent
	if z.left == nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = minimum(z.right)
		y_original_color = y.color
		x = y.right

		if y.parent == z {
			if x == nil {
				parent = y
			} else {
				x.parent = y
			}
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if y_original_color == BLACK {
		t.rbDeleteFixUp(x, parent)
	}
	t.size -= 1
}

//Helper function for fixing up the tree after insertion.
func (t *Tree) rbInsertFixUp(z *node) {
	var y *node
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = BLACK
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

//Helper function for fixing up the tree after deletion.
func (t *Tree) rbDeleteFixUp(x, parent *node) {
	var w *node

	for x != t.root && getColor(x) == BLACK {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			w = parent.right
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.leftRotate(x.parent)
				w = parent.right
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.right) == BLACK {
					if w.left != nil {
						w.left.color = BLACK
					}
					w.color = RED
					t.rightRotate(w)
					w = parent.right
				}
				w.color = parent.color
				parent.color = BLACK
				if w.right != nil {
					w.right.color = BLACK
				}
				t.leftRotate(parent)
				x = t.root
			}
		} else {
			w = parent.left
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.rightRotate(parent)
				w = parent.left
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.left) == BLACK {
					if w.right != nil {
						w.right.color = BLACK
					}
					w.color = RED
					t.leftRotate(w)
					w = parent.left
				}
				w.color = parent.color
				parent.color = BLACK
				if w.left != nil {
					w.left.color = BLACK
				}
				t.rightRotate(parent)
				x = t.root
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

//Left-rotate the tree.
func (t *Tree) leftRotate(x *node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

//Right-rotate the tree.
func (t *Tree) rightRotate(x *node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = x
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.right = x
	x.parent = y
}

//Display the preorder of the tree.
func (t *Tree) Preorder() {
	if t.root != nil {
		t.root.preorder()
	}
	fmt.Println()
}

//Find the node by its key.
func (t *Tree) findNode(key TLesser) *node {
	x := t.root
	for x != nil {
		if key.TLess(x.Key) {
			x = x.left
		} else {
			if !key.TLess(x.Key) && !x.Key.TLess(key) {
				return x
			} else {
				x = x.right
			}
		}
	}
	return nil
}

//Transplant the subtrees by their root nodes.
func (t *Tree) transplant(u, v *node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v == nil {
		return
	}
	v.parent = u.parent
}

//Return the successor of. the node.
func (n *node) Next() *node {
	return successor(n)
}

//Helper preorder function that does the work recursively.
func (n *node) preorder() {
	fmt.Printf("(%v %v)", n.Key, n.Value)
	if n.parent == nil {
		fmt.Printf("nil")
	} else {
		fmt.Printf(", whose parent is %v", n.parent.Key)
	}
	if n.color == RED {
		fmt.Println(" and color RED")
	} else {
		fmt.Println(" and color BLACK")
	}
	if n.left != nil {
		fmt.Printf("%v's left child is ", n.Key)
		n.left.preorder()
	}
	if n.right != nil {
		fmt.Printf("%v's right child is ", n.Key)
		n.right.preorder()
	}
}

//Helper function for returning the successor of the node.
func successor(x *node) *node {
	if x.right != nil {
		return minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}
	return y
}

//Get the node's color.
func getColor(n *node) int {
	if n == nil {
		return BLACK
	}
	return n.color
}

//Find the minimum node of the subtree.
func minimum(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

//Find the maximum node of the subtree.
func maximum(n *node) *node {
	for n.right != nil {
		n = n.right
	}
	return n
}