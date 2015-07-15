package lister

import (
	"fmt"
	"errors"
	"reflect"
)

type List struct {
	startingNode *node
	length       int
}

type node struct {
	nextNode *node
	data     interface{}
}

func InitializeList() *List {
	l := new(List)
	l.startingNode = nil
	l.length = 0
	return l
}

func InitializeListVar(d ...interface{}) *List {
	l := new(List)
	if len(d) == 0 {
		return InitializeList()
	} else {
		l.startingNode = nil
		l.length = 0
		for _, v := range d {
			l.Append(v)
		}
	}
	return l
}

func InitializeListA(d interface{}) *List {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return InitializeList()
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	l := new(List)
	l.startingNode = nil
	l.length = 0
	for _, v := range temp2 {
		l.Append(v)
	}
	return l
}

func (l *List) IsEmpty() bool {
	if l.length != 0 {
		return false
	}
	return true
}

func (l *List) Append(d interface{}) {
	n := new(node)
	n.nextNode = nil
	n.data = d
	if l.startingNode == nil {
		l.startingNode = n
		l.length = 1
	} else {
		t := new(node)
		t = l.startingNode
		for ; t.nextNode != nil; t = t.nextNode {
		}
		t.nextNode = n
		l.length++
	}
}

func (l *List) AppendA(d interface{}) {
	temp := reflect.ValueOf(d)
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	for _, v := range temp2 {
		l.Append(v)
	}
}

func (l *List) AppendVar(d ...interface{}) {
	
}

func (l *List) PrintAll(sepstr ...string) {
	if l.length == 0 {
		return
	}
	sep := ", "
	if len(sepstr) == 1 {
		sep = sepstr[0]
	}
	fmt.Print("[")
	for n := l.startingNode; n != nil; n = n.nextNode {
		fmt.Print(n.data)
		if n.nextNode == nil {
			break
		}
		fmt.Print(sep)
	}
	fmt.Print("]")
}

func (l *List) PrintAllln(sepstr ...string) {
	if l.length == 0 {
		return
	}
	sep := ", "
	if len(sepstr) == 1 {
		sep = sepstr[0]
	}
	fmt.Print("[")
	for n := l.startingNode; n != nil; n = n.nextNode {
		fmt.Print(n.data)
		if n.nextNode == nil {
			break
		}
		fmt.Print(sep)
	}
	fmt.Println("]")
}

func (l *List) Search(d interface{}) bool {
	for temp := l.startingNode; temp != nil; temp = temp.nextNode {
		if temp.data == d {
			return true
		}
	}
	return false
}


func (l *List) Delete(d interface{}) (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("The list is empty.")
	} else if l.length == 1 && l.startingNode.data == d {
		ret := l.startingNode.data
		l = InitializeList()
		return ret, nil
	}
	temp := new(node)
	toDelete := new(node)
	for toDelete = l.startingNode; toDelete.data != d; toDelete = toDelete.nextNode {
		if toDelete == nil {
			return nil, errors.New("Haven't found the element to delete.")
		}
	}
	for temp = l.startingNode; temp.nextNode != toDelete; temp = temp.nextNode {
	}
	
}

