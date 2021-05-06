package slice

import (
	"errors"
	"unsafe"
)

type size = int

const (
	TEN   size = 10
	HND   size = 100
	THOUS size = 1000
)

type slice struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
	Size size
}

func Createslice(length, cap int, elm ...int) (*slice, error) {
	if length < 0 || cap < 0 {
		return nil, errors.New("len or cap must be positive number.")
	}

	if length > cap {
		return nil, errors.New("len must be less than cap.")
	}

	if cap < 10 {
		var backgroundArr [TEN]int
		if elm != nil {
			for elmLen, i := len(elm), 0; i < elmLen; i++ {
				backgroundArr[i] = elm[i]
			}
		}
		slice := &slice{Cap: cap, Len: length, Data: unsafe.Pointer(&backgroundArr)}
		return slice, nil
	} else if cap < 100 {
		var backgroundArr [HND]int

		if elm != nil {
			for elmLen, i := len(elm), 0; i < elmLen; i++ {
				backgroundArr[i] = elm[i]
			}
		}

		slice := &slice{Cap: cap, Len: length, Data: unsafe.Pointer(&backgroundArr)}
		return slice, nil
	} else if cap < 1000 {
		var backgroundArr [THOUS]int

		if elm != nil {
			for elmLen, i := len(elm), 0; i < elmLen; i++ {
				backgroundArr[i] = elm[i]
			}
		}

		slice := &slice{Cap: cap, Len: length, Data: unsafe.Pointer(&backgroundArr)}
		return slice, nil
	} else {
		return nil, errors.New("cannot create slice which cap is over 1000.")
	}
}

func (s *slice) at(index int) (unsafe.Pointer, error) {
	if index < 0 {
		return nil, errors.New("invalid index value, at's index must be positive number.")
	}
	if index > s.Cap-1 {
		return nil, errors.New("invalid reference. index must be less than cap.")
	}

	return unsafe.Pointer(uintptr(s.Data) + uintptr(index)*unsafe.Sizeof(int(0))), nil
}

// TODO: errじゃなくてpanicを起こしてもいいかも
func (s *slice) Get(index int) (int, error) {
	if index < 0 {
		return 0, errors.New("index must be positive value.")
	}
	if index > s.Len {
		return 0, errors.New("index must be less than slice's cap.")
	}
	ptr, err := s.at(index)
	if err != nil {
		return 0, err
	}
	return *(*int)(ptr), nil
}

func (s *slice) Set(index, value int) error {
	if index < 0 {
		return errors.New("index must be positive value.")
	}
	if index > s.Len {
		return errors.New("index must be less than slice's cap.")
	}
	ptr, err := s.at(index)
	if err != nil {
		return err
	}
	*(*int)(ptr) = value

	return nil
}

// ※user入力のsliceは許可する
func (s *slice) Append(input []int) error {
	return nil
}

// callbackはerrを返さない想定.
// errは起こらないはず(slice構造体内部のデータからLenを取得してるので)なので、errは握り潰す
func (s *slice) Map(callback func(int) int) {
	for i := 0; i < s.Len; i++ {
		val, _ := s.Get(i)
		result := callback(val)
		s.Set(i, result)
	}
}

func (s *slice) Filter(callback func(int) int) {
}
