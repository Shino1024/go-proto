package seter

import (
//	"fmt"
	"reflect"
	"errors"
	"sort"
)

type Set struct {
	startingNode *node
	endingNode   *node
	length       int
	lowest       interface{}
	highest      interface{}
}

func merge(a []interface{}, b []interface{}) []interface{} {
	r := make([]interface{}, len(a) + len(b))
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i + j] = a[i]
			i++
		} else {
			r[i + j] = b[j]
			j++
		}
	}
	for i < len(a) {
		r[i + j] = a[i]
		i++
	}
	for j < len(b) {
		r[i + j] = b[j]
		j++
	}
	return r
}

func mergeSort(items []interface{}) []interface{} {
	if len(items) < 2 {
		return items
	}
	middle := len(items) / 2
	a := mergeSort(items[:middle])
	b := mergeSort(items[middle:])
	return merge(a, b)
}

func InitializeSet(d ...interface{}) (*Set, error) {
	temp := reflect.ValueOf(d)
	s := new(Set)
	if temp.Len() == 0 {
		s.startingNode = nil
		s.endingNode = nil
		s.length = 0
		return s
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	if temp.Len() == 1 {
		s.startingNode = new(node)
		s.startingNode.nextNode = nil
		s.startingNode.data = temp2[0]
		s.endingNode = nil
		s.length = 1
		return s
	}
	for i := 0; i < temp.Len(); i++ {
		for j := i + 1; j < temp.Len(); j++ {
			if temp2[i] == temp2[j] {
				return nil, errors.New("Found the same elements in the given data.")
			}
		}
	}
	temp2 = mergeSort(temp2)
	temp3 := new(node)
	temp3.nextNode = nil
	s.startingNode = temp3
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			s.endingNode = temp3
			temp3.nextNode = nil
			break
		}
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	s.length = temp.Len()
	s.lowest = temp2[0]
	s.highest = temp2[len(temp2) - 1]
	return s, nil
}

func InitializeSetA(d interface{}) (*Set, error) {
	temp := reflect.ValueOf(d)
	s := new(Set)
	if temp.Len() == 0 {
		s.startingNode = nil
		s.endingNode = nil
		s.length = 0
		return s
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	if temp.Len() == 1 {
		s.startingNode = new(node)
		s.startingNode.nextNode = nil
		s.startingNode.data = temp2[0]
		s.endingNode = nil
		s.length = 1
		return s
	}
	for i := 0; i < temp.Len(); i++ {
		for j := i + 1; j < temp.Len(); j++ {
			if temp2[i] == temp2[j] {
				return nil, errors.New("Found the same elements in the given data.")
			}
		}
	}
	temp2 = mergeSort(temp2)
	temp3 := new(node)
	temp3.nextNode = nil
	s.startingNode = temp3
	for k, v := range temp2 {
		temp3.data = v
		if k == len(temp2) - 1 {
			s.endingNode = temp3
			temp3.nextNode = nil
			break
		}
		temp3.nextNode = new(node)
		temp3 = temp3.nextNode
	}
	s.length = temp.Len()
	s.lowest = temp2[0]
	s.highest = temp2[len(temp2) - 1]
	return s, nil
}

func (s *Set) Append(d ...interface{}) error {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	for i := 0; i < temp.Len(); i++ {
		for j := i + 1; j < temp.Len(); j++ {
			if temp2[i] == temp2[j] {
				return errors.New("Found the same elements in the given data.")
			}
		}
	}
	temp2 = mergeSort(temp2)
	tempLow, tempHigh := s.lowest, s.highest
	tempNode := new(node)
	count := 0
	lastPoint := l.startingNode
	for _, v := range temp2 {
		if tempLow > v {
			tempNode2 := new(node)
			tempNode2.data = v
			tempNode2.nextNode = s.startingNode
			s.startingNode = tempNode
			s.length++
		} else if tempHigh > v {
			for tempNode = lastPoint; tempNode.nextNode != nil && tempNode.data > tempNode.nextNode.data; lastPoint, tempNode = tempNode, tempNode.nextNode {
			}
			
		} else {
			
		}
	}
}

