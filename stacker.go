package stacker

import (
	"errors"
	"fmt"
	"reflect"
)

type Stack struct {
	data   []interface{}
}

const MaxCap = 1024

func InitializeStack(d ...interface{}) *Stack {
	s := new(Stack)
	s.data = make([]interface{}, 0, MaxCap)
	s.Push(d...)
	return s
}

func InitializeStackL(l int) *Stack {
	s := new(Stack)
	if l < 1 || l > 10000000 {
		s.data = make([]interface{}, 0, l)
	} else {
		s.data = make([]interface{}, 0, MaxCap)
	}
	return s
}

func InitializeStackA(d interface{}) *Stack {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return InitializeStack()
	}
	s := new(Stack)
	s.data = make([]interface{}, temp.Len(), MaxCap)
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
	if len(s.data) == cap(s.data) {
		return errors.New("Cannot push more, not enough space.")
	}
	temp := reflect.ValueOf(a)
	temp2 := make([]interface{}, temp.Len(), MaxCap)
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	if cap(s.data) < len(s.data) + temp.Len() {
		return errors.New("Insufficient space for the elements.")
	}
	temp3 := make([]interface{}, temp.Len() + len(s.data), cap(s.data))
	copy(temp3[0:len(s.data)], s.data)
	copy(temp3[len(s.data):len(s.data) + temp.Len()], temp2)
	s.data = temp3
	return nil
}

func (s *Stack) Push(a ...interface{}) error {
	if len(s.data) == cap(s.data) {
		return errors.New("Cannot push more, not enough space.")
	}
	temp := reflect.ValueOf(a)
	temp2 := make([]interface{}, temp.Len(), MaxCap)
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	if cap(s.data) < len(s.data) + temp.Len() {
		return errors.New("Insufficient space for the elements.")
	}
	temp3 := make([]interface{}, temp.Len() + len(s.data), cap(s.data))
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
	temp := make([]interface{}, len(s.data) - 1, cap(s.data))
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
	if length > cap(s.data) {
		return errors.New("Concatenation failed, not enough space.")
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

func (s *Stack) ChangeCapBy(a int) error {
	if cap(s.data) + a < 1 {
		return errors.New("Attempted to lower the capacity to a non-positive level.")
	} else if len(s.data) > cap(s.data) + a {
		return errors.New("The length cannot be greater than the maximal capacity.")
	}
	temp := make([]interface{}, len(s.data), cap(s.data) + a)
	copy(temp, s.data)
	s.data = temp
	return nil
}

func (s *Stack) ChangeCapTo(a int) error {
	if a < 1 {
		return errors.New("Attempted to lower the capacity to a non-positive level.");
	} else if len(s.data) > a {
		return errors.New("The length cannot be greater than the maximal capacity.")
	}
	temp := make([]interface{}, len(s.data), cap(s.data))
	copy(temp, s.data)
	s.data = temp
	return nil
}

func (s *Stack) Maxlen() int {
	return cap(s.data)
}

func (s *Stack) Length() int {
	return len(s.data)
}

