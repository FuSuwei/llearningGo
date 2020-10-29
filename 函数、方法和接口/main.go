package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	twice(a)
	fmt.Println(a)
	var b = IntSliceHeader{
		Data: []int{1, 2, 3},
		Len: 3,
		Cap: 3,
	}
	fmt.Println(b.Data)
	twiceIntSliceHeader(b)
	fmt.Println(b.Data)
	fmt.Println(b.Cap)
}

func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

type IntSliceHeader struct {
	Data []int
	Len  int
	Cap  int
}

func twiceIntSliceHeader(x IntSliceHeader) {
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2
	}
	x.Cap += 1
}