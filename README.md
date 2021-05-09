## golang_low_level_slice
It's imitation of go slice, without using native go slice.  

I refered to https://tenntenn.dev/ja/posts/qiita-5229bce80ddb688a708a/

## install
`go get -u https://github.com/mox692/golang_low_level_slice`

## usage

```
// Get
s := slice.Createslice(5, 5, 1, 2, 3, 4, 5)
for i := 0; i < s.Len; i++ {
  fmt.Printf("%d ", s.Get(i)) // -> 1 2 3 4 5
}

// Set
s := slice.Createslice(5, 5, 1, 2, 3, 4, 5)
for i := 0; i < s.Len; i++ {
  s.Set(i, s.Get(i)*2)
  fmt.Printf("%d ", s.Get(i)) // -> 2 4 6 8 10 
}

// Append
s = slice.Createslice(3, 3, 1, 2, 3)
add := []int{6, 7}
s.Append(add)
for i := 0; i < s.Len; i++ {
  fmt.Printf("%d ", s.Get(i)) // -> 1 2 3 6 7 
}

// Map
callback := func(a int) int {
  return a * a
}
s.Map(callback)
for i := 0; i < s.Len; i++ {
  fmt.Printf("%d ", s.Get(i))  // -> 1 2 3 4 5
}
```
