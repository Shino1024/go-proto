//Package seter provides a simple set structure.
package seter

import (
	"fmt"
	"errors"
)

//Interface SLesser to be implemented for objects that will be stored in the set.
type SLesser interface {
	//SLess takes an SLesser and returns a bool. Compare between the receiver and other SLesser.
	SLess(other SLesser) bool
}

//Set data structure.
type Set struct {
	data []SLesser
}

//Initialize the set. Use any number of arguments. Use objects of a single type implementing SLesser.
func InitializeSet(d ...SLesser) *Set {
	s := new(Set)
	if len(d) == 0 {
		s.data = make([]SLesser, 0)
		return s
	}
	s.Insert(d...)
	return s
}

//Initialize the set. Use an array of objects implementing SLesser. Use objects of a single type implementing SLesser.
func InitializeSetA(d []SLesser) *Set {
	s := new(Set)
	if len(d) == 0 {
		s.data = make([]SLesser, 0)
		return s
	}
	s.Insert(d...)
	return s
}

//Insert the objects with this function. Use any number of arguments implementing SLesser.
func (s *Set) Insert(d ...SLesser) {
	for j := 0; j < len(d); j++ {
		OUTSIDELOOP:
		for i := 0; ; i++ {
			switch {
				case i == s.Len():
				s.data = append(s.data, d[j])
				break OUTSIDELOOP

				case !d[j].SLess(s.data[i]) && !s.data[i].SLess(d[j]):
				break OUTSIDELOOP

				case !s.data[i].SLess(d[j]):
				s.data = append(s.data[:i], append([]SLesser{d[j]}, s.data[i:]...)...)
				break OUTSIDELOOP
			}
		}
	}
}

//Insert the objects with this function. Use an array of objects implementing SLesser.
func (s *Set) InsertA(d []SLesser) {
	for j := 0; j < len(d); j++ {
		OUTSIDELOOP:
		for i := 0; ; i++ {
			switch {
				case i == s.Len():
				s.data = append(s.data, d[j])
				break OUTSIDELOOP

				case !d[j].SLess(s.data[i]) && !s.data[i].SLess(d[j]):
				break OUTSIDELOOP

				case !s.data[i].SLess(d[j]):
				s.data = append(s.data[:i], append([]SLesser{d[j]}, s.data[i:]...)...)
				break OUTSIDELOOP
			}
		}
	}
}

//Check the position of the object in the set. -1 means that there is no such object in the set.
func (s *Set) Search(d SLesser) int {
	min, max := 0, len(s.data) - 1
	for max >= min {
		mid := (max + min) / 2
		if !s.data[mid].SLess(d) && !d.SLess(s.data[mid]) {
			return mid
		} else if s.data[mid].SLess(d) {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

//Remove any number of objects implementing SLesser.
func (s *Set) Delete(d ...SLesser) {
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

//Remove objects implementing SLesser placed in an array.
func (s *Set) DeleteA(d []SLesser) {
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

//Get the object from any position. It may return an error when the index is incorrect. Returns SLesser.
func (s *Set) Get(w int) (SLesser, error) {
	if w < 0 {
		return nil, errors.New("Negative subscript.")
	} else if w > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	}
	return s.data[w], nil
}

//Get a slice of objects from any position. It may return an error when the indexes are incorrect. Returns []SLesser.
func (s *Set) GetFromTo(f, t int) ([]SLesser, error) {
	if f < 0 || t < 0 {
		return nil, errors.New("Negative subscript.")
	} else if t > len(s.data) - 1 || f > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	} else if t < f {
		return nil, errors.New("The second argument is smaller than the second one.")
	}
	return s.data[f:t], nil
}

//Get the whole slice placed in the set. Returns SLesser.
func (s *Set) GetAll() []SLesser {
	return s.data
}

//Check whether the set is empty. Returns bool.
func (s *Set) IsEmpty() bool {
	if len(s.data) != 0 {
		return false
	}
	return true
}

//Print all objects from the set. They, of course, should implement Stringer as well. An optional separator can be provided.
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

//Print all objects from the set and a new line. They, of course, should implement Stringer as well. An optional separator can be provided.
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

//Get the biggest object from the set.
func (s *Set) Max() SLesser {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data) - 1]
}

//Get the smallest object from the set.
func (s *Set) Min() SLesser {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[0]
}

//Erase everything from the set and return it in []SLesser.
func (s *Set) Empty() []SLesser {
	if len(s.data) == 0 {
		return nil
	}
	ret := s.data
	s.data = make([]SLesser, 0)
	return ret
}

//Get the length of the set.
func (s *Set) Len() int {
	return len(s.data)
}

//Concatenate any number of sets.
func (s *Set) Concat(other ...*Set) {
	for _, v := range other {
		s.InsertA(v.GetAll())
	}
}

//Concatenate and return any number of sets.
func ConcatSets(s ...*Set) *Set {
	ret := InitializeSet()
	for _, v := range s {
		ret.InsertA(v.GetAll())
	}

	return ret
}

//Perform the union operation on any number of sets.
func Union(s ...*Set) *Set {
	ret := InitializeSet()
	for _, v := range s {
		ret.InsertA(v.GetAll())
	}

	return ret
}

//Substract any number of sets from the receiver set.
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

//Perform the difference operation on any number of sets.
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

//Perform the intersection operation on any number of sets.
func Intersection(s ...*Set) *Set {
	ret := Union(s...)
	ret.Subtract(Difference(s...))

	return ret
}