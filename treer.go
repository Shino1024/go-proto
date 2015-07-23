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
	data  {}interface
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
	for i := 0; i < temp.Len(); i++ {
		temp3 := new(node)
		temp3 = ret.root
		temp3.data = temp2[i]
		for temp3.left != nil || temp3.right != nil {
			if 
		}
	}
}

func InitializeTreeA(d []interface{}) *Tree {
	
}

