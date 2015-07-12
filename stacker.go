package stacker

import (
	"errors"
	. "strconv"
	"fmt"
)

type Stack struct {
	data   []int
	length int
	maxlen int
}

const MaxCap = 1024

func InitializeStack() *Stack {
	s := new(Stack)
	s.data = make([]int, 0, 4096)
	s.length = 0
	s.maxlen = 4096
	return s
}

func InitializeStackL(l int) *Stack {
	s := new(Stack)
	s.data = make([]int, 0, l)
	s.length = 0
	s.maxlen = l
	return s
}

func InitializeStackN(d []int, l int) *Stack {
	if l < len(d) {
		l = len(d)
	}
	s := new(Stack)
	s.data = make([]int, len(d), l)
	copy(s.data, d)
	s.length = len(d)
	s.maxlen = l
	return s
}

func InitializeStackA(d []int) *Stack {
	s := new(Stack)
	s.data = make([]int, len(d), MaxCap)
	copy(s.data, d)
	s.length = len(d)
	s.maxlen = MaxCap
	return s
}

func InitializeStackVar(d ...int) *Stack {
	s := new(Stack)
	s.data = make([]int, len(d), MaxCap)
	fmt.Println("err", d)
	for _, v := range d {
		s.Push(v)
	}
	s.length = len(d)
	fmt.Println("err", s.data)
	s.maxlen = MaxCap
	return s
}

func (s *Stack) IsEmpty() bool {
	if s.length != 0 {
		return false
	}
	return true
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() == true {
		return 0, errors.New("Can't get the top, because the stack is empty.")
	}
	return s.data[s.length - 1], nil
}

func (s *Stack) Push(a int) error {
	if s.length + 1 == s.maxlen {
		return errors.New("Cannot push more, not enough space.")
	}
	s.length++
	temp := make([]int, s.length)
	copy(temp, s.data)
	temp[s.length - 1] = a
	s.data = temp
	return nil
}

func (s *Stack) PushA(a []int) error {
	for _, v := range a {
		if len(a) + s.length >= s.maxlen {
			return errors.New("Haven't pushed all of the elements, not enough space.")
		}
		s.Push(v)
	}
	return nil
}

func (s *Stack) PushVar(a ...int) error {
	for _, v := range a {
		if len(a) + s.length >= s.maxlen {
			return errors.New("Haven't pushed all of the elements, not enough space.")
		}
		s.Push(v)
	}
	return nil
}

func(s *Stack) Pop() (int, error) {
	if s.IsEmpty() == true {
		return 0, errors.New("Can't pop anymore.")
	}
	s.length--
	ret := s.data[s.length]
	temp := make([]int, s.length, s.maxlen)
	copy(temp, s.data[:s.length])
	s.data = temp
	return ret, nil
}

func (s *Stack) PopN(a int) []int {
	ret := make([]int, a)
	for i := 0; i < a; i++ {
		if _, e := s.Top(); e != nil {
			break
		}
		ret[i], _ = s.Pop()
	}
	for i, j := 0, len(ret) - 1; i < j; i, j = i + 1, j - 1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

func (s *Stack) Empty() []int {
	return s.PopN(s.length)
}

func (s *Stack) PrintAll() {
	fmt.Print("[")
	for k, v := range s.data {
		fmt.Print(v)
		if k == s.length - 1 {
			break
		}
		fmt.Print(", ")
	}
	fmt.Print("]")
}

func (s *Stack) String() string {
	ret := ""
	ret += "["
	for k, v := range s.data {
		ret += Itoa(v)
		if k == s.length - 1 {
			break
		}
		ret += ", "
	}
	ret += "]"
	return ret
}

func (s *Stack) Concat(s2 *Stack) error {
	if s.length + s2.length > s.maxlen {
		return errors.New("The second stack is too long.")
	}
	s.length += s2.length
	fmt.Println("s b4 con:", s)
	for _, v := range s2.data {
		s.Push(v)
		fmt.Println("appended", v)
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
	ret := new(Stack)
	var maxlen, length int
	for _, v := range s {
		maxlen += v.maxlen
		length += v.length
	}
	ret.maxlen = 2 * maxlen
	ret.length = length
	for _, v := range s {
		ret.Concat(v)
	}
	return ret
}

func (s *Stack) ChangeCapBy(a int) error {
	if s.maxlen + a < 1 {
		return errors.New("Attempted to lower the capacity to a non-positive level.")
	} else if s.length > s.maxlen + a {
		return errors.New("The length cannot be greater than the maximal capacity.")
	}
	s.maxlen += a
	return nil
}

func (s *Stack) ChangeCapTo(a int) error {
	if a < 1 {
		return errors.New("Attempted to lower the capacity to a non-positive level.");
	} else if s.length > a {
		return errors.New("The length cannot be greater than the maximal capacity.")
	}
	s.maxlen = a
	return nil
}

func (s *Stack) Maxlen() int {
	return s.maxlen
}

func (s *Stack) Length() int {
	return s.length
}
