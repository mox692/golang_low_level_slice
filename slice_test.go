package slice

import (
	"fmt"
	"testing"
)

type testInput struct {
	len int
	cap int
	elm []int
}

func Test_createslice(t *testing.T) {

	// negative value -> err, positive value -> noerr
	inputs := map[int]testInput{
		1:  {2, 3, []int{2}},
		2:  {3, 3, []int{}},
		-2: {-1, 3, []int{2}},
		-3: {1, -3, []int{2}},
	}

	for i, v := range inputs {
		fmt.Println(len(inputs))
		ptr, err := createslice(v.len, v.cap, v.elm...)
		if err == nil && ptr == nil {
			t.Errorf("ptr is nil\n")
		}
		if err != nil && i > 0 {
			t.Errorf("expect noerror(%d) but got err: %+v\n", i, err)
		}
		if err == nil && i < 0 {
			t.Errorf("expect error(%d) but got noerr: %+v\n", i, err)
		}
	}
}
