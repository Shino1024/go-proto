package treer

import (
//	"fmt"
	"reflect"
)

type Tree struct {
	root   *node
	length int
	depth  int
}

type node struct {
	left  *node
	right *node
	data  interface{}
}

func InitializeTree(d interface{}) *Tree {
	ret := new(Tree)
	ret.root.left = nil
	ret.root.right = nil
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		ret.root = nil
		ret.length = 0
		ret.depth = 0
		return ret
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.IndexOf(i).Interface()
	}
	depth := 0
	for i := 0; i < temp.Len(); i++ {
		temp3 := new(node)
		temp3 = ret.root
		for temp3.left != nil || temp3.right != nil {
			depth++
			if temp3.data > d {
				if temp3.right == nil {
					temp3 = temp3.right
					temp3.left, temp3.right = nil, nil
					temp3.data = d
				} else {
					temp3 = temp3.right
				}
			} else {
				if temp3.left == nil {
					temp3 = temp3.left
					temp3.left, temp3.right = nil, nil
					temp3.data = d
				} else {
					temp3 = temp3.left
				}
			}
		}
	}
	t.length = temp.Len()
	if depth > ret.depth {
		ret.depth = depth
	}
}

func InitializeTreeA(d []interface{}) *Tree {
	ret := new(Tree)
	ret.root.left = nil
	ret.root.right = nil
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		ret.root = nil
		ret.length = 0
		ret.depth = 0
		return ret
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.IndexOf(i).Interface()
	}
	depth := 0
	for i := 0; i < temp.Len(); i++ {
		temp3 := new(node)
		temp3 = ret.root
		for temp3.left != nil || temp3.right != nil {
			depth++
			if temp3.data > d {
				if temp3.right == nil {
					temp3 = temp3.right
					temp3.left, temp3.right = nil, nil
					temp3.data = d
				} else {
					temp3 = temp3.right
				}
			} else {
				if temp3.left == nil {
					temp3 = temp3.left
					temp3.left, temp3.right = nil, nil
					temp3.data = d
				} else {
					temp3 = temp3.left
				}
			}
		}
	}
	t.length = temp.Len()
	if depth > ret.depth {
		ret.depth = depth
	}
}

func (t *Tree) Append(d ...interface{}) {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		ret.root = nil
		ret.length = 0
		ret.depth = 0
		return ret
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.IndexOf(i).Interface()
	}
	depth := 0
	for i := 0; i < temp.Len(); i++ {
		temp3 := new(node)
		temp3 = ret.root
		for temp3.left != nil || temp3.right != nil {
			depth++
			if temp3.data > d {
				if temp3.right == nil {
					temp3 = temp3.right
					temp3.left, temp3.right = nil, nil
					temp3.data = d
				} else {
					temp3 = temp3.right
				}
			} else {
				if temp3.left == nil {
					temp3 = temp3.left
					temp3.left, temp3.right = nil, nil
					temp3.data = d
				} else {
					temp3 = temp3.left
				}
			}
		}
	}
	t.length += temp.Len()
	if depth > ret.depth {
		ret.depth = depth
	}
}

func (t *Tree) AppendA(d []interface{}) {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		ret.root = nil
		ret.length = 0
		ret.depth = 0
		return ret
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.IndexOf(i).Interface()
	}
	depth := 0
	for i := 0; i < temp.Len(); i++ {
		temp3 := new(node)
		temp3 = ret.root
		for temp3.left != nil || temp3.right != nil {
			depth++
			if temp3.data > d {
				if temp3.right == nil {
					temp3 = temp3.right
					temp3.left, temp3.right = nil, nil
					temp3.data = d
				} else {
					temp3 = temp3.right
				}
			} else {
				if temp3.left == nil {
					temp3 = temp3.left
					temp3.left, temp3.right = nil, nil
					temp3.data = d
				} else {
					temp3 = temp3.left
				}
			}
		}
	}
	t.length = temp.Len()
	if depth > ret.depth {
		ret.depth = depth
	}
}

func (t *Tree) Search(d interface{}) bool {
	temp := t.root
	if temp == nil {
		return false
	}
	if temp.data == d {
		return true
	}
	for temp.left != nil || temp.right != nil {
		if temp.data == d {
			return true
		} else if temp.data > d {
			if temp.left == nil {
				return false
			}
			temp = temp.left
		} else {
			if temp.right == nil {
				return false
			}
			temp = temp.right
		}
	}
	return false
}

func (t *Tree) IsEmpty() bool {
	if t.depth == 0 {
		return true
	}
	return false
}

func emptyHelper(n *node) []interface{} {
	
}

func (t *Tree) Empty() []interface{} {

}

func (t *Tree) Max(d interface{}) (interface{}, error) {

}

func (t *Tree) Min(d interface{}) (interface{}, error) {
	
}

func (t *Tree) Delete(d interface{}) error {
	
}
