package stacker

import (
	"errors"
	"fmt"
	"reflect"
)

type Stack struct {
	data   []interface{}
}

func InitializeStack(d ...interface{}) *Stack {
	s := new(Stack)
	s.data = make([]interface{}, 0)
	s.Push(d...)
	return s
}

func InitializeStackA(d interface{}) *Stack {
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

func (s *Stack) Get(w int) (interface{}, error) {
	if w < 0 {
		return nil, errors.New("Negative subscript.")
	} else if w > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	}
	fmt.Println(s.data[w])
	return s.data[w], nil
}

func (s *Stack) GetFromTo(f, t int) ([]interface{}, error) {
	if f < 0 || t < 0 {
		return nil, errors.New("Negative subscript.")
	} else if t > len(s.data) - 1 || f > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	} else if t < f {
		return nil, errors.New("The second argument is smaller than the second one.")
	}
	return s.data[f:t], nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.data) != 0 {
		return false
	}
	return true
}

func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() == true {
		return nil, errors.New("Can't get the top, because the stack is empty.")
	}
	return s.data[len(s.data) - 1], nil
}

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

func (s *Stack) Push(a ...interface{}) error {
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

func(s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() == true {
		return nil, errors.New("Can't pop anymore.")
	}
	ret := s.data[len(s.data) - 1]
	temp := make([]interface{}, len(s.data) - 1)
	copy(temp, s.data[:len(s.data) - 1])
	s.data = temp
	return ret, nil
}

func (s *Stack) PopN(a int) ([]interface{}, error) {
	if a > len(s.data) {
		return nil, errors.New("Attempted to pop too much at once.")
	}
	ret := make([]interface{}, a)
	copy(ret, s.data[len(s.data) - a:])
	s.data = s.data[:len(s.data) - a]
	return ret, nil
}

func (s *Stack) Empty() []interface{} {
	if len(s.data) == 0 {
		return nil
	}
	ret, _ := s.PopN(len(s.data))
	return ret
}

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

func (s *Stack) Concat(s2 ...*Stack) error {
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
	return nil
}

func ConcatStacks(s ...*Stack) *Stack {
	if len(s) > 0 {
		ret := InitializeStack()
		ret.Concat(s...)
		return ret
	}
	return InitializeStack()
}

func (s *Stack) Len() int {
	return len(s.data)
}