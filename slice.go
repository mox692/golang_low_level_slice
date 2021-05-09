package slice

import (
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

func Createslice(length, cap int, elm ...int) *slice {
	if length < 0 || cap < 0 {
		panic("len or cap must be positive number.")
	}

	if length > cap {
		panic("len must be less than cap.")
	}

	if cap < 10 {
		var backgroundArr [TEN]int
		if elm != nil {
			for elmLen, i := len(elm), 0; i < elmLen; i++ {
				backgroundArr[i] = elm[i]
			}
		}
		slice := &slice{Cap: cap, Len: length, Data: unsafe.Pointer(&backgroundArr)}
		return slice
	} else if cap < 100 {
		var backgroundArr [HND]int

		if elm != nil {
			for elmLen, i := len(elm), 0; i < elmLen; i++ {
				backgroundArr[i] = elm[i]
			}
		}

		slice := &slice{Cap: cap, Len: length, Data: unsafe.Pointer(&backgroundArr)}
		return slice
	} else if cap < 1000 {
		var backgroundArr [THOUS]int

		if elm != nil {
			for elmLen, i := len(elm), 0; i < elmLen; i++ {
				backgroundArr[i] = elm[i]
			}
		}

		slice := &slice{Cap: cap, Len: length, Data: unsafe.Pointer(&backgroundArr)}
		return slice
	} else {
		panic("cannot create slice which cap is over 1000.")
	}
}

func (s *slice) at(index int) unsafe.Pointer {
	if index < 0 {
		panic("invalid index value, at's index must be positive number.")
	}
	if index > s.Cap-1 {
		panic("invalid reference. index must be less than cap.")
	}

	return unsafe.Pointer(uintptr(s.Data) + uintptr(index)*unsafe.Sizeof(int(0)))
}

// TODO: errじゃなくてpanicを起こしてもいいかも
func (s *slice) Get(index int) int {
	if index < 0 {
		panic("index must be positive value.")
	}
	if index > s.Len {
		panic("index must be less than slice's cap.")
	}
	ptr := s.at(index)
	return *(*int)(ptr)
}

func (s *slice) Set(index, value int) {
	if index < 0 {
		panic("index must be positive value.")
	}
	if index > s.Len {
		panic("index must be less than slice's cap.")
	}
	ptr := s.at(index)
	*(*int)(ptr) = value

}

// ※user入力のsliceは許可する
func (s *slice) Append(input []int) {
	appended := s.Len + len(input)
	if appended <= s.Cap {
		if s.Cap < 10 {
			var newArr [TEN]int
			for i := 0; i < appended; i++ {
				if i < s.Len {
					val := s.Get(i)
					newArr[i] = val
				} else {
					newArr[i] = input[i-s.Len]
				}
			}
			s.Len = appended
			s.Data = unsafe.Pointer(&newArr)
		} else if s.Cap < 100 {
			var newArr [HND]int
			for i := 0; i < appended; i++ {
				if i < s.Len {
					val := s.Get(i)
					newArr[i] = val
				} else {
					newArr[i] = input[i-s.Len]
				}
			}
			s.Len = appended
			s.Data = unsafe.Pointer(&newArr)
		} else if s.Cap < 1000 {
			var newArr [THOUS]int
			for i := 0; i < appended; i++ {
				if i < s.Len {
					val := s.Get(i)
					newArr[i] = val
				} else {
					newArr[i] = input[i-s.Len]
				}
			}
			s.Len = appended
			s.Data = unsafe.Pointer(&newArr)
		} else if appended >= 1000 {
			panic("this slice can't have cap over 1000!")
		} else {
			panic("Unexpected processing happen!")
		}
	} else if appended > s.Cap {
		if appended < 10 {
			var newArr [TEN]int
			for i := 0; i < appended; i++ {
				if i < s.Len {
					val := s.Get(i)
					newArr[i] = val
				} else {
					newArr[i] = input[i-s.Len]
				}
			}
			s.Len = appended
			s.Cap = TEN
			s.Data = unsafe.Pointer(&newArr)
		} else if appended < 100 {
			var newArr [HND]int
			for i := 0; i < appended; i++ {
				if i < s.Len {
					val := s.Get(i)
					newArr[i] = val
				} else {
					newArr[i] = input[i-s.Len]
				}
			}
			s.Len = appended
			s.Cap = HND
			s.Data = unsafe.Pointer(&newArr)
		} else if appended < 1000 {
			var newArr [THOUS]int
			for i := 0; i < appended; i++ {
				if i < s.Len {
					val := s.Get(i)
					newArr[i] = val
				} else {
					newArr[i] = input[i-s.Len]
				}
			}
			s.Len = appended
			s.Cap = THOUS
			s.Data = unsafe.Pointer(&newArr)
		} else if appended >= 1000 {
			panic("this slice can't have cap over 1000!")
		} else {
			panic("Unexpected processing happen!")
		}
	}
	return
}

// callbackはerrを返さない想定.
// errは起こらないはず(slice構造体内部のデータからLenを取得してるので)なので、errは握り潰す
func (s *slice) Map(callback func(int) int) {
	for i := 0; i < s.Len; i++ {
		val := s.Get(i)
		result := callback(val)
		s.Set(i, result)
	}
}

func (s *slice) Filter(callback func(int) int) {
}
