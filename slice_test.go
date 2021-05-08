package slice

import (
	"testing"
)

/*
sliceのpointerまでをtestするのは難しいので、errが返ってくるか返ってこないかでtestを行う
*/

/*
	Createslice's test.
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

/*
	at's test.
*/
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

/*
	Get's test.
*/
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

/*
	Set's test.
*/
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

/*
	Map's test.
*/
func cb1(a int) int {
	return a * 2
}
func cb2(a int) int {
	if a < 0 {
		return -1 * a
	} else {
		return a
	}
}

type map_testcase struct {
	len      int
	callback func(int) int
	init     []int
	expected []int
}

func Test_Map(t *testing.T) {

	test_case := []map_testcase{
		{
			len:      5,
			callback: cb1,
			init:     []int{3, 45, 6, 7, 3},
			expected: []int{6, 90, 12, 14, 6},
		},
	}

	for _, v := range test_case {
		slice, err := Createslice(v.len, v.len, v.init...)
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}

		slice.Map(v.callback)

		for i := 0; i < slice.Len; i++ {
			val, _ := slice.Get(i)
			if val != v.expected[i] {
				t.Errorf("expect %d, but got %d\n", v.expected[i], val)
			}
		}
	}

}

/*
	Append's test.
*/
type appendTestCase struct {
	input     []int
	result    []int
	errString string
}

func Test_Append(t *testing.T) {
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		t.Errorf("err: %+v\n", err)
	// 	}
	// }()

	testCase := []appendTestCase{
		{
			input:     []int{3, 4},
			result:    []int{12, 32, 1, 3, 4},
			errString: "",
		},
		{
			input:     []int{3, 412, 32, 1, 3, 4, 12, 32, 1, 3, 4, 3, 412, 32, 1, 3, 4, 12, 32, 1, 3, 4},
			result:    []int{12, 32, 1, 3, 412, 32, 1, 3, 4, 12, 32, 1, 3, 4, 3, 412, 32, 1, 3, 4, 12, 32, 1, 3, 4}, //
			errString: "",
		},
	}

	for _, v := range testCase {
		sampleSlice, err := Createslice(3, 10, 12, 32, 1)
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		sampleSlice.Append(v.input)
		for s := 0; s < sampleSlice.Len; s++ {
			if val, _ := sampleSlice.Get(s); val != v.result[s] {
				t.Errorf("expect %d, but got %d\n", v.result[s], val)
			}
		}
	}
}
