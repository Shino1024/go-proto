package stacker

import (
	"errors"
	"fmt"
	"reflect"
)

type Stack struct {
	data   []interface{}
	length int
	maxlen int
}

const MaxCap = 1024

func InitializeStack() *Stack {
	s := new(Stack)
	s.data = make([]interface{}, 0, MaxCap)
	s.length = 0
	s.maxlen = MaxCap
	return s
}

func InitializeStackL(l int) *Stack {
	s := new(Stack)
	s.data = make([]interface{}, 0, l)
	s.length = 0
	s.maxlen = l
	return s
}

func InitializeStackN(d interface{}, l int) *Stack {
	temp := reflect.ValueOf(d)
	s := new(Stack)
	s.data = make([]interface{}, temp.Len(), MaxCap)
	for i := 0; i < temp.Len(); i++ {
		s.data[i] = temp.Index(i).Interface()
	}
	s.length = temp.Len()
	if l < temp.Len() {
		l = temp.Len()
	}
	s.maxlen = l
	return s
}

func InitializeStackA(d interface{}) *Stack {
	temp := reflect.ValueOf(d)
	s := new(Stack)
	s.data = make([]interface{}, temp.Len(), MaxCap)
	for i := 0; i < temp.Len(); i++ {
		s.data[i] = temp.Index(i).Interface()
	}
	s.length = temp.Len()
	s.maxlen = MaxCap
	return s
}

func InitializeStackVar(d ...interface{}) *Stack {
	s := new(Stack)
	s.data = make([]interface{}, len(d), MaxCap)
	for _, v := range d {
		s.Push(v)
	}
	s.length = len(d)
	s.maxlen = MaxCap
	return s
}

func (s *Stack) IsEmpty() bool {
	if s.length != 0 {
		return false
	}
	return true
}

func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() == true {
		return nil, errors.New("Can't get the top, because the stack is empty.")
	}
	return s.data[s.length - 1], nil
}

func (s *Stack) Push(a interface{}) error {
	if s.length + 1 == s.maxlen {
		return errors.New("Cannot push more, not enough space.")
	}
	s.length++
	temp := make([]interface{}, s.length)
	copy(temp, s.data)
	temp[s.length - 1] = a
	s.data = temp
	return nil
}

func (s *Stack) PushA(a interface{}) error {
	temp := reflect.ValueOf(a)
	temp2 := make([]interface{}, temp.Len(), MaxCap)
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	for _, v := range temp2 {
		if temp.Len() + s.length >= s.maxlen {
			return errors.New("Haven't pushed all of the elements, not enough space.")
		}
		s.Push(v)
	}
	return nil
}

func (s *Stack) PushVar(a ...interface{}) error {
	for _, v := range a {
		if len(a) + s.length >= s.maxlen {
			return errors.New("Haven't pushed all of the elements, not enough space.")
		}
		s.Push(v)
	}
	return nil
}

func(s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() == true {
		return nil, errors.New("Can't pop anymore.")
	}
	s.length--
	ret := s.data[s.length]
	temp := make([]interface{}, s.length, s.maxlen)
	copy(temp, s.data[:s.length])
	s.data = temp
	return ret, nil
}

func (s *Stack) PopN(a int) ([]interface{}, error) {
	if a > s.length {
		return nil, errors.New("Attempted to pop too much at once.")
	}
	ret := make([]interface{}, a)
	copy(ret, s.data[len(s.data) - a:])
	s.data = s.data[:len(s.data) - a]
	fmt.Println("s.data - length:", len(s.data))
	s.length -= a
	fmt.Println("s.length = ", s.length)
	fmt.Println(s)
	return ret, nil
}

func (s *Stack) Empty() []interface{} {
	ret, _ := s.PopN(s.length)
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
		if k == s.length - 1 {
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
		if k == s.length - 1 {
			break
		}
		fmt.Print(sep)
	}
	fmt.Println("]")
}

func (s *Stack) Concat(s2 *Stack) error {
	if s.length + s2.length > s.maxlen {
		return errors.New("The second stack is too long.")
	}
	for _, v := range s2.data {
		s.Push(v)
	}
	return nil
}

func (s *Stack) ConcatVar(s2 ...*Stack) error {
	for _, v := range s2 {
		if s.Concat(v) != nil {
			return errors.New("Not all stacks were concatenated, the overall capacity is too great.")
		}
	}
	return nil
}

func ConcatRetVar(s ...*Stack) *Stack {
	if len(s) > 0 {
		ret := new(Stack)
		var maxlen int
		for _, v := range s {
			maxlen += v.maxlen
		}
		ret.maxlen = 2 * maxlen
		for _, v := range s {
			ret.Concat(v)
		}
		return ret
	}
	return InitializeStack()
}

func (s *Stack) ChangeCapBy(a int) error {
	if s.maxlen + a < 1 {
		return errors.New("Attempted to lower the capacity to a non-positive level.")
	} else if s.length > s.maxlen + a {
		return errors.New("The length cannot be greater than the maximal capacity.")
	}
	s.maxlen += a
	temp := make([]interface{}, s.length, s.maxlen)
	copy(temp, s.data)
	s.data = temp
	return nil
}

func (s *Stack) ChangeCapTo(a int) error {
	if a < 1 {
		return errors.New("Attempted to lower the capacity to a non-positive level.");
	} else if s.length > a {
		return errors.New("The length cannot be greater than the maximal capacity.")
	}
	s.maxlen = a
	temp := make([]interface{}, s.length, s.maxlen)
	copy(temp, s.data)
	s.data = temp
	return nil
}

func (s *Stack) Maxlen() int {
	return s.maxlen
}

func (s *Stack) Length() int {
	return s.length
}
