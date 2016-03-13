//Simple list data structure.
package lister

import (
	"fmt"
	"errors"
	"reflect"
)

//List data structure.
type List struct {
	startingNode *node
	endingNode   *node
	length       int
}

//Helper node data structure of which lists consist from.
type node struct {
	nextNode *node
	data     interface{}
}

//Initialize the list. Use any number of interface{}'s.
func InitializeList(d ...interface{}) *List {
	temp := reflect.ValueOf(d)
	l := new(List)
	if temp.Len() == 0 {
		l.startingNode = nil
		l.endingNode = nil
		l.length = 0
		return l
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	temp3 := new(node)
	temp3.nextNode = nil
	l.startingNode = temp3
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			l.endingNode = temp3
			temp3.nextNode = nil
			break
		}
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	l.length = temp.Len()
	return l
}

//Initialize the list. Use an array of interface{}'s.
func InitializeListA(d interface{}) *List {
	temp := reflect.ValueOf(d)
	l := new(List)
	if temp.Len() == 0 {
		l.startingNode = nil
		l.endingNode = nil
		l.length = 0
		return l
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	temp3 := new(node)
	temp3.nextNode = nil
	l.startingNode = temp3
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			l.endingNode = temp3
			temp3.nextNode = nil
			break
		} 
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	l.length = temp.Len()
	return l
}

//Check if the list is empty.
func (l *List) IsEmpty() bool {
	if l.length != 0 {
		return false
	}
	return true
}

//Return the length of the list.
func (l *List) Len() int {
	return l.length
}

//Append any number of interface{} objects to the list.
func (l *List) Append(d ...interface{}) {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	temp3 := new(node)
	temp3.nextNode = nil
	if l.length == 0 {
		l.startingNode = temp3
	} else {
		l.endingNode.nextNode = temp3
	}
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			l.endingNode = temp3
			temp3.nextNode = nil
			break
		} 
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	l.length += temp.Len()
}

//Append an array of interface{} objects to the list.
func (l *List) AppendA(d interface{}) {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	temp3 := new(node)
	temp3.nextNode = nil
	if l.length == 0 {
		l.startingNode = temp3
	} else {
		l.endingNode.nextNode = temp3
	}
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			l.endingNode = temp3
			temp3.nextNode = nil
			break
		} 
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	l.length += temp.Len()
}

//Print all objects from the list. They, of course, should implement Stringer as well. An optional separator can be provided.
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

