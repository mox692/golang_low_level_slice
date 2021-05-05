package slice

import (
	"testing"
)

/*
sliceのpointerまでをtestするのは難しいので、errが返ってくるか返ってこないかでtestを行う
*/

type testInput struct {
	len int
	cap int
	elm []int
}

func Test_Createslice(t *testing.T) {

	// negative value -> err, positive value -> noerr
	inputs := map[int]testInput{
		1:  {2, 3, []int{2}},
		2:  {3, 3, []int{}},
		-2: {-1, 3, []int{2}},
		-3: {1, -3, []int{2}},
	}

	for i, v := range inputs {
		ptr, err := Createslice(v.len, v.cap, v.elm...)
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

func Test_at(t *testing.T) {
	slice, err := Createslice(3, 3, 1, 2, 3)
	if err != nil {
		t.Errorf("Createslice Err: %+v\n", err)
	}
	inputs := map[int]int{
		1:  2,
		2:  2,
		-1: -1,
		-2: 3,
	}

	for i, v := range inputs {
		ptr, err := slice.at(v)
		if i < 0 && err == nil {
			t.Errorf("expect err(%d), but got noerr\n", i)
		}
		if i > 0 && ptr == nil {
			t.Errorf("expect noerr(%d), but got noptr\n", i)
		}
	}
}

type getResult struct {
	input  int
	expect int
}

func Test_Get(t *testing.T) {
	testSlice, err := Createslice(2, 4, 1, 2)
	if err != nil {
		t.Errorf("createSliceErr: %+v\n", err)
	}
	inputs := map[int]getResult{
		1:  {input: 1, expect: 2},
		2:  {input: 0, expect: 1},
		-1: {input: 3, expect: 0},
		-2: {input: -1, expect: 0},
	}

	for i, v := range inputs {
		result, err := testSlice.Get(v.input)
		if i < 0 && err == nil {
			t.Errorf("expect err(%d), but got noerr\n", i)
		}
		if i > 0 && err != nil {
			t.Errorf("expect noerr(%d), but got err: %+v\n", i, err)
		}
		if i > 0 && result != v.expect {
			t.Errorf("expect %d, but got %d\n", v.expect, result)
		}
	}
}

type setInput struct {
	index int
	value int
}

func Test_Set(t *testing.T) {
	testSlice, err := Createslice(2, 4, 1, 2)
	if err != nil {
		t.Errorf("createSliceErr: %+v\n", err)
	}

	inputs := map[int]setInput{
		1:  {index: 0, value: 100},
		2:  {index: 1, value: -10000},
		-1: {index: -1, value: 100},
		-2: {index: 4, value: 100},
	}

	for i, v := range inputs {
		err := testSlice.Set(v.index, v.value)
		result, _ := testSlice.Get(v.index)
		if i < 0 && err == nil {
			t.Errorf("expect err(%d), but got noerr\n", i)
		}
		if i > 0 && err != nil {
			t.Errorf("expect noerr(%d), but got err: %+v\n", i, err)
		}
		if i > 0 && v.value != result {
			t.Errorf("expect %d, but got %d\n", v.value, result)
		}
	}
}