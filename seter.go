package seter

import (
	"fmt"
	"errors"
)

type Lesser interface {
	Less(other interface{}) bool
}

type Set struct {
	data []Lesser
}

func InitializeSet(d Lesser) *Set {
	s := new(Set)
	if len(d) == 0 {
		s.data = make([]Lesser, 0)
		return s
	}
	s.Insert(d...)
	return s
}

func InitializeSetA(d []Lesser) *Set {
	s := new(Set)
	if len(d) == 0 {
		s.data = make([]Lesser, 0)
		return s
	}
	s.Insert(d...)
	return s
}

func (s *Set) Insert(d ...Lesser) {
	for j := 0; j < len(d); j++ {
		OUTSIDELOOP:
		for i := 0; ; i++ {
			switch {
				case i == s.Len():
				s.data = append(s.data, d[j])
				break OUTSIDELOOP

				case !d[j].Less(s.data[i]) && !s.data[i].Less(d[j]):
				break OUTSIDELOOP

				case !s.data[i].Less(d[j]):
				s.data = append(s.data[:i], append([]Lesser{d[j]}, s.data[i:]...)...)
				break OUTSIDELOOP
			}
		}
	}
}

func (s *Set) InsertA(d []Lesser) error {
	for j := 0; j < len(d); j++ {
		OUTSIDELOOP:
		for i := 0; ; i++ {
			switch {
				case i == s.Len():
				s.data = append(s.data, d[j])
				break OUTSIDELOOP

				case !d[j].Less(s.data[i]) && !s.data[i].Less(d[j]):
				break OUTSIDELOOP

				case !s.data[i].Less(d[j]):
				s.data = append(s.data[:i], append([]Lesser{d[j]}, s.data[i:]...)...)
				break OUTSIDELOOP
			}
		}
	}
}

func (s *Set) Search(d Lesser) int {
	min, max := 0, len(s.data) - 1
	for max >= min {
		mid := (max + min) / 2
		if !s.data[mid].Less(d) && !d.Less(s.data[mid]) {
			return mid
		} else if s.data[mid].Less(d) {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

func (s *Set) Delete(d ...Lesser) {
	if len(s.data) == 0 {
		return
	}
	for _, v := range d {
		if check := s.Search(v); check == -1 {
			continue
		} else if len(s.data) - 1 == check {
			s.data = s.data[0:len(s.data) - 1]
		} else {
			temp := s.data[check + 1:]
			s.data = s.data[:check]
			s.data = append(s.data, temp...)
		}
	}
}

func (s *Set) DeleteA(d []Lesser) {
	if len(s.data) == 0 {
		return
	}
	for _, v := range d {
		if check := s.Search(v); check == -1 {
			continue
		} else if len(s.data) - 1 == check {
			s.data = s.data[0:len(s.data) - 1]
		} else {
			temp := s.data[check + 1:]
			s.data = s.data[:check]
			s.data = append(s.data, temp...)
		}
	}
}

func (s *Set) Get(w int) (Lesser, error) {
	if w < 0 {
		return nil, errors.New("Negative subscript.")
	} else if w > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	}
	return s.data[w], nil
}

func (s *Set) GetFromTo(f, t int) ([]Lesser, error) {
	if f < 0 || t < 0 {
		return nil, errors.New("Negative subscript.")
	} else if t > len(s.data) - 1 || f > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	} else if t < f {
		return nil, errors.New("The second argument is smaller than the second one.")
	}
	return s.data[f:t], nil
}

func (s *Set) GetAll() []Lesser {
	return s.data
}

func (s *Set) IsEmpty() bool {
	if len(s.data) != 0 {
		return false
	}
	return true
}

func (s *Set) PrintAll(sepstr ...string) {
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

func (s *Set) PrintAllln(sepstr ...string) {
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

func (s *Set) Max() Lesser {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data) - 1]
}

func (s *Set) Min() Lesser {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[0]
}

func (s *Set) Empty() []Lesser {
	if len(s.data) == 0 {
		return nil
	}
	ret, _ := s.GetFromTo(0, len(s.data) - 1)
	s.data := make([]Lesser, 0)
	return ret
}

func (s *Set) Len() int {
	return len(s.data)
}

func (s *Set) Concat(other ...*Set) {
	for _, v := range other {
		s.InsertA(other.GetAll())
	}
}

func ConcatSets(s ...*Set) *Set {
	ret := InitializeSet()
	for _, v := range s {
		ret.InsertA(v.GetAll())
	}

	return ret
}

func Union(s ...*Set) *Set {
	ret := InitializeSet()
	for _, v := range s {
		ret.InsertA(v.GetAll())
	}

	return ret
}

func (s *Set) Subtract(other ...*Set) {
	for _, v := range other {
		for i := 0; i < v.Len(); i++ {
			temp, err := v.Get(i)
			if err == nil {
				s.Delete(temp)
			}
		}
	}
}

func Difference(s ...*Set) *Set {
	if len(s) == 0 {
		return InitializeSet()
	} else if len(s) == 1 {
		return s[0]
	}

	ret := InitializeSet()
	ret.InsertA(s[0].GetAll())
	for i := 1; i < len(s); i++ {
		for j := 0; j < s[i].Len(); j++ {
			temp, _ := s[i].Get(j)
			if ret.Search(temp) == -1 {
				ret.Insert(temp)
			} else {
				ret.Delete(temp)
			}
		}
	}

	return ret
}

func Intersection(s ...*Set) *Set {
	ret := Union(s...)
	ret.Subtract(Difference(s...))

	return ret
}
