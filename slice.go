package slice

import (
	"errors"
	"unsafe"
)

type slice struct {
	Data uintptr
	Len  int
	Cap  int
}

func createslice(length, cap int, elm ...int) (*slice, error) {
	if length < 0 || cap < 0 {
		return nil, errors.New("len or cap must be positive number.")
	}

	if length > cap {
		return nil, errors.New("len must be less than cap.")
	}

	backgroundArr := make([]int, length)
	if elm != nil {
		for elmLen, i := 0, len(elm); i < elmLen; i++ {
			backgroundArr[i] = elm[i]
		}
	}

	slice := &slice{}
	slice.Cap = cap
	slice.Len = length
	slice.Data = uintptr(unsafe.Pointer(&backgroundArr))
	return slice, nil
}
