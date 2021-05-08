package slice

import (
	"errors"
	"unsafe"
)

func CreatesliceE(length, cap int, elm ...int) (*slice, error) {
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

func (s *slice) atE(index int) (unsafe.Pointer, error) {
	if index < 0 {
		return nil, errors.New("invalid index value, at's index must be positive number.")
	}
	if index > s.Cap-1 {
		return nil, errors.New("invalid reference. index must be less than cap.")
	}

	return unsafe.Pointer(uintptr(s.Data) + uintptr(index)*unsafe.Sizeof(int(0))), nil
}

func (s *slice) GetE(index int) (int, error) {
	if index < 0 {
		return 0, errors.New("index must be positive value.")
	}
	if index > s.Len {
		return 0, errors.New("index must be less than slice's cap.")
	}
	ptr, err := s.atE(index)
	if err != nil {
		return 0, err
	}
	return *(*int)(ptr), nil
}

func (s *slice) SetE(index, value int) error {
	if index < 0 {
		return errors.New("index must be positive value.")
	}
	if index > s.Len {
		return errors.New("index must be less than slice's cap.")
	}
	ptr, err := s.atE(index)
	if err != nil {
		return err
	}
	*(*int)(ptr) = value

	return nil
}

func (s *slice) AppendE(input []int) error {
	appended := s.Len + len(input)
	if appended < s.Cap {
		if s.Cap < 10 {
			var newArr [TEN]int
			for i := 0; i < appended; i++ {
				if val, err := s.GetE(i); i < s.Len {
					if err != nil {
						return err
					}
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
				if val, err := s.GetE(i); i < s.Len {
					if err != nil {
						return err
					}
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
				if val, err := s.GetE(i); i < s.Len {
					if err != nil {
						return err
					}
					newArr[i] = val
				} else {
					newArr[i] = input[i-s.Len]
				}
			}
			s.Len = appended
			s.Data = unsafe.Pointer(&newArr)
		} else if appended >= 1000 {
			return errors.New("this slice can't have cap over 1000!")
		} else {
			return errors.New("Unexpected processing happen!")
		}
	} else if appended > s.Cap {
		if appended < 10 {
			var newArr [TEN]int
			for i := 0; i < appended; i++ {
				if val, err := s.Get(i); i < s.Len {
					if err != nil {
						return err
					}
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
				if val, _ := s.Get(i); i < s.Len {
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
				if val, _ := s.Get(i); i < s.Len {
					newArr[i] = val
				} else {
					newArr[i] = input[i-s.Len]
				}
			}
			s.Len = appended
			s.Cap = THOUS
			s.Data = unsafe.Pointer(&newArr)
		} else if appended >= 1000 {
			return errors.New("this slice can't have cap over 1000!")
		} else {
			return errors.New("Unexpected processing happen!")
		}
	}
	return nil
}

func (s *slice) MapE(callback func(int) int) error {
	for i := 0; i < s.Len; i++ {
		val, err := s.GetE(i)
		if err != nil {
			return err
		}
		result := callback(val)
		err = s.SetE(i, result)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *slice) FilterE(callback func(int) int) {
}
