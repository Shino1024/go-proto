package multiseter

import (
	"fmt"
	"errors"
)

type Lesser interface {
	Less(other Lesser) bool
}

type Multiset struct {
	data []Lesser
}

func InitializeMultiset(d ...Lesser) *Multiset {
	s := new(Multiset)
	if len(d) == 0 {
		s.data = make([]Lesser, 0)
		return s
	}
	s.Insert(d...)
	return s
}

func InitializeMultisetA(d []Lesser) *Multiset {
	s := new(Multiset)
	if len(d) == 0 {
		s.data = make([]Lesser, 0)
		return s
	}
	s.Insert(d...)
	return s
}

func (s *Multiset) Insert(d ...Lesser) {
	for j := 0; j < len(d); j++ {
		OUTSIDELOOP:
		for i := 0; ; i++ {
			switch {
				case i == s.Len():
				s.data = append(s.data, d[j])
				break OUTSIDELOOP

				case !s.data[i].Less(d[j]):
				s.data = append(s.data[:i], append([]Lesser{d[j]}, s.data[i:]...)...)
				break OUTSIDELOOP
			}
		}
	}
}

func (s *Multiset) InsertA(d []Lesser) {
	for j := 0; j < len(d); j++ {
		OUTSIDELOOP:
		for i := 0; ; i++ {
			switch {
				case i == s.Len():
				s.data = append(s.data, d[j])
				break OUTSIDELOOP

				case !s.data[i].Less(d[j]):
				s.data = append(s.data[:i], append([]Lesser{d[j]}, s.data[i:]...)...)
				break OUTSIDELOOP
			}
		}
	}
}

func (s *Multiset) Delete(d ...Lesser) {
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

func (s *Multiset) DeleteA(d []Lesser) {
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

func (s *Multiset) DeleteAll(d ...Lesser) {
	if len(s.data) == 0 {
		return
	}
	for _, v := range d {
		if checkLeft, checkRight := s.RangeSearch(v); checkLeft == -1 {
			continue
		} else {
			temp := s.data[checkRight + 1:]
			s.data = s.data[:checkLeft]
			s.data = append(s.data, temp...)
		}
	}
}

func (s *Multiset) DeleteAllA(d []Lesser) {
	if len(s.data) == 0 {
		return
	}
	for _, v := range d {
		if checkLeft, checkRight := s.RangeSearch(v); checkLeft == -1 {
			continue
		} else {
			temp := s.data[checkRight + 1:]
			s.data = s.data[:checkLeft]
			s.data = append(s.data, temp...)
		}
	}
}

func (s *Multiset) Search(d Lesser) int {
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

func (s *Multiset) RangeSearch(d Lesser) (int, int) {
	mid := s.Search(d)
	if mid == -1 {
		return -1, -1
	}

	var left, right int
	for left = mid; ; left-- {
		if left == -1 {
			left++
			break
		} else if !s.data[left].Less(d) && !d.Less(s.data[left]) {
			continue
		} else {
			left++
			break
		}
	}

	for right = mid; ; right++ {
		if right == s.Len() {
			right--
			break
		} else if !s.data[right].Less(d) && !d.Less(s.data[right]) {
			continue
		} else {
			right--
			break
		}
	}

	return left, right
}

func (s *Multiset) Get(w int) (Lesser, error) {
	if w < 0 {
		return nil, errors.New("Negative subscript.")
	} else if w > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	}
	return s.data[w], nil
}

func (s *Multiset) GetFromTo(f, t int) ([]Lesser, error) {
	if f < 0 || t < 0 {
		return nil, errors.New("Negative subscript.")
	} else if t > len(s.data) - 1 || f > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	} else if t < f {
		return nil, errors.New("The second argument is smaller than the second one.")
	}
	return s.data[f:t], nil
}

func (m *Multiset) GetAll() []Lesser {
	return m.data
}

func (s *Multiset) IsEmpty() bool {
	if len(s.data) != 0 {
		return false
	}
	return true
}

func (s *Multiset) PrintAll(sepstr ...string) {
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

func (s *Multiset) PrintAllln(sepstr ...string) {
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

func (s *Multiset) Max() Lesser {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data) - 1]
}

func (s *Multiset) Min() Lesser {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[0]
}

func (s *Multiset) Empty() []Lesser {
	if len(s.data) == 0 {
		return nil
	}
	ret, _ := s.GetFromTo(0, len(s.data) - 1)
	s.data := make([]Lesser, 0)
	return ret
}

func (s *Multiset) Len() int {
	return len(s.data)
}

func Union(s ...*Multiset) *Multiset {
	ret := InitializeMultiset()
	if len(s) == 0 {
		return ret
	} else if len(s) == 1 {
		return s[0]
	}

	for _, v := range s {
		for i := 0; i < v.Len(); i++ {
			temp, _ := v.Get(i)
			tempLeft, tempRight := v.RangeSearch(temp)
			if tempLeft != -1 {
				tempLeft1, tempRight1 := ret.RangeSearch(temp)
				if tempLeft1 != -1 {
					for j := 0; j < tempRight - tempLeft + tempLeft1 - tempRight1; j++ {
						ret.Insert(temp)
					}
				} else {
					for j := 0; j < tempRight - tempLeft + 1; j++ {
						ret.Insert(temp)
					}
				}

				i = tempRight
			}
		}
	}

	return ret
}

func (s *Multiset) Subtract(other ...*Multiset) {
	for _, v := range other {
		for i := 0; i < v.Len(); i++ {
			temp, err := v.Get(i)
			if err == nil {
				s.Delete(temp)
			}
		}
	}
}

func Intersection(s ...*Multiset) *Multiset {
	ret := Union(s...)
	ret.Subtract(Difference(s...))
	return ret
}

func Difference(s ...*Multiset) *Multiset {
	if len(s) == 0 {
		return InitializeMultiset()
	} else if len(s) == 1 {
		return s[0]
	}

	ret := InitializeMultiset()
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