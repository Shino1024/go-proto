package seter

import (
	"fmt"
	"reflect"
	"errors"
)

type Set struct {
	data   []interface{}
}

const MaxCap = 1024

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
		s.data = make([]interface{}, 0)
		s.length = 0
		return s
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	for i := 0; i < temp.Len(); i++ {
		for j := i + 1; j < temp.Len(); j++ {
			if temp2[i] == temp2[j] {
				return nil, errors.New("Found the same elements in the given data.")
			}
		}
	}
	temp2 = mergeSort(temp2)
	s.data = temp2
	return s, nil
}

func InitializeSetA(d interface{}) (*Set, error) {
	temp := reflect.ValueOf(d)
	s := new(Set)
	if temp.Len() == 0 {
		s.data = make([]interface{}, 0)
		return s
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	for i := 0; i < temp.Len(); i++ {
		for j := i + 1; j < temp.Len(); j++ {
			if temp2[i] == temp2[j] {
				return nil, errors.New("Found the same elements in the given data.")
			}
		}
	}
	temp2 = mergeSort(temp2)
	s.data = temp2
	return s, nil
}

func (s *Set) Insert(d ...interface{}) error {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return nil
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
	temp3 = make([]interface{}, len(s.data) + temp.Len(), cap(s.data))
	for i, j := 0, 0; ; i, j = i + 1, j + 1 {
		if s.data[i] > temp2[j] && j < len(temp2) {
			temp3[i + j] = temp2[j]
			j++
		} else if i < len(s.data) {
			temp3[i + j] = s.data[i]
			i++
		}
	}
	s.data = temp3
}

func (s *Set) InsertA(d interface{}) error {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return nil
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
	temp3 = make([]interface{}, len(s.data) + temp.Len(), cap(s.data))
	for i, j := 0, 0; ; i, j = i + 1, j + 1 {
		if s.data[i] > temp2[j] && j < len(temp2) {
			temp3[i + j] = temp2[j]
			j++
		} else if i < len(s.data) {
			temp3[i + j] = s.data[i]
			i++
		}
	}
	s.data = temp3
}

func (s *Set) Search(d interface{}) int {
	min, max := 0, len(s.data) - 1
	for max >= min {
		mid := (max + min) / 2
		if s.data[mid] == d {
			return mid
		} else if s.data[mid] < d {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

func (s *Set) Delete(d ...interface{}) {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	for _, v := temp2 {
		if check := s.Search(v); check == -1 {
			continue
		} else if len(s.data) - 1 == check {
			s.data = s.data[len(s.data) - 2]
		} else {
			temp3 := s.data[check + 1:]
			s.data = s.data[:check]
			s.data = append(s.data, temp3)
		}
	}
}

func (s *Set) DeleteA(d interface{}) {
	temp := reflect.ValueOf(d)
	if temp.Len() == 0 {
		return
	}
	temp2 := make([]interface{}, temp.Len())
	for i := 0; i < temp.Len(); i++ {
		temp2[i] = temp.Index(i).Interface()
	}
	for _, v := temp2 {
		if check := s.Search(v); check == -1 {
			continue
		} else if len(s.data) - 1 == check {
			s.data = s.data[len(s.data) - 2]
		} else {
			temp3 := s.data[check + 1:]
			s.data = s.data[:check]
			s.data = append(s.data, temp3)
		}
	}
}

func (s *Set) Get(w int) (interface{}, error) {
	if w < 0 {
		return nil, errors.New("Negative subscript.")
	} else if w > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	}
	fmt.Println(s.data[w])
	return s.data[w], nil
}

func (s *Set) GetFromTo(f, t int) ([]interface{}, error) {
	if f < 0 || t < 0 {
		return nil, errors.New("Negative subscript.")
	} else if t > len(s.data) - 1 || f > len(s.data) - 1 {
		return nil, errors.New("Subscript beyond the scope.")
	} else if t < f {
		return nil, errors.New("The second argument is smaller than the second one.")
	}
	return s.data[f:t], nil
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

func (s *Set) Max() interface{} {
	if s.IsEmpty() != false {
		return nil
	}
	return s.data[len(s.data) - 1]
}

func (s *Set) Min() interface{} {
	if s.IsEmpty() != false {
		return nil
	}
	return s.data[0]
}

func (s *Set) Empty() []interface{} {
	if len(s.data) == 0 {
		return nil
	}
	ret, _ := s.PopN(len(s.data))
	return ret
}

func (s *Set) ChangeCapBy(a int) error {
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

func (s *Set) ChangeCapTo(a int) error {
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

func (s *Set) Maxlen() int {
	return cap(s.data)
}

func (s *Set) Length() int {
	return len(s.data)
}