//Print all objects from the list and a new line. They, of course, should implement Stringer as well. An optional separator can be provided.
func (l *List) PrintAllln(sepstr ...string) {
	if l.length == 0 {
		fmt.Println("[]")
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

//Check if the object is in the list.
func (l *List) Search(d interface{}) bool {
	for temp := l.startingNode; temp != nil; temp = temp.nextNode {
		if reflect.DeepEqual(temp.data, d) {
			return true
		}
	}
	return false
}

//Remove any number of objects from the list.
func (l *List) Delete(d ...interface{}) error {
	for _, v := range d {
		if l.length == 0 {
			return errors.New("The list is empty.")
		} else if reflect.DeepEqual(l.startingNode.data, v) == true {
			if l.length == 1 {
				l = InitializeList()
			} else {
				l.startingNode = l.startingNode.nextNode
			}
			l.length--
			return nil
		} else if l.length == 1 && reflect.DeepEqual(l.startingNode.data, v) == false {
			return errors.New("Haven't found the element to delete.")
		}
		toDelete := new(node)
		for toDelete = l.startingNode; reflect.DeepEqual(toDelete.data, v) == false; toDelete = toDelete.nextNode {
			if toDelete.nextNode == nil && reflect.DeepEqual(toDelete.data, v) == false {
				return errors.New("Haven't found the element to delete.")
			}
		}
		temp := new(node)
		for temp = l.startingNode; temp.nextNode != toDelete; temp = temp.nextNode {
		}
		if toDelete.nextNode == nil {
			temp.nextNode = nil
			l.length--
			return nil
		}
		temp.nextNode = temp.nextNode.nextNode
		toDelete = nil
		l.length--
	}

	return nil
}

//Remove an array of objects from the list.
func (l *List) DeleteA(d []interface{}) error {
	for _, v := range d {
		if l.length == 0 {
			return errors.New("The list is empty.")
		} else if reflect.DeepEqual(l.startingNode.data, v) == true {
			if l.length == 1 {
				l = InitializeList()
			} else {
				l.startingNode = l.startingNode.nextNode
			}
			l.length--
			return nil
		} else if l.length == 1 && reflect.DeepEqual(l.startingNode.data, v) == false {
			return errors.New("Haven't found the element to delete.")
		}
		toDelete := new(node)
		for toDelete = l.startingNode; reflect.DeepEqual(toDelete.data, v) == false; toDelete = toDelete.nextNode {
			if toDelete.nextNode == nil && reflect.DeepEqual(toDelete.data, v) == false {
				return errors.New("Haven't found the element to delete.")
			}
		}
		temp := new(node)
		for temp = l.startingNode; temp.nextNode != toDelete; temp = temp.nextNode {
		}
		if toDelete.nextNode == nil {
			temp.nextNode = nil
			l.length--
			return nil
		}
		temp.nextNode = temp.nextNode.nextNode
		toDelete = nil
		l.length--
	}

	return nil
}

//Remove all instances of any number of objects from the list.
func (l *List) DeleteAll(d ...interface{}) {
	for _, v := range d {
		for ; l.Search(v) == true; {
			l.Delete(v)
		}
	}
}

//Remove all instances of an array of objects from the list.
func (l *List) DeleteAllA(d []interface{}) {
	for _, v := range d {
		for ; l.Search(v) == true; {
			l.Delete(v)
		}
	}
}

//Get the object from any position. It may return nil when the index is incorrect.
func (l *List) Get(a int) interface{} {
	if l.length == 0 || a > l.length - 1 {
		return nil
	}

	temp := new(node)
	temp = l.startingNode
	for i := 0; i < a; i++ {
		temp = temp.nextNode
	}

	return temp.data
}

//Get a slice of objects from any position. It may return an error when the indexes are incorrect.
func (l *List) GetFromTo(f, t int) []interface{} {
	if f < 0 || t < 0 {
		return []interface{}
	} else if t > l.Len() - 1 || f > l.Len() - 1 {
		return []interface{}
	} else if t < f {
		return []interface{}
	}
	
	ret := make([]interface{}, t - f + 1)
	for i := 0; i < t - f + 1; i++ {
		ret[i] = l.Get(i)
	}

	return ret
}

//Get the whole slice of objects stored in the list.
func (l *List) GetAll() []interface{} {
	return l.GetFromTo(0, l.Len())
}

//Erase everything from the stack and return it as []interface{}. Returns an error when the list is empty already.
func (l *List) Empty() ([]interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("The list is empty already.")
	}
	ret := make([]interface{}, l.length)
	for it, it2 := l.startingNode, 0; it != nil; it, it2 = it.nextNode, it2 + 1 {
		ret[it2] = it.data
	}
	for e, e2 := new(node), l.startingNode; e2.nextNode != nil; e, e2 = e2, e2.nextNode {
		e.nextNode = nil
	}
	l.startingNode = nil
	l.endingNode = nil
	l.length = 0
	return ret, nil
}

//Concatenate any number of lists.
func (l *List) Concat(d ...*List) {
	for _, v := range d {
		if v.length == 0 {
			continue
		}
		for temp2 := v.startingNode; temp2 != nil; temp2 = temp2.nextNode {
			l.Append(temp2.data)
		}
		l.length += v.length
	}
}

//Concatenate and return any number of multisets.
func ConcatLists(d ...*List) *List {
	ret := InitializeList()
	for _, v := range d {
		if v.length == 0 {
			continue
		}
		for temp2 := v.startingNode; temp2 != nil; temp2 = temp2.nextNode {
			ret.Append(temp2.data)
		}
		ret.length += v.length
	}
	return ret
}