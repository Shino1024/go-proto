//Simple multiseter data structure.
package multiseter

import (
	"fmt"
	"errors"
)

//Interface MLesser to be implemented for objects that will be stored in the multiset.
type MLesser interface {
	//MLess takes an MLesser and returns a bool. Compare between the receiver and other MLesser.
	MLess(other MLesser) bool
}

//Multiset data structure.
type Multiset struct {
	data []MLesser
}

//Initialize the multiset. Use any number of arguments. Use objects of a single type implementing MLesser.
func InitializeMultiset(d ...MLesser) *Multiset {
	s := new(Multiset)
	if len(d) == 0 {
		s.data = make([]MLesser, 0)
		return s
	}
	s.Insert(d...)
	return s
}

//Initialize the multiset. Use an array of objects implementing MLesser. Use objects of a single type implementing MLesser.
func InitializeMultisetA(d []MLesser) *Multiset {
	s := new(Multiset)
	if len(d) == 0 {
		s.data = make([]MLesser, 0)
		return s
	}
	s.Insert(d...)
	return s
}

//Insert the objects with this function. Use any number of arguments implementing MLesser.
func (s *Multiset) Insert(d ...MLesser) {
	for j := 0; j < len(d); j++ {
		OUTSIDELOOP:
		for i := 0; ; i++ {
			switch {
				case i == s.Len():
				s.data = append(s.data, d[j])
				break OUTSIDELOOP

				case !s.data[i].MLess(d[j]):
				s.data = append(s.data[:i], append([]MLesser{d[j]}, s.data[i:]...)...)
				break OUTSIDELOOP
			}
		}
	}
}

//Insert the objects with this function. Use an array of objects implementing MLesser.
func (s *Multiset) InsertA(d []MLesser) {
	for j := 0; j < len(d); j++ {
		OUTSIDELOOP:
		for i := 0; ; i++ {
			switch {
				case i == s.Len():
				s.data = append(s.data, d[j])
				break OUTSIDELOOP

				case !s.data[i].MLess(d[j]):
				s.data = append(s.data[:i], append([]MLesser{d[j]}, s.data[i:]...)...)
				break OUTSIDELOOP
			}
		}
	}
}

//Remove any number of objects implementing MLesser.
func (s *Multiset) Delete(d ...MLesser) {
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

//Remove objects implementing MLesser placed in an array.
func (s *Multiset) DeleteA(d []MLesser) {
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

//Remove all instances of any number of objects implementing MLesser.
func (s *Multiset) DeleteAll(d ...MLesser) {
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

//Remove any number of objects implementing MLesser placed in an array.
func (s *Multiset) DeleteAllA(d []MLesser) {
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

//Check the position of the object in the multiset. -1 means that there is no such object in the multiset.
func (s *Multiset) Search(d MLesser) int {
	min, max := 0, len(s.data) - 1
	for max >= min {
		mid := (max + min) / 2
		if !s.data[mid].MLess(d) && !d.MLess(s.data[mid]) {
			return mid
		} else if s.data[mid].MLess(d) {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

//Check the first and the last position of the object in the multiset. -1 means that there is no such object in the multiset.
func (s *Multiset) RangeSearch(d MLesser) (int, int) {
	mid := s.Search(d)
	if mid == -1 {
		return -1, -1
	}

	var left, right int
	for left = mid; ; left-- {
		if left == -1 {
			left++
			break
		} else if !s.data[left].MLess(d) && !d.MLess(s.data[left]) {
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
		} else if !s.data[right].MLess(d) && !d.MLess(s.data[right]) {
			continue
		} else {
			right--
			break
		}
	}

	return left, right
}

//Count the number of instances of an MLesser object.
func (s *Multiset) Count(a MLesser) uint {
	if left, right := s.RangeSearch(a); left == -1 {
		return 0
	} else {
		return uint(right) - uint(left) + 1
	}
}

//Get the object from any position. It may return an error when the index is incorrect. Returns MLesser.
func (s *Multiset) Get(w int) (MLesser, error) {
	if w < 0 {
		return nil, errors.New("Negative subscript.")
	} else if w > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	}
	return s.data[w], nil
}

//Get a slice of objects from any position. It may return an error when the indexes are incorrect. Returns []MLesser.
func (s *Multiset) GetFromTo(f, t int) ([]MLesser, error) {
	if f < 0 || t < 0 {
		return nil, errors.New("Negative subscript.")
	} else if t > len(s.data) - 1 || f > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	} else if t < f {
		return nil, errors.New("The second argument is smaller than the second one.")
	}
	return s.data[f:t], nil
}

//Get the whole slice placed in the multiset. Returns MLesser.
func (s *Multiset) GetAll() []MLesser {
	return s.data
}

//Check whether the multiset is empty. Returns bool.
func (s *Multiset) IsEmpty() bool {
	if len(s.data) != 0 {
		return false
	}
	return true
}

//Print all objects from the multiset. They, of course, should implement Stringer as well. An optional separator can be provided.
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

//Print all objects from the multiset and a new line. They, of course, should implement Stringer as well. An optional separator can be provided.
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

//Get the biggest object from the multiset.
func (s *Multiset) Max() MLesser {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data) - 1]
}

//Get the smallest object from the multiset.
func (s *Multiset) Min() MLesser {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[0]
}

//Erase everything from the multiset and return it in []MLesser.
func (s *Multiset) Empty() []MLesser {
	if len(s.data) == 0 {
		return nil
	}
	ret, _ := s.GetFromTo(0, len(s.data) - 1)
	s.data = make([]MLesser, 0)
	return ret
}

//Get the length of the multiset.
func (s *Multiset) Len() int {
	return len(s.data)
}

//Concatenate any number of multisets.
func (s *Multiset) Concat(other ...*Multiset) {
	for _, v := range other {
		s.InsertA(v.GetAll())
	}
}

//Concatenate and return any number of multisets.
func ConcatMultisets(s ...*Multiset) *Multiset {
	ret := InitializeMultiset()
	for _, v := range s {
		ret.InsertA(v.GetAll())
	}

	return ret
}

//Perform the union operation on any number of multisets.
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

//Substract any number of multisets from the receiver multiset.
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

//Perform the difference operation on any number of multisets.
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

//Perform the intersection operation on any number of multisets.
func Intersection(s ...*Multiset) *Multiset {
	ret := Union(s...)
	ret.Subtract(Difference(s...))
	return ret
}