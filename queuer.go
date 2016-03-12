package queuer

import (
	"fmt"
	"reflect"
)

type Queue struct {
	startingNode *node
	endingNode   *node
	length       int
}

type node struct {
	nextNode *node
	data     interface{}
}

func InitializeQueue(d ...interface{}) *Queue {
	temp := reflect.ValueOf(d)
	q := new(Queue)
	if temp.Len() == 0 {
		q.startingNode = nil
		q.endingNode = nil
		q.length = 0
		return q
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	temp3 := new(node)
	temp3.nextNode = nil
	q.startingNode = temp3
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			q.endingNode = temp3
			temp3.nextNode = nil
			break
		}
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	q.length = temp.Len()
	return q
}

func InitializeQueueA(d interface{}) *Queue {
	temp := reflect.ValueOf(d)
	q := new(Queue)
	if temp.Len() == 0 {
		q.startingNode = nil
		q.endingNode = nil
		q.length = 0
		return q
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	temp3 := new(node)
	temp3.nextNode = nil
	q.startingNode = temp3
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			q.endingNode = temp3
			temp3.nextNode = nil
			break
		} 
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	q.length = temp.Len()
	return q
}

func (q *Queue) IsEmpty() bool {
	if q.length != 0 {
		return false
	}
	return true
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) Push(d ...interface{}) {
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
	if q.length == 0 {
		q.startingNode = temp3
	} else {
		q.endingNode.nextNode = temp3
	}
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			q.endingNode = temp3
			temp3.nextNode = nil
			break
		} 
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	q.length += temp.Len()
}

func (q *Queue) PushA(d interface{}) {
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
	if q.length == 0 {
		q.startingNode = temp3
	} else {
		q.endingNode.nextNode = temp3
	}
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			q.endingNode = temp3
			temp3.nextNode = nil
			break
		} 
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	q.length += temp.Len()
}

func (q *Queue) PrintAll(sepstr ...string) {
	if q.length == 0 {
		return
	}
	sep := ", "
	if len(sepstr) == 1 {
		sep = sepstr[0]
	}
	fmt.Print("[")
	for n := q.startingNode; n != nil; n = n.nextNode {
		fmt.Print(n.data)
		if n.nextNode == nil {
			break
		}
		fmt.Print(sep)
	}
	fmt.Print("]")
}

func (q *Queue) PrintAllln(sepstr ...string) {
	if q.length == 0 {
		fmt.Println("[]")
		return
	}
	sep := ", "
	if len(sepstr) == 1 {
		sep = sepstr[0]
	}
	fmt.Print("[")
	for n := q.startingNode; n != nil; n = n.nextNode {
		fmt.Print(n.data)
		if n.nextNode == nil {
			break
		}
		fmt.Print(sep)
	}
	fmt.Println("]")
}

func (q *Queue) Search(d interface{}) bool {
	for temp := q.startingNode; temp != nil; temp = temp.nextNode {
		if reflect.DeepEqual(temp.data, d) {
			return true
		}
	}
	return false
}

func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}
	ret := q.startingNode.data
	q.startingNode = q.startingNode.nextNode
	q.length--
	return ret
}

func (q *Queue) PopN(n int) []interface{} {
	if n > q.length {
		return nil
	} else if n == q.length {
		return q.Empty()
	}
	ret := make([]interface{}, n)
	for e, e2, it := new(node), q.startingNode, 0; it < n; e, e2, it = e2, e2.nextNode, it + 1 {
		ret[it] = e2.data
		e.nextNode = nil
		if it + 1 == n {
			q.startingNode = e2.nextNode
		}
	}
	q.length -= n
	return ret
}

func (q *Queue) Empty() []interface{} {
	if q.length == 0 {
		return nil
	}
	ret := make([]interface{}, q.length)
	for it, it2 := q.startingNode, 0; it != nil; it, it2 = it.nextNode, it2 + 1 {
		ret[it2] = it.data
	}
	for e, e2 := new(node), q.startingNode; e2.nextNode != nil; e, e2 = e2, e2.nextNode {
		e.nextNode = nil
	}
	q.startingNode = nil
	q.endingNode = nil
	q.length = 0
	return ret
}

func (q *Queue) Concat(d ...*Queue) {
	for _, v := range d {
		if v.length == 0 {
			continue
		}
		for temp2 := v.startingNode; temp2 != nil; temp2 = temp2.nextNode {
			q.Push(temp2.data)
		}
		q.length += v.length
	}
}

func ConcatLists(d ...*Queue) *Queue {
	ret := InitializeQueue()
	for _, v := range d {
		if v.length == 0 {
			continue
		}
		for temp2 := v.startingNode; temp2 != nil; temp2 = temp2.nextNode {
			ret.Push(temp2.data)
		}
		ret.length += v.length
	}

	return ret
}