//Package stacker provides a simple stack.
package stacker

import (
	"errors"
	"fmt"
	"reflect"
)

//Stack data structure.
type Stack struct {
	data   []interface{}
}

//Initialize the stack. Use any number of interface{}'s.
func InitializeStack(d ...interface{}) *Stack {
	s := new(Stack)
	s.data = make([]interface{}, 0)
	s.Push(d...)
	return s
}

//Initialize the stack. Use an array of interface{}'s.
func InitializeStackA(d []interface{}) *Stack {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return InitializeStack()
	}
	s := new(Stack)
	s.data = make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		s.data[i] = temp.Index(i).Interface()
	}
	return s
}

//Check if the stack is empty.
func (s *Stack) IsEmpty() bool {
	if len(s.data) != 0 {
		return false
	}
	return true
}

//Get the top object of the stack. It may return an error.
func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() == true {
		return nil, errors.New("Can't get the top, because the stack is empty.")
	}
	return s.data[len(s.data) - 1], nil
}

//Push any number of interface{}'s.
func (s *Stack) Push(a ...interface{}) {
	temp := reflect.ValueOf(a)
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	temp3 := make([]interface{}, temp.Len() + len(s.data))
	copy(temp3[0:len(s.data)], s.data)
	copy(temp3[len(s.data):len(s.data) + temp.Len()], temp2)
	s.data = temp3
}

//Push an array of interface{}'s.
func (s *Stack) PushA(a interface{}) error {
	temp := reflect.ValueOf(a)
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	temp3 := make([]interface{}, temp.Len() + len(s.data))
	copy(temp3[0:len(s.data)], s.data)
	copy(temp3[len(s.data):len(s.data) + temp.Len()], temp2)
	s.data = temp3
	return nil
}

//Pop from the stack as an interface{}. Remember to perform type assertion on the returned object in order to make it usable.
func(s *Stack) Pop() interface{} {
	if s.IsEmpty() == true {
		return nil
	}
	ret := s.data[len(s.data) - 1]
	temp := make([]interface{}, len(s.data) - 1)
	copy(temp, s.data[:len(s.data) - 1])
	s.data = temp
	return ret
}

//Pop any number of elements from the stack as []interface{}. Remember to perform type assertion on the returned object in order to make it usable.
func (s *Stack) PopN(a int) []interface{} {
	if a > len(s.data) {
		return nil
	}
	ret := make([]interface{}, a)
	copy(ret, s.data[len(s.data) - a:])
	s.data = s.data[:len(s.data) - a]
	return ret
}

//Erase everything from the stack and return it as []interface{}.
func (s *Stack) Empty() []interface{} {
	if len(s.data) == 0 {
		return nil
	}
	ret := s.PopN(len(s.data))
	return ret
}

//Print all objects from the stack. They, of course, should implement Stringer as well. An optional separator can be provided.
func (s *Stack) PrintAll(sepstr ...string) {
	sep := ", "
	if len(sepstr) == 1 {
		sep = sepstr[0]
	}
	fmt.Print("[")
	for k, v := range s.data {
		fmt.Print(v)
		if k == len(s.data) - 1 {
			break
		}
		fmt.Print(sep)
	}
	fmt.Print("]")
}

//Print all objects from the stack and a new line. They, of course, should implement Stringer as well. An optional separator can be provided.
func (s *Stack) PrintAllln(sepstr ...string) {
	sep := ", "
	if len(sepstr) == 1 {
		sep = sepstr[0]
	}
	fmt.Print("[")
	for k, v := range s.data {
		fmt.Print(v)
		if k == len(s.data) - 1 {
			break
		}
		fmt.Print(sep)
	}
	fmt.Println("]")
}

//Concatenate any number of stacks.
func (s *Stack) Concat(s2 ...*Stack) {
	var length int
	for _, v := range s2 {
		length += len(v.data)
	}
	temp := make([]interface{}, length)
	temp2 := 0
	for _, v := range s2 {
		copy(temp[temp2:temp2 + len(v.data)], v.data)
		temp2 += len(v.data)
	}
	s.Push(temp...)
}

//Concatenate and return any number of stacks.
func ConcatStacks(s ...*Stack) *Stack {
	if len(s) > 0 {
		ret := InitializeStack()
		ret.Concat(s...)
		return ret
	}
	return InitializeStack()
}

//Get the length of the stack.
func (s *Stack) Len() int {
	return len(s.data)
}